// Package l2tp represents the VPP binary API of the 'l2tp' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//l2tp.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package l2tp

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x76ecd7d3

// L2tpv3CreateTunnel represents the VPP binary API message 'l2tpv3_create_tunnel'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 6:
//
//        ["l2tpv3_create_tunnel",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "client_address", 16],
//            ["u8", "our_address", 16],
//            ["u8", "is_ipv6"],
//            ["u32", "local_session_id"],
//            ["u32", "remote_session_id"],
//            ["u64", "local_cookie"],
//            ["u64", "remote_cookie"],
//            ["u8", "l2_sublayer_present"],
//            ["u32", "encap_vrf_id"],
//            {"crc" : "0xd9839424"}
//        ],
//
type L2tpv3CreateTunnel struct {
	ClientAddress     []byte `struc:"[16]byte"`
	OurAddress        []byte `struc:"[16]byte"`
	IsIpv6            uint8
	LocalSessionID    uint32
	RemoteSessionID   uint32
	LocalCookie       uint64
	RemoteCookie      uint64
	L2SublayerPresent uint8
	EncapVrfID        uint32
}

func (*L2tpv3CreateTunnel) GetMessageName() string {
	return "l2tpv3_create_tunnel"
}
func (*L2tpv3CreateTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*L2tpv3CreateTunnel) GetCrcString() string {
	return "d9839424"
}
func NewL2tpv3CreateTunnel() api.Message {
	return &L2tpv3CreateTunnel{}
}

// L2tpv3CreateTunnelReply represents the VPP binary API message 'l2tpv3_create_tunnel_reply'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 21:
//
//        ["l2tpv3_create_tunnel_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x7007338f"}
//        ],
//
type L2tpv3CreateTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*L2tpv3CreateTunnelReply) GetMessageName() string {
	return "l2tpv3_create_tunnel_reply"
}
func (*L2tpv3CreateTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*L2tpv3CreateTunnelReply) GetCrcString() string {
	return "7007338f"
}
func NewL2tpv3CreateTunnelReply() api.Message {
	return &L2tpv3CreateTunnelReply{}
}

// L2tpv3SetTunnelCookies represents the VPP binary API message 'l2tpv3_set_tunnel_cookies'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 28:
//
//        ["l2tpv3_set_tunnel_cookies",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u64", "new_local_cookie"],
//            ["u64", "new_remote_cookie"],
//            {"crc" : "0x3e783c95"}
//        ],
//
type L2tpv3SetTunnelCookies struct {
	SwIfIndex       uint32
	NewLocalCookie  uint64
	NewRemoteCookie uint64
}

func (*L2tpv3SetTunnelCookies) GetMessageName() string {
	return "l2tpv3_set_tunnel_cookies"
}
func (*L2tpv3SetTunnelCookies) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*L2tpv3SetTunnelCookies) GetCrcString() string {
	return "3e783c95"
}
func NewL2tpv3SetTunnelCookies() api.Message {
	return &L2tpv3SetTunnelCookies{}
}

// L2tpv3SetTunnelCookiesReply represents the VPP binary API message 'l2tpv3_set_tunnel_cookies_reply'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 37:
//
//        ["l2tpv3_set_tunnel_cookies_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x0219718b"}
//        ],
//
type L2tpv3SetTunnelCookiesReply struct {
	Retval int32
}

func (*L2tpv3SetTunnelCookiesReply) GetMessageName() string {
	return "l2tpv3_set_tunnel_cookies_reply"
}
func (*L2tpv3SetTunnelCookiesReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*L2tpv3SetTunnelCookiesReply) GetCrcString() string {
	return "0219718b"
}
func NewL2tpv3SetTunnelCookiesReply() api.Message {
	return &L2tpv3SetTunnelCookiesReply{}
}

// SwIfL2tpv3TunnelDetails represents the VPP binary API message 'sw_if_l2tpv3_tunnel_details'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 43:
//
//        ["sw_if_l2tpv3_tunnel_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "interface_name", 64],
//            ["u8", "client_address", 16],
//            ["u8", "our_address", 16],
//            ["u32", "local_session_id"],
//            ["u32", "remote_session_id"],
//            ["u64", "local_cookie", 2],
//            ["u64", "remote_cookie"],
//            ["u8", "l2_sublayer_present"],
//            {"crc" : "0x6e3e43d4"}
//        ],
//
type SwIfL2tpv3TunnelDetails struct {
	SwIfIndex         uint32
	InterfaceName     []byte `struc:"[64]byte"`
	ClientAddress     []byte `struc:"[16]byte"`
	OurAddress        []byte `struc:"[16]byte"`
	LocalSessionID    uint32
	RemoteSessionID   uint32
	LocalCookie       []uint64 `struc:"[2]uint64"`
	RemoteCookie      uint64
	L2SublayerPresent uint8
}

