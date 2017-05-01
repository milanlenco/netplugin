// Generates Go bindings for all VPP APIs located in the json directory.
// go:generate binapi_generator --input-dir=json --output-dir=go

package srv

import (
	"errors"
	"fmt"

	"github.com/fdio-stack/govpp"
	"github.com/fdio-stack/govpp/api"
	"github.com/fdio-stack/govpp/core"
	"github.com/fdio-stack/govpp/core/bin_api/acl"
	"github.com/fdio-stack/govpp/core/bin_api/af_packet"
	"github.com/fdio-stack/govpp/core/bin_api/interfaces"
	"github.com/fdio-stack/govpp/core/bin_api/l2"
	"github.com/fdio-stack/govpp/core/bin_api/vpe"
	"github.com/fdio-stack/govpp/core/bin_api/vxlan"
)

type vppBridgeDomain struct {
	name         string
	bridgeID     uint32
	hasInterface bool
}

type vppInterface struct {
	name      string
	swIfIndex uint32
	adminUp   uint8
	ipAddr    string
}

type vppRuleT struct {
	index uint32
}

type vppVxlanTunnel struct {
	name      string
	swIfIndex uint32
	tunnels   []vxlanTunnels
}

type vxlanTunnels struct {
	srcAddr string
	dstAddr string
}

// Keeps a map of the associated Contiv Network ID and VPP bridge domains
var vppBridgeByID = make(map[string]*vppBridgeDomain)
var vxlanByID = make(map[string]*vppVxlanTunnel)
var vppIntfByName = make(map[string]*vppInterface)
var vppRuleByID = make(map[string]*vppRuleT)

/*
 ***************************************************************

 *** PUBLIC functions

 ***************************************************************
 */

// VppConnect export the VPP connect function to the public
func VppConnect() (*core.Connection, *api.Channel, error) {
	conn, err := govpp.Connect()
	if err != nil {
		return nil, nil, err
	}
	chann, err := conn.NewAPIChannel()
	if err != nil {
		return nil, nil, err
	}
	return conn, chann, nil
}

// VppAddDelBridgeDomain creates a bridge domain inside VPP
func VppAddDelBridgeDomain(id string, pktTag uint32, isAdd uint8) error {
	if isAdd == 1 {
		vppBridge := vppBridgeDomain{
			id, pktTag, false}
		err := vpp_add_del_l2_bridge_domain(pktTag, isAdd)
		if err != nil {
			return err
		}
		vppBridgeByID[id] = &vppBridge
		return nil
	}
	err := vpp_add_del_l2_bridge_domain(pktTag, isAdd)
	if err != nil {
		return err
	}
	delete(vppBridgeByID, id)
	return nil
}

// VppAddInterface creates an af_packet interface in VPP
func VppAddInterface(vppIntf string) error {
	err := vpp_add_af_packet_interface(vppIntf)
	if err != nil {
		return err
	}
	return nil
}

// VppDelInterface creates an af_packet interface in VPP
func VppDelInterface(vppIntf string) error {
	err := vpp_del_af_packet_interface(vppIntf)
	if err != nil {
		return err
	}
	return nil
}

// VppInterfaceAdminUp sets interface flags state up
func VppInterfaceAdminUp(vppIntf string) error {
	err := vpp_set_vpp_interface_adminup(vppIntf)
	if err != nil {
		return err
	}
	return nil
}

// VppSetInterfaceL2Bridge requests bridge mode for interface
func VppSetInterfaceL2Bridge(id string, interfaceType int, vppIntf string, tunnelIfIndex uint32, shg uint8) error {
	bdid := vppBridgeByID[id].bridgeID
	switch interfaceType {
	case 0:
		intfIndex := vppIntfByName[vppIntf].swIfIndex
		err := vpp_set_interface_l2_bridge(bdid, intfIndex, shg)
		if err != nil {
			return err
		}
	case 1:
		intfIndex := tunnelIfIndex
		err := vpp_set_interface_l2_bridge(bdid, intfIndex, shg)
		if err != nil {
			return err
		}
	default:
		return errors.New("Unkonwn interface type")
	}
	return nil
}

// VppVxlanAddDelTunnel creates or deletes a VXLAN tunnel
func VppVxlanAddDelTunnel(isAdd uint8, isIPv6 uint8, srcAddr []byte,
	dstAddr []byte, vni uint32) (uint32, error) {
	tunnelIfIndex, err := vpp_vxlan_add_del_tunnel(isAdd, isIPv6, srcAddr,
		dstAddr, vni)
	if err != nil {
		return 0, err
	}
	return tunnelIfIndex, nil
}

// VppACLAddReplaceRule adds/replaces a rule in VPP
func VppACLAddReplaceRule(vppRule *acl.ACLRule) error {
	err := vpp_acl_add_replace_rule(vppRule)
	if err != nil {
		return err
	}
	return nil
}

// VppACLDelRule deletes an ACL Rule in vpp
func VppACLDelRule(vppRule *acl.ACLRule) error {
	err := vpp_acl_del_rule(vppRule)
	if err != nil {
		return err
	}
	return nil
}

/*
 ***************************************************************

 *** VPP BRIDGE DOMAIN

 ***************************************************************
 */

func vpp_add_del_l2_bridge_domain(bdid uint32, isAdd uint8) error {

	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	req := &l2.BridgeDomainAddDel{
		BdID:    bdid,
		Flood:   1,
		UuFlood: 1,
		Forward: 1,
		Learn:   1,
		ArpTerm: 1,
		MacAge:  0,
		IsAdd:   isAdd,
	}
	// brecode - change to argument values

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &l2.BridgeDomainAddDelReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	if reply.Retval != 0 {
		return errors.New("Could not add/del bridge domain")
	}

	return nil
}

