/***
Copyright 2017 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vppd

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/contiv/netplugin/core"
	"github.com/contiv/netplugin/netmaster/mastercfg"
	log "github.com/ligato/cn-infra/logging/logrus"

	agentCore "github.com/ligato/cn-infra/core"
	"github.com/ligato/vpp-agent/clientv1/linux/localclient"
	"github.com/ligato/vpp-agent/flavours/linuxlocal"
	vpp_l2 "github.com/ligato/vpp-agent/plugins/defaultplugins/l2plugin/model/l2"
)

// VppDriverOperState carries operational state of the VppDriver.
type VppDriverOperState struct {
	core.CommonState

	// Add any vppdriver specific state for endpoint/network etc
}

// Write the state
func (s *VppDriverOperState) Write() error {
	key := fmt.Sprintf(vppOperPath, s.ID)
	return s.StateDriver.WriteState(key, s, json.Marshal)
}

// Read the state given an ID.
func (s *VppDriverOperState) Read(id string) error {
	key := fmt.Sprintf(vppOperPath, id)
	return s.StateDriver.ReadState(key, s, json.Unmarshal)
}

// ReadAll reads all the state
func (s *VppDriverOperState) ReadAll() ([]core.State, error) {
	return s.StateDriver.ReadAllState(vppOperPathPrefix, s, json.Unmarshal)
}

// Clear removes the state.
func (s *VppDriverOperState) Clear() error {
	key := fmt.Sprintf(vppOperPath, s.ID)
	return s.StateDriver.ClearState(key)
}

const vppDriverID agentCore.PluginName = "contiv-vpp-driver"

// VppDriver holds the operational state of vpp driver
type VppDriver struct {
	oper     VppDriverOperState // Oper state of the driver
	localIP  string             // Local IP address
	lock     sync.Mutex         // lock for modifying shared state
	vppAgent *agentCore.Agent   // VPP agent
}

// Init initializes the VPP driver with VPP agent.
func (d *VppDriver) Init(info *core.InstanceInfo) error {

	if info == nil || info.StateDriver == nil {
		return core.Errorf("Invalid arguments. instance-info: %+v", info)
	}
	d.oper.StateDriver = info.StateDriver
	d.localIP = info.VtepIP
	// restore the driver's runtime state if it exists
	err := d.oper.Read(info.HostLabel)
	if core.ErrIfKeyExists(err) != nil {
		log.Errorf("Failed to read driver oper state for key %q. Error: %s",
			info.HostLabel, err)
		return err
	} else if err != nil {
		// create the oper state as it is first time start up
		d.oper.ID = info.HostLabel

		// write the oper
		err = d.oper.Write()
		if err != nil {
			return err
		}
		// write the oper
		err = d.oper.Write()
		if err != nil {
			return err
		}
	}

	log.Infof("Initializing vppdriver")

	flavour := linuxlocal.Flavour{}
	d.vppAgent = agentCore.NewAgent(log.StandardLogger(), 15*time.Second, flavour.Plugins()...)

	err = d.vppAgent.Start()
	if err != nil {
		log.Fatalf("Error starting VPP agent, Err: %v", err)
		return err
	}

	return nil
}

// Deinit closes the driver. Primarily it stops the associated VPP agent.
func (d *VppDriver) Deinit() {
	log.Infof("Cleaning up vppdriver")

	err := d.vppAgent.Stop()
	if err != nil {
		log.Warnf("Error stopping VPP agent, Err: '%v'", err)
	}
}

// CreateNetwork creates a bridge domain network for a given ID in VPP
// We get the Tenant/vrf and network/subnet info from contiv in this API
func (d *VppDriver) CreateNetwork(id string) error {
	cfgNw := mastercfg.CfgNetworkState{}
	cfgNw.StateDriver = d.oper.StateDriver
	err := cfgNw.Read(id)
	if err != nil {
		log.Errorf("Failed to read network id='%s'", id)
		return err
	}

	log.Infof("Create network id='%s', config='%+v'", id, cfgNw)

	bd := &vpp_l2.BridgeDomains_BridgeDomain{
		Name:                "bd-" + id,
		Flood:               true,
		UnknownUnicastFlood: true,
		Forward:             true,
		Learn:               true,
		ArpTermination:      false,
		MacAge:              0,
		Interfaces:          nil,
	}

	err = localclient.DataChangeRequest(vppDriverID).
		Put().
		BD(bd). /* TODO: VXLANs, ACL */
		Send().ReceiveReply()
	if err != nil {
		log.Errorf("Failed to create network id='%s', Err: %v", id, err)
		return err
	}

	return nil
}

