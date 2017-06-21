// Package vxlan_gpe represents the VPP binary API of the 'vxlan_gpe' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//vxlan_gpe.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package vxlan_gpe

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xf2ec9253

// VxlanGpeAddDelTunnel represents the VPP binary API message 'vxlan_gpe_add_del_tunnel'.
// Generated from '/usr/share/vpp/api//vxlan_gpe.api.json', line 6:
//
//        ["vxlan_gpe_add_del_tunnel",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u8", "local", 16],
//            ["u8", "remote", 16],
//            ["u32", "encap_vrf_id"],
//            ["u32", "decap_vrf_id"],
//            ["u8", "protocol"],
//            ["u32", "vni"],
//            ["u8", "is_add"],
//            {"crc" : "0x39488c0f"}
//        ],
//
type VxlanGpeAddDelTunnel struct {
	IsIpv6     uint8
	Local      []byte `struc:"[16]byte"`
	Remote     []byte `struc:"[16]byte"`
	EncapVrfID uint32
	DecapVrfID uint32
	Protocol   uint8
	Vni        uint32
	IsAdd      uint8
}

func (*VxlanGpeAddDelTunnel) GetMessageName() string {
	return "vxlan_gpe_add_del_tunnel"
}
func (*VxlanGpeAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeAddDelTunnel) GetCrcString() string {
	return "39488c0f"
}
func NewVxlanGpeAddDelTunnel() api.Message {
	return &VxlanGpeAddDelTunnel{}
}

// VxlanGpeAddDelTunnelReply represents the VPP binary API message 'vxlan_gpe_add_del_tunnel_reply'.
// Generated from '/usr/share/vpp/api//vxlan_gpe.api.json', line 20:
//
//        ["vxlan_gpe_add_del_tunnel_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x563fedf9"}
//        ],
//
type VxlanGpeAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*VxlanGpeAddDelTunnelReply) GetMessageName() string {
	return "vxlan_gpe_add_del_tunnel_reply"
}
func (*VxlanGpeAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeAddDelTunnelReply) GetCrcString() string {
	return "563fedf9"
}
func NewVxlanGpeAddDelTunnelReply() api.Message {
	return &VxlanGpeAddDelTunnelReply{}
}

// VxlanGpeTunnelDump represents the VPP binary API message 'vxlan_gpe_tunnel_dump'.
// Generated from '/usr/share/vpp/api//vxlan_gpe.api.json', line 27:
//
//        ["vxlan_gpe_tunnel_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x14423443"}
//        ],
//
type VxlanGpeTunnelDump struct {
	SwIfIndex uint32
}

func (*VxlanGpeTunnelDump) GetMessageName() string {
	return "vxlan_gpe_tunnel_dump"
}
func (*VxlanGpeTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeTunnelDump) GetCrcString() string {
	return "14423443"
}
func NewVxlanGpeTunnelDump() api.Message {
	return &VxlanGpeTunnelDump{}
}

// VxlanGpeTunnelDetails represents the VPP binary API message 'vxlan_gpe_tunnel_details'.
// Generated from '/usr/share/vpp/api//vxlan_gpe.api.json', line 34:
//
//        ["vxlan_gpe_tunnel_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "local", 16],
//            ["u8", "remote", 16],
//            ["u32", "vni"],
//            ["u8", "protocol"],
//            ["u32", "encap_vrf_id"],
//            ["u32", "decap_vrf_id"],
//            ["u8", "is_ipv6"],
//            {"crc" : "0xda8ca593"}
//        ]
//
type VxlanGpeTunnelDetails struct {
	SwIfIndex  uint32
	Local      []byte `struc:"[16]byte"`
	Remote     []byte `struc:"[16]byte"`
	Vni        uint32
	Protocol   uint8
	EncapVrfID uint32
	DecapVrfID uint32
	IsIpv6     uint8
}

func (*VxlanGpeTunnelDetails) GetMessageName() string {
	return "vxlan_gpe_tunnel_details"
}
func (*VxlanGpeTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeTunnelDetails) GetCrcString() string {
	return "da8ca593"
}
func NewVxlanGpeTunnelDetails() api.Message {
	return &VxlanGpeTunnelDetails{}
}
