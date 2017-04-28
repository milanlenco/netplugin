// Package gre represents the VPP binary API of the 'gre' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//gre.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package gre

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x6a50ea63

// GreAddDelTunnel represents the VPP binary API message 'gre_add_del_tunnel'.
// Generated from '/usr/share/vpp/api//gre.api.json', line 6:
//
//        ["gre_add_del_tunnel",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_ipv6"],
//            ["u8", "teb"],
//            ["u8", "src_address", 16],
//            ["u8", "dst_address", 16],
//            ["u32", "outer_fib_id"],
//            {"crc" : "0x8ab92528"}
//        ],
//
type GreAddDelTunnel struct {
	IsAdd      uint8
	IsIpv6     uint8
	Teb        uint8
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
	OuterFibID uint32
}

func (*GreAddDelTunnel) GetMessageName() string {
	return "gre_add_del_tunnel"
}
func (*GreAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GreAddDelTunnel) GetCrcString() string {
	return "8ab92528"
}
func NewGreAddDelTunnel() api.Message {
	return &GreAddDelTunnel{}
}

// GreAddDelTunnelReply represents the VPP binary API message 'gre_add_del_tunnel_reply'.
// Generated from '/usr/share/vpp/api//gre.api.json', line 18:
//
//        ["gre_add_del_tunnel_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x754a5956"}
//        ],
//
type GreAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*GreAddDelTunnelReply) GetMessageName() string {
	return "gre_add_del_tunnel_reply"
}
func (*GreAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GreAddDelTunnelReply) GetCrcString() string {
	return "754a5956"
}
func NewGreAddDelTunnelReply() api.Message {
	return &GreAddDelTunnelReply{}
}

// GreTunnelDump represents the VPP binary API message 'gre_tunnel_dump'.
// Generated from '/usr/share/vpp/api//gre.api.json', line 25:
//
//        ["gre_tunnel_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x23d04dc0"}
//        ],
//
type GreTunnelDump struct {
	SwIfIndex uint32
}

func (*GreTunnelDump) GetMessageName() string {
	return "gre_tunnel_dump"
}
func (*GreTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GreTunnelDump) GetCrcString() string {
	return "23d04dc0"
}
func NewGreTunnelDump() api.Message {
	return &GreTunnelDump{}
}

// GreTunnelDetails represents the VPP binary API message 'gre_tunnel_details'.
// Generated from '/usr/share/vpp/api//gre.api.json', line 32:
//
//        ["gre_tunnel_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "is_ipv6"],
//            ["u8", "teb"],
//            ["u8", "src_address", 16],
//            ["u8", "dst_address", 16],
//            ["u32", "outer_fib_id"],
//            {"crc" : "0xdd1d50aa"}
//        ]
//
type GreTunnelDetails struct {
	SwIfIndex  uint32
	IsIpv6     uint8
	Teb        uint8
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
	OuterFibID uint32
}

func (*GreTunnelDetails) GetMessageName() string {
	return "gre_tunnel_details"
}
func (*GreTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GreTunnelDetails) GetCrcString() string {
	return "dd1d50aa"
}
func NewGreTunnelDetails() api.Message {
	return &GreTunnelDetails{}
}