// DeleteNetwork deletes a network for a given ID from VPP
func (d *VppDriver) DeleteNetwork(id, subnet, nwType, encap string, pktTag, extPktTag int, gateway string, tenant string) error {
	log.Infof("Delete network id='%s'", id)

	err := localclient.DataChangeRequest(vppDriverID).
		Delete().
		BD("bd-" + id). /* TODO: VXLANs, ACL */
		Send().ReceiveReply()
	if err != nil {
		log.Errorf("Failed to delete network id='%s', Err: %v", id, err)
		return err
	}

	return nil
}

// CreateEndpoint is not implemented.
func (d *VppDriver) CreateEndpoint(id string) error {
	log.Infof("Not implemented")
	return nil
}

//UpdateEndpointGroup is not implemented.
func (d *VppDriver) UpdateEndpointGroup(id string) error {
	log.Infof("Not implemented")
	return nil
}

// DeleteEndpoint is not implemented.
func (d *VppDriver) DeleteEndpoint(id string) (err error) {
	log.Infof("Not implemented")
	return nil
}

// CreateRemoteEndpoint is not implemented.
func (d *VppDriver) CreateRemoteEndpoint(id string) error {
	log.Infof("Not implemented")
	return nil
}

// DeleteRemoteEndpoint is not implemented.
func (d *VppDriver) DeleteRemoteEndpoint(id string) (err error) {
	log.Infof("Not implemented")
	return nil
}

// CreateHostAccPort is not implemented.
func (d *VppDriver) CreateHostAccPort(id, a string, nw int) (string, error) {
	log.Infof("Not implemented")
	return "", nil
}

// DeleteHostAccPort is not implemented.
func (d *VppDriver) DeleteHostAccPort(id string) (err error) {
	log.Infof("Not implemented")
	return nil
}

// AddPeerHost is not implemented.
func (d *VppDriver) AddPeerHost(node core.ServiceInfo) error {
	log.Infof("Not implemented")
	return nil
}

// DeletePeerHost is not implemented.
func (d *VppDriver) DeletePeerHost(node core.ServiceInfo) error {
	log.Infof("Not implemented")
	return nil
}

// AddMaster is not implemented
func (d *VppDriver) AddMaster(node core.ServiceInfo) error {
	log.Infof("Not implemented")
	return nil
}

// DeleteMaster is not implemented
func (d *VppDriver) DeleteMaster(node core.ServiceInfo) error {
	log.Infof("Not implemented")
	return nil
}

// AddBgp is not implemented.
func (d *VppDriver) AddBgp(id string) (err error) {
	log.Infof("Not implemented")
	return nil
}

// DeleteBgp is not implemented.
func (d *VppDriver) DeleteBgp(id string) (err error) {
	log.Infof("Not implemented")
	return nil
}

// AddSvcSpec is not implemented.
func (d *VppDriver) AddSvcSpec(svcName string, spec *core.ServiceSpec) error {
	log.Infof("Not implemented")
	return nil
}

// DelSvcSpec is not implemented.
func (d *VppDriver) DelSvcSpec(svcName string, spec *core.ServiceSpec) error {
	log.Infof("Not implemented")
	return nil
}

// SvcProviderUpdate is not implemented.
func (d *VppDriver) SvcProviderUpdate(svcName string, providers []string) {
}

// GetEndpointStats is not implemented
func (d *VppDriver) GetEndpointStats() ([]byte, error) {
	log.Infof("Not implemented")
	return []byte{}, nil
}

// InspectState is not implemented
func (d *VppDriver) InspectState() ([]byte, error) {
	log.Infof("Not implemented")
	return []byte{}, nil
}

// InspectBgp is not implemented
func (d *VppDriver) InspectBgp() ([]byte, error) {
	log.Infof("Not implemented")
	return []byte{}, nil
}

// GlobalConfigUpdate is not implemented
func (d *VppDriver) GlobalConfigUpdate(inst core.InstanceInfo) error {
	log.Infof("Not implemented")
	return nil
}

// InspectNameserver is not implemented
func (d *VppDriver) InspectNameserver() ([]byte, error) {
	log.Infof("Not implemented")
	return []byte{}, nil
}

// AddPolicyRule is not implemented
func (d *VppDriver) AddPolicyRule(id string) error {
	log.Infof("Not implemented")
	return nil
}

// DelPolicyRule is not implemented
func (d *VppDriver) DelPolicyRule(id string) error {
	log.Infof("Not implemented")
	return nil
}
