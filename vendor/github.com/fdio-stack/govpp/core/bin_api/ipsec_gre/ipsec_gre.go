// Package ipsec_gre represents the VPP binary API of the 'ipsec_gre' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//ipsec_gre.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package ipsec_gre

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x163cb35d

// IpsecGreAddDelTunnel represents the VPP binary API message 'ipsec_gre_add_del_tunnel'.
// Generated from '/usr/share/vpp/api//ipsec_gre.api.json', line 6:
//
//        ["ipsec_gre_add_del_tunnel",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "local_sa_id"],
//            ["u32", "remote_sa_id"],
//            ["u8", "is_add"],
//            ["u8", "src_address", 4],
//            ["u8", "dst_address", 4],
//            {"crc" : "0x8e39c05c"}
//        ],
//
type IpsecGreAddDelTunnel struct {
	LocalSaID  uint32
	RemoteSaID uint32
	IsAdd      uint8
	SrcAddress []byte `struc:"[4]byte"`
	DstAddress []byte `struc:"[4]byte"`
}

func (*IpsecGreAddDelTunnel) GetMessageName() string {
	return "ipsec_gre_add_del_tunnel"
}
func (*IpsecGreAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecGreAddDelTunnel) GetCrcString() string {
	return "8e39c05c"
}
func NewIpsecGreAddDelTunnel() api.Message {
	return &IpsecGreAddDelTunnel{}
}

// IpsecGreAddDelTunnelReply represents the VPP binary API message 'ipsec_gre_add_del_tunnel_reply'.
// Generated from '/usr/share/vpp/api//ipsec_gre.api.json', line 17:
//
//        ["ipsec_gre_add_del_tunnel_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x91f136aa"}
//        ],
//
type IpsecGreAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*IpsecGreAddDelTunnelReply) GetMessageName() string {
	return "ipsec_gre_add_del_tunnel_reply"
}
func (*IpsecGreAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecGreAddDelTunnelReply) GetCrcString() string {
	return "91f136aa"
}
func NewIpsecGreAddDelTunnelReply() api.Message {
	return &IpsecGreAddDelTunnelReply{}
}

// IpsecGreTunnelDump represents the VPP binary API message 'ipsec_gre_tunnel_dump'.
// Generated from '/usr/share/vpp/api//ipsec_gre.api.json', line 24:
//
//        ["ipsec_gre_tunnel_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x63d70659"}
//        ],
//
type IpsecGreTunnelDump struct {
	SwIfIndex uint32
}

func (*IpsecGreTunnelDump) GetMessageName() string {
	return "ipsec_gre_tunnel_dump"
}
func (*IpsecGreTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecGreTunnelDump) GetCrcString() string {
	return "63d70659"
}
func NewIpsecGreTunnelDump() api.Message {
	return &IpsecGreTunnelDump{}
}

// IpsecGreTunnelDetails represents the VPP binary API message 'ipsec_gre_tunnel_details'.
// Generated from '/usr/share/vpp/api//ipsec_gre.api.json', line 31:
//
//        ["ipsec_gre_tunnel_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "local_sa_id"],
//            ["u32", "remote_sa_id"],
//            ["u8", "src_address", 4],
//            ["u8", "dst_address", 4],
//            {"crc" : "0x1fe2b580"}
//        ]
//
type IpsecGreTunnelDetails struct {
	SwIfIndex  uint32
	LocalSaID  uint32
	RemoteSaID uint32
	SrcAddress []byte `struc:"[4]byte"`
	DstAddress []byte `struc:"[4]byte"`
}

func (*IpsecGreTunnelDetails) GetMessageName() string {
	return "ipsec_gre_tunnel_details"
}
func (*IpsecGreTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecGreTunnelDetails) GetCrcString() string {
	return "1fe2b580"
}
func NewIpsecGreTunnelDetails() api.Message {
	return &IpsecGreTunnelDetails{}
}