func (*SwIfL2tpv3TunnelDetails) GetMessageName() string {
	return "sw_if_l2tpv3_tunnel_details"
}
func (*SwIfL2tpv3TunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwIfL2tpv3TunnelDetails) GetCrcString() string {
	return "6e3e43d4"
}
func NewSwIfL2tpv3TunnelDetails() api.Message {
	return &SwIfL2tpv3TunnelDetails{}
}

// SwIfL2tpv3TunnelDump represents the VPP binary API message 'sw_if_l2tpv3_tunnel_dump'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 57:
//
//        ["sw_if_l2tpv3_tunnel_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x597e7092"}
//        ],
//
type SwIfL2tpv3TunnelDump struct {
}

func (*SwIfL2tpv3TunnelDump) GetMessageName() string {
	return "sw_if_l2tpv3_tunnel_dump"
}
func (*SwIfL2tpv3TunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwIfL2tpv3TunnelDump) GetCrcString() string {
	return "597e7092"
}
func NewSwIfL2tpv3TunnelDump() api.Message {
	return &SwIfL2tpv3TunnelDump{}
}

// L2tpv3InterfaceEnableDisable represents the VPP binary API message 'l2tpv3_interface_enable_disable'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 63:
//
//        ["l2tpv3_interface_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "enable_disable"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xc74900bf"}
//        ],
//
type L2tpv3InterfaceEnableDisable struct {
	EnableDisable uint8
	SwIfIndex     uint32
}

func (*L2tpv3InterfaceEnableDisable) GetMessageName() string {
	return "l2tpv3_interface_enable_disable"
}
func (*L2tpv3InterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*L2tpv3InterfaceEnableDisable) GetCrcString() string {
	return "c74900bf"
}
func NewL2tpv3InterfaceEnableDisable() api.Message {
	return &L2tpv3InterfaceEnableDisable{}
}

// L2tpv3InterfaceEnableDisableReply represents the VPP binary API message 'l2tpv3_interface_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 71:
//
//        ["l2tpv3_interface_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x541a4157"}
//        ],
//
type L2tpv3InterfaceEnableDisableReply struct {
	Retval int32
}

func (*L2tpv3InterfaceEnableDisableReply) GetMessageName() string {
	return "l2tpv3_interface_enable_disable_reply"
}
func (*L2tpv3InterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*L2tpv3InterfaceEnableDisableReply) GetCrcString() string {
	return "541a4157"
}
func NewL2tpv3InterfaceEnableDisableReply() api.Message {
	return &L2tpv3InterfaceEnableDisableReply{}
}

// L2tpv3SetLookupKey represents the VPP binary API message 'l2tpv3_set_lookup_key'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 77:
//
//        ["l2tpv3_set_lookup_key",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "key"],
//            {"crc" : "0xb7152584"}
//        ],
//
type L2tpv3SetLookupKey struct {
	Key uint8
}

func (*L2tpv3SetLookupKey) GetMessageName() string {
	return "l2tpv3_set_lookup_key"
}
func (*L2tpv3SetLookupKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*L2tpv3SetLookupKey) GetCrcString() string {
	return "b7152584"
}
func NewL2tpv3SetLookupKey() api.Message {
	return &L2tpv3SetLookupKey{}
}

// L2tpv3SetLookupKeyReply represents the VPP binary API message 'l2tpv3_set_lookup_key_reply'.
// Generated from '/usr/share/vpp/api//l2tp.api.json', line 84:
//
//        ["l2tpv3_set_lookup_key_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x300e69f4"}
//        ]
//
type L2tpv3SetLookupKeyReply struct {
	Retval int32
}

func (*L2tpv3SetLookupKeyReply) GetMessageName() string {
	return "l2tpv3_set_lookup_key_reply"
}
func (*L2tpv3SetLookupKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*L2tpv3SetLookupKeyReply) GetCrcString() string {
	return "300e69f4"
}
func NewL2tpv3SetLookupKeyReply() api.Message {
	return &L2tpv3SetLookupKeyReply{}
}