func vpp_set_interface_l2_bridge(bdid uint32, intfIndex uint32, shg uint8) error {
	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	req := &vpe.SwInterfaceSetL2Bridge{
		RxSwIfIndex: intfIndex,
		BdID:        bdid,
		Shg:         shg,
		Bvi:         0,
		Enable:      1,
	}

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &vpe.SwInterfaceSetL2BridgeReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	if reply.Retval != 0 {
		return errors.New("Could not set bridge mode for interface")
	}
	return nil
}

/*
 ***************************************************************

 *** VPP Interface Add / Del, Set Flags

 ***************************************************************
 */

func vpp_add_af_packet_interface(vppIntf string) error {
	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	req := &af_packet.AfPacketCreate{
		HostIfName:      []byte(vppIntf),
		UseRandomHwAddr: 1,
	}

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &af_packet.AfPacketCreateReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	if reply.Retval != 0 {
		return errors.New("Could not add ad_packet interface")
	}

	vppInt := vppInterface{
		vppIntf,
		reply.SwIfIndex,
		0,
		""}
	vppIntfByName[vppIntf] = &vppInt
	return nil
}

func vpp_del_af_packet_interface(vppIntf string) error {
	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	req := &af_packet.AfPacketDelete{
		HostIfName: []byte(vppIntf),
	}

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &af_packet.AfPacketDeleteReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	if reply.Retval != 0 {
		return errors.New("Could not delete ad_packet interface")
	}

	delete(vppIntfByName, vppIntf)
	return nil
}

func vpp_set_vpp_interface_adminup(vppIntf string) error {
	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	_, ok := vppIntfByName[vppIntf]
	if !ok {
		return errors.New("Interface not found in vppIntfByName")
	}

	req := &interfaces.SwInterfaceSetFlags{
		SwIfIndex:   vppIntfByName[vppIntf].swIfIndex,
		AdminUpDown: 1,
	}

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &interfaces.SwInterfaceSetFlagsReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	if reply.Retval != 0 {
		return errors.New("Could not add set af_packet interface flag, admin state up")
	}
	return nil
}

/*
 ***************************************************************

 *** VPP VXLAN

 ***************************************************************
 */

func vpp_vxlan_add_del_tunnel(isAdd uint8, isIPv6 uint8, srcAddr []byte,
	dstAddr []byte, vni uint32) (uint32, error) {
	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	req := &vxlan.VxlanAddDelTunnel{
		IsAdd:          isAdd,
		IsIpv6:         isIPv6,
		SrcAddress:     srcAddr,
		DstAddress:     dstAddr,
		DecapNextIndex: 0xffffffff,
		Vni:            vni,
	}

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &vxlan.VxlanAddDelTunnelReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	fmt.Printf("%+v\n", reply)
	if reply.Retval != 0 {
		return 0, errors.New("Could not create vxlanTunnel")
	}
	return reply.SwIfIndex, nil
}

/*
 ***************************************************************

 *** VPP ACL

 ***************************************************************
 */

func vpp_acl_add_replace_rule(vppRule *acl.ACLRule) error {
	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	req := &acl.ACLAddReplace{
		ACLIndex: ^uint32(0),
		Tag:      []byte(vppRule.RuleId),
		R: []acl.ACLRule{
			{
				IsPermit:       vppRule.IsPermit,
				SrcIPAddr:      vppRule.SrcIPAddr,
				SrcIPPrefixLen: vppRule.SrcIPPrefixLen,
				DstIPAddr:      vppRule.DstIPAddr,
				DstIPPrefixLen: vppRule.DstIPPrefixLen,
				Proto:          vppRule.Proto,
				SrcportOrIcmptypeFirst: vppRule.SrcportOrIcmptypeFirst,
				SrcportOrIcmptypeLast:  vppRule.SrcportOrIcmptypeLast,
				DstportOrIcmpcodeFirst: vppRule.DstportOrIcmpcodeFirst,
				DstportOrIcmpcodeLast:  vppRule.DstportOrIcmpcodeLast,
			},
		},
	}

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &acl.ACLAddReplaceReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	fmt.Printf("%+v\n", reply)
	if reply.Retval != 0 {
		return errors.New("Could not add set af_packet interface flag, admin state up")
	}
	vppIndexValue := vppRuleT{
		reply.ACLIndex,
	}
	vppRuleByID[vppRule.RuleId] = &vppIndexValue
	return nil
}

func vpp_acl_del_rule(vppRule *acl.ACLRule) error {
	conn, _ := govpp.Connect()
	defer conn.Disconnect()

	ch, _ := conn.NewAPIChannel()
	defer ch.Close()

	req := &acl.ACLDel{
		ACLIndex: vppRuleByID[vppRule.RuleId].index,
	}

	// send the request - channel API instead of SendRequest
	ch.ReqChan <- &api.VppRequest{Message: req}

	// receive the response - channel API instead of ReceiveReply
	vppReply := <-ch.ReplyChan
	reply := &acl.ACLDelReply{}
	ch.MsgDecoder.DecodeMsg(vppReply.Data, reply)

	fmt.Printf("%+v\n", reply)
	if reply.Retval != 0 {
		return errors.New("Could not add set af_packet interface flag, admin state up")
	}
	return nil
}
