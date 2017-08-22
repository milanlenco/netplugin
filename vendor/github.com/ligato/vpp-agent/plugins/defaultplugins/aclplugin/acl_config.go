// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package aclplugin is the implementation of the ACL Plugin.
package aclplugin

import (
	"fmt"

	"git.fd.io/govpp.git/api"
	log "github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/cn-infra/utils/safeclose"
	"github.com/ligato/vpp-agent/idxvpp"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/aclplugin/model/acl"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/aclplugin/vppcalls"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/ifplugin/ifaceidx"
	"github.com/ligato/vpp-agent/plugins/govppmux"
)

// ACLConfigurator runs in the background in its own goroutine where it watches for any changes
// in the configuration of ACLs as modelled by the proto file "../model/acl/acl.proto" and stored
// in ETCD under the key "/vnf-agent/{agent-label}/vpp/config/v1/acl/". Updates received from the northbound API
// are compared with the VPP run-time configuration and differences are applied through the VPP binary API.
type ACLConfigurator struct {
	GoVppmux       *govppmux.GOVPPPlugin
	ACLL3L4Indexes idxvpp.NameToIdxRW
	ACLL2Indexes   idxvpp.NameToIdxRW // mapping for L2 ACLs
	SwIfIndexes    ifaceidx.SwIfIndex
	vppChannel     *api.Channel
}

// Init goroutines, channels and mappings
func (plugin *ACLConfigurator) Init() (err error) {
	log.Infof("Initializing plugin ACL plugin")

	// Init VPP API channel
	plugin.vppChannel, err = plugin.GoVppmux.NewAPIChannel()
	if err != nil {
		return err
	}

	err = vppcalls.CheckMsgCompatibilityForACL(plugin.vppChannel)

	// todo possibly check acl plugin version on vpp using bin api acl_plugin_get_version

	return err
}

// Close GOVPP channel
func (plugin *ACLConfigurator) Close() {
	safeclose.Close(plugin.vppChannel)
}

