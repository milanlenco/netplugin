package vpppolicy

import (
	"errors"
	"fmt"
	"net"
)

// PolicyRule rules the world
type PolicyRule struct {
	Rule *VppnetPolicyRule // rule definition
}

// VppPolicy should be me
type VppPolicy struct {
	Rules map[string]*PolicyRule // rules database
}

// AddRule  adds a policy to the vpp driver
func AddRule(rule *VppnetPolicyRule) error {
	fmt.Println(rule)
	return nil
}

// DelRule  deletes a policy to the vpp driver
func DelRule(rule *VppnetPolicyRule) error {
	fmt.Println(rule)
	return nil
}

// DstGroupMetadata returns metadata for dst group
func DstGroupMetadata(groupId int) (uint64, uint64) {
	metadata := uint64(groupId) << 1
	metadataMask := uint64(0xfffe)
	metadata = metadata & metadataMask

	return metadata, metadataMask
}

// SrcGroupMetadata returns metadata for src group
func SrcGroupMetadata(groupId int) (uint64, uint64) {
	metadata := uint64(groupId) << 16
	metadataMask := uint64(0x7fff0000)
	metadata = metadata & metadataMask

	return metadata, metadataMask
}

// AddRule adds a security rule to policy table
func (self *PolicyAgent) AddRule(rule *OfnetPolicyRule, ret *bool) error {
	var ipDa *net.IP = nil
	var ipDaMask *net.IP = nil
	var ipSa *net.IP = nil
	var ipSaMask *net.IP = nil
	var md *uint64 = nil
	var mdm *uint64 = nil
	var flag, flagMask uint16
	var flagPtr, flagMaskPtr *uint16
	var err error

	// make sure switch is connected
	if !self.agent.IsSwitchConnected() {
		self.agent.WaitForSwitchConnection()
	}

	// check if we already have the rule
	self.mutex.RLock()
	if self.Rules[rule.RuleId] != nil {
		oldRule := self.Rules[rule.RuleId].Rule

		if ruleIsSame(oldRule, rule) {
			self.mutex.RUnlock()
			return nil
		} else {
			self.mutex.RUnlock()
			log.Errorf("Rule already exists. new rule: {%+v}, old rule: {%+v}", rule, oldRule)
			return errors.New("Rule already exists")
		}
	}
	self.mutex.RUnlock()

	log.Infof("Received AddRule: %+v", rule)

	// Parse dst ip
	if rule.DstIpAddr != "" {
		ipDa, ipDaMask, err = ParseIPAddrMaskString(rule.DstIpAddr)
		if err != nil {
			log.Errorf("Error parsing dst ip %s. Err: %v", rule.DstIpAddr, err)
			return err
		}
	}

	// parse src ip
	if rule.SrcIpAddr != "" {
		ipSa, ipSaMask, err = ParseIPAddrMaskString(rule.SrcIpAddr)
		if err != nil {
			log.Errorf("Error parsing src ip %s. Err: %v", rule.SrcIpAddr, err)
			return err
		}
	}

	// parse source/dst endpoint groups
	if rule.SrcEndpointGroup != 0 && rule.DstEndpointGroup != 0 {
		srcMetadata, srcMetadataMask := SrcGroupMetadata(rule.SrcEndpointGroup)
		dstMetadata, dstMetadataMask := DstGroupMetadata(rule.DstEndpointGroup)
		metadata := srcMetadata | dstMetadata
		metadataMask := srcMetadataMask | dstMetadataMask
		md = &metadata
		mdm = &metadataMask
	} else if rule.SrcEndpointGroup != 0 {
		srcMetadata, srcMetadataMask := SrcGroupMetadata(rule.SrcEndpointGroup)
		md = &srcMetadata
		mdm = &srcMetadataMask
	} else if rule.DstEndpointGroup != 0 {
		dstMetadata, dstMetadataMask := DstGroupMetadata(rule.DstEndpointGroup)
		md = &dstMetadata
		mdm = &dstMetadataMask
	}

	// Setup TCP flags
	if rule.IpProtocol == 6 && rule.TcpFlags != "" {
		switch rule.TcpFlags {
		case "syn":
			flag = TCP_FLAG_SYN
			flagMask = TCP_FLAG_SYN
		case "syn,ack":
			flag = TCP_FLAG_ACK | TCP_FLAG_SYN
			flagMask = TCP_FLAG_ACK | TCP_FLAG_SYN
		case "ack":
			flag = TCP_FLAG_ACK
			flagMask = TCP_FLAG_ACK
		case "syn,!ack":
			flag = TCP_FLAG_SYN
			flagMask = TCP_FLAG_ACK | TCP_FLAG_SYN
		case "!syn,ack":
			flag = TCP_FLAG_ACK
			flagMask = TCP_FLAG_ACK | TCP_FLAG_SYN
		default:
			log.Errorf("Unknown TCP flags: %s, in rule: %+v", rule.TcpFlags, rule)
			return errors.New("Unknown TCP flag")
		}

		flagPtr = &flag
		flagMaskPtr = &flagMask
	}
	// Install the rule in policy table
	ruleFlow, err := self.policyTable.NewFlow(ofctrl.FlowMatch{
		Priority:     uint16(FLOW_POLICY_PRIORITY_OFFSET + rule.Priority),
		Ethertype:    0x0800,
		IpDa:         ipDa,
		IpDaMask:     ipDaMask,
		IpSa:         ipSa,
		IpSaMask:     ipSaMask,
		IpProto:      rule.IpProtocol,
		TcpSrcPort:   rule.SrcPort,
		TcpDstPort:   rule.DstPort,
		UdpSrcPort:   rule.SrcPort,
		UdpDstPort:   rule.DstPort,
		Metadata:     md,
		MetadataMask: mdm,
		TcpFlags:     flagPtr,
		TcpFlagsMask: flagMaskPtr,
	})
	if err != nil {
		log.Errorf("Error adding flow for rule {%v}. Err: %v", rule, err)
		return err
	}

	// Point it to next table
	if rule.Action == "allow" {
		err = ruleFlow.Next(self.nextTable)
		if err != nil {
			log.Errorf("Error installing flow {%+v}. Err: %v", ruleFlow, err)
			return err
		}
	} else if rule.Action == "deny" {
		err = ruleFlow.Next(self.ofSwitch.DropAction())
		if err != nil {
			log.Errorf("Error installing flow {%+v}. Err: %v", ruleFlow, err)
			return err
		}
	} else {
		log.Errorf("Unknown action in rule {%+v}", rule)
		return errors.New("Unknown action in rule")
	}

	// save the rule
	pRule := PolicyRule{
		Rule: rule,
		flow: ruleFlow,
	}
	self.mutex.Lock()
	self.Rules[rule.RuleId] = &pRule
	self.mutex.Unlock()

	return nil
}
