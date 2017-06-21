// Package lb represents the VPP binary API of the 'lb' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//lb.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package lb

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xaddf1b29

// LbConf represents the VPP binary API message 'lb_conf'.
// Generated from '/usr/share/vpp/api//lb.api.json', line 6:
//
//        ["lb_conf",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "ip4_src_address"],
//            ["u8", "ip6_src_address", 16],
//            ["u32", "sticky_buckets_per_core"],
//            ["u32", "flow_timeout"],
//            {"crc" : "0x3e9c476c"}
//        ],
//
type LbConf struct {
	IP4SrcAddress        uint32
	IP6SrcAddress        []byte `struc:"[16]byte"`
	StickyBucketsPerCore uint32
	FlowTimeout          uint32
}

func (*LbConf) GetMessageName() string {
	return "lb_conf"
}
func (*LbConf) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LbConf) GetCrcString() string {
	return "3e9c476c"
}
func NewLbConf() api.Message {
	return &LbConf{}
}

// LbConfReply represents the VPP binary API message 'lb_conf_reply'.
// Generated from '/usr/share/vpp/api//lb.api.json', line 16:
//
//        ["lb_conf_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb6b1c34e"}
//        ],
//
type LbConfReply struct {
	Retval int32
}

func (*LbConfReply) GetMessageName() string {
	return "lb_conf_reply"
}
func (*LbConfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LbConfReply) GetCrcString() string {
	return "b6b1c34e"
}
func NewLbConfReply() api.Message {
	return &LbConfReply{}
}

// LbAddDelVip represents the VPP binary API message 'lb_add_del_vip'.
// Generated from '/usr/share/vpp/api//lb.api.json', line 22:
//
//        ["lb_add_del_vip",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "ip_prefix", 16],
//            ["u8", "prefix_length"],
//            ["u8", "is_gre4"],
//            ["u32", "new_flows_table_length"],
//            ["u8", "is_del"],
//            {"crc" : "0xf611e4a4"}
//        ],
//
type LbAddDelVip struct {
	IPPrefix            []byte `struc:"[16]byte"`
	PrefixLength        uint8
	IsGre4              uint8
	NewFlowsTableLength uint32
	IsDel               uint8
}

func (*LbAddDelVip) GetMessageName() string {
	return "lb_add_del_vip"
}
func (*LbAddDelVip) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LbAddDelVip) GetCrcString() string {
	return "f611e4a4"
}
func NewLbAddDelVip() api.Message {
	return &LbAddDelVip{}
}

// LbAddDelVipReply represents the VPP binary API message 'lb_add_del_vip_reply'.
// Generated from '/usr/share/vpp/api//lb.api.json', line 33:
//
//        ["lb_add_del_vip_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x446cc0f6"}
//        ],
//
type LbAddDelVipReply struct {
	Retval int32
}

func (*LbAddDelVipReply) GetMessageName() string {
	return "lb_add_del_vip_reply"
}
func (*LbAddDelVipReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LbAddDelVipReply) GetCrcString() string {
	return "446cc0f6"
}
func NewLbAddDelVipReply() api.Message {
	return &LbAddDelVipReply{}
}

// LbAddDelAs represents the VPP binary API message 'lb_add_del_as'.
// Generated from '/usr/share/vpp/api//lb.api.json', line 39:
//
//        ["lb_add_del_as",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "vip_ip_prefix", 16],
//            ["u8", "vip_prefix_length"],
//            ["u8", "as_address", 16],
//            ["u8", "is_del"],
//            {"crc" : "0xa74233e4"}
//        ],
//
type LbAddDelAs struct {
	VipIPPrefix     []byte `struc:"[16]byte"`
	VipPrefixLength uint8
	AsAddress       []byte `struc:"[16]byte"`
	IsDel           uint8
}

func (*LbAddDelAs) GetMessageName() string {
	return "lb_add_del_as"
}
func (*LbAddDelAs) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LbAddDelAs) GetCrcString() string {
	return "a74233e4"
}
func NewLbAddDelAs() api.Message {
	return &LbAddDelAs{}
}

// LbAddDelAsReply represents the VPP binary API message 'lb_add_del_as_reply'.
// Generated from '/usr/share/vpp/api//lb.api.json', line 49:
//
//        ["lb_add_del_as_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x6c407528"}
//        ]
//
type LbAddDelAsReply struct {
	Retval int32
}

func (*LbAddDelAsReply) GetMessageName() string {
	return "lb_add_del_as_reply"
}
func (*LbAddDelAsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LbAddDelAsReply) GetCrcString() string {
	return "6c407528"
}
func NewLbAddDelAsReply() api.Message {
	return &LbAddDelAsReply{}
}