// ConfigureACL creates access list with provided rules and sets this list to every relevant interface
func (plugin *ACLConfigurator) ConfigureACL(acl *acl.AccessLists_Acl) error {
	log.Infof("Configuring new ACL %v", acl.AclName)

	if acl.Rules != nil && len(acl.Rules) > 0 {
		// Validate rules
		rules, isL2MacIP := validateRules(acl.AclName, acl.Rules)
		// Configure ACL rules
		var vppACLIndex uint32
		var err error
		if isL2MacIP {
			vppACLIndex, err = vppcalls.AddMacIPAcl(rules, acl.AclName, plugin.vppChannel)
			if err != nil {
				return err
			}
			// Index used for L2 registration is ACLIndex + 1 (ACL indexes start from 0)
			agentACLIndex := vppACLIndex + 1
			plugin.ACLL2Indexes.RegisterName(acl.AclName, agentACLIndex, nil)
			log.Debugf("ACL %v registered with index %v", acl.AclName, agentACLIndex)
		} else {
			vppACLIndex, err = vppcalls.AddIPAcl(rules, acl.AclName, plugin.vppChannel)
			if err != nil {
				return err
			}
			// Index used for L3L4 registration is aclIndex + 1 (ACL indexes start from 0)
			agentACLIndex := vppACLIndex + 1
			plugin.ACLL3L4Indexes.RegisterName(acl.AclName, agentACLIndex, nil)
			log.Debugf("ACL %v registered with index %v", acl.AclName, agentACLIndex)
		}

		// Set ACL to interfaces
		if acl.Interfaces != nil {
			if isL2MacIP {
				err := vppcalls.SetMacIPAclToInterface(vppACLIndex, acl.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
			} else {
				err := vppcalls.SetACLToInterfacesAsIngress(vppACLIndex, acl.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
				err = vppcalls.SetACLToInterfacesAsEgress(vppACLIndex, acl.Interfaces.Egress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
			}
		} else {
			log.Infof("No interface configured for ACL %v", acl.AclName)
		}
	}

	return nil
}

// ModifyACL modifies previously created access list. L2 access list is removed and recreated, L3/L4 access list is
// modified directly. List of interfaces is refreshed as well
func (plugin *ACLConfigurator) ModifyACL(oldACL *acl.AccessLists_Acl, newACL *acl.AccessLists_Acl) error {
	log.Infof("Modifying ACL %v", oldACL.AclName)

	if newACL.Rules != nil {
		// Validate rules
		rules, isL2MacIP := validateRules(newACL.AclName, newACL.Rules)
		var vppACLIndex uint32
		if isL2MacIP {
			agentACLIndex, _, found := plugin.ACLL2Indexes.LookupIdx(newACL.AclName)
			if !found {
				log.Infof("Acl %v index not found", newACL.AclName)
				return nil
			}
			// Index used in VPP = index used in mapping - 1
			vppACLIndex = agentACLIndex - 1
		} else {
			agentACLIndex, _, found := plugin.ACLL3L4Indexes.LookupIdx(newACL.AclName)
			if !found {
				log.Infof("Acl %v index not found", newACL.AclName)
				return nil
			}
			vppACLIndex = agentACLIndex - 1
		}
		if isL2MacIP {
			// L2 ACL
			err := vppcalls.DeleteMacIPAcl(vppACLIndex, plugin.vppChannel)
			if err != nil {
				return err
			}
			plugin.ACLL2Indexes.UnregisterName(newACL.AclName)
			newVppACLIndex, err := vppcalls.AddMacIPAcl(rules, newACL.AclName, plugin.vppChannel)
			if err != nil {
				return err
			}
			// Create agent index by incrementing vpp one
			newAgentACLIndex := newVppACLIndex + 1
			plugin.ACLL2Indexes.RegisterName(newACL.AclName, newAgentACLIndex, nil)
		} else {
			// L3/L4 ACL can be modified directly
			err := vppcalls.ModifyIPAcl(vppACLIndex, rules, newACL.AclName, plugin.vppChannel)
			if err != nil {
				return err
			}
			// There is no need to update index because modified ACL keeps the old one
		}

		// Update interfaces
		if isL2MacIP {
			// Remove L2 ACL from old interfaces
			if oldACL.Interfaces != nil {
				err := vppcalls.RemoveMacIPIngressACLFromInterfaces(vppACLIndex, oldACL.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
			}
			// Put L2 ACL to new interfaces
			if newACL.Interfaces != nil {
				err := vppcalls.SetMacIPAclToInterface(vppACLIndex, newACL.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
			}

		} else {
			// Remove L3/L4 ACL from old interfaces
			if oldACL.Interfaces != nil {
				err := vppcalls.RemoveIPIngressACLFromInterfaces(vppACLIndex, oldACL.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
				err = vppcalls.RemoveIPEgressACLFromInterfaces(vppACLIndex, oldACL.Interfaces.Egress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
			}
			// Put L3/L4 ACL to new interfaces
			if newACL.Interfaces != nil {
				err := vppcalls.SetACLToInterfacesAsIngress(vppACLIndex, newACL.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
				err = vppcalls.SetACLToInterfacesAsEgress(vppACLIndex, newACL.Interfaces.Egress, plugin.SwIfIndexes, plugin.vppChannel)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// DeleteACL removes existing ACL. To detach ACL from interfaces, list of interfaces has to be provided
func (plugin *ACLConfigurator) DeleteACL(acl *acl.AccessLists_Acl) error {
	log.Infof("Deleting ACL %v", acl.AclName)
	// Get ACL index. Keep in mind that all ACL Indices were incremented by 1
	agentL2AclIndex, _, l2AclFound := plugin.ACLL2Indexes.LookupIdx(acl.AclName)
	agentL3L4AclIndex, _, l3l4AclFound := plugin.ACLL3L4Indexes.LookupIdx(acl.AclName)
	if !l2AclFound && !l3l4AclFound {
		return fmt.Errorf("Acl %v not found, cannot be removed", acl.AclName)
	}
	if l2AclFound {
		// Remove interfaces from L2 ACL
		vppACLIndex := agentL2AclIndex - 1
		if acl.Interfaces != nil {
			err := vppcalls.RemoveMacIPIngressACLFromInterfaces(vppACLIndex, acl.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
			if err != nil {
				return err
			}
		}
		// Remove ACL L2
		err := vppcalls.DeleteMacIPAcl(vppACLIndex, plugin.vppChannel)
		if err != nil {
			return err
		}
		// Unregister
		plugin.ACLL2Indexes.UnregisterName(acl.AclName)
	}
	if l3l4AclFound {
		// Remove interfaces
		vppACLIndex := agentL3L4AclIndex - 1
		if acl.Interfaces != nil {
			err := vppcalls.RemoveIPIngressACLFromInterfaces(vppACLIndex, acl.Interfaces.Ingress, plugin.SwIfIndexes, plugin.vppChannel)
			if err != nil {
				return err
			}
			err = vppcalls.RemoveIPEgressACLFromInterfaces(vppACLIndex, acl.Interfaces.Egress, plugin.SwIfIndexes, plugin.vppChannel)
			if err != nil {
				return err
			}
		}
		// Remove ACL L3/L4
		err := vppcalls.DeleteIPAcl(vppACLIndex, plugin.vppChannel)
		if err != nil {
			return err
		}
		// Unregister
		plugin.ACLL3L4Indexes.UnregisterName(acl.AclName)
	}

	return nil
}

// Validate rules provided in ACL. Every rule has to contain actions and matches.
// Current limitation: L2 and L3/4 have to be split to different ACLs, there cannot be L2 rules and L3/4 rules in the
// same ACL
func validateRules(aclName string, rules []*acl.AccessLists_Acl_Rule) ([]*acl.AccessLists_Acl_Rule, bool) {
	var validL3L4Rules []*acl.AccessLists_Acl_Rule
	var validL2Rules []*acl.AccessLists_Acl_Rule

	for index, rule := range rules {
		if rule.Actions == nil {
			log.Warnf("Rule %v from acl %v does not contain actions", index, aclName)
			continue
		}
		if rule.Matches == nil {
			log.Warnf("Rule %v from acl %v does not contain matches", index, aclName)
			continue
		}
		if rule.Matches.IpRule != nil {
			validL3L4Rules = append(validL3L4Rules, rule)
		}
		if rule.Matches.MacipRule != nil {
			validL2Rules = append(validL2Rules, rule)
		}
	}
	if len(validL3L4Rules) > 0 && len(validL2Rules) > 0 {
		log.Errorf("Acl %v contains even L2 rules and L3/L4 rules. This case is not supported yet, only L3/L4 rules will be resolved",
			aclName)
		return validL3L4Rules, false
	} else if len(validL3L4Rules) > 0 {
		return validL3L4Rules, false
	} else {
		return validL2Rules, true
	}
}
