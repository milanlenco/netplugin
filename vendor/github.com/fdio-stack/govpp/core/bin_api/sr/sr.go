// Package sr represents the VPP binary API of the 'sr' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//sr.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package sr

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x430b3265

// SrLocalsidAddDel represents the VPP binary API message 'sr_localsid_add_del'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 6:
//
//        ["sr_localsid_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_del"],
//            ["u8", "localsid_addr", 16],
//            ["u8", "end_psp"],
//            ["u8", "behavior"],
//            ["u32", "sw_if_index"],
//            ["u32", "vlan_index"],
//            ["u32", "fib_table"],
//            ["u8", "nh_addr", 16],
//            {"crc" : "0xccf1dfb7"}
//        ],
//
type SrLocalsidAddDel struct {
	IsDel        uint8
	LocalsidAddr []byte `struc:"[16]byte"`
	EndPsp       uint8
	Behavior     uint8
	SwIfIndex    uint32
	VlanIndex    uint32
	FibTable     uint32
	NhAddr       []byte `struc:"[16]byte"`
}

func (*SrLocalsidAddDel) GetMessageName() string {
	return "sr_localsid_add_del"
}
func (*SrLocalsidAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrLocalsidAddDel) GetCrcString() string {
	return "ccf1dfb7"
}
func NewSrLocalsidAddDel() api.Message {
	return &SrLocalsidAddDel{}
}

// SrLocalsidAddDelReply represents the VPP binary API message 'sr_localsid_add_del_reply'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 20:
//
//        ["sr_localsid_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x8ed0f725"}
//        ],
//
type SrLocalsidAddDelReply struct {
	Retval int32
}

func (*SrLocalsidAddDelReply) GetMessageName() string {
	return "sr_localsid_add_del_reply"
}
func (*SrLocalsidAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrLocalsidAddDelReply) GetCrcString() string {
	return "8ed0f725"
}
func NewSrLocalsidAddDelReply() api.Message {
	return &SrLocalsidAddDelReply{}
}

// SrPolicyAdd represents the VPP binary API message 'sr_policy_add'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 26:
//
//        ["sr_policy_add",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "bsid_addr", 16],
//            ["u32", "weight"],
//            ["u8", "is_encap"],
//            ["u8", "type"],
//            ["u32", "fib_table"],
//            ["u8", "n_segments"],
//            ["u8", "segments", 0],
//            {"crc" : "0xd2c9c4a7"}
//        ],
//
type SrPolicyAdd struct {
	BsidAddr  []byte `struc:"[16]byte"`
	Weight    uint32
	IsEncap   uint8
	Type      uint8
	FibTable  uint32
	NSegments uint8
	Segments  []byte
}

func (*SrPolicyAdd) GetMessageName() string {
	return "sr_policy_add"
}
func (*SrPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyAdd) GetCrcString() string {
	return "d2c9c4a7"
}
func NewSrPolicyAdd() api.Message {
	return &SrPolicyAdd{}
}

// SrPolicyAddReply represents the VPP binary API message 'sr_policy_add_reply'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 39:
//
//        ["sr_policy_add_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x1fa4bf69"}
//        ],
//
type SrPolicyAddReply struct {
	Retval int32
}

func (*SrPolicyAddReply) GetMessageName() string {
	return "sr_policy_add_reply"
}
func (*SrPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyAddReply) GetCrcString() string {
	return "1fa4bf69"
}
func NewSrPolicyAddReply() api.Message {
	return &SrPolicyAddReply{}
}

// SrPolicyMod represents the VPP binary API message 'sr_policy_mod'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 45:
//
//        ["sr_policy_mod",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "bsid_addr", 16],
//            ["u32", "sr_policy_index"],
//            ["u32", "fib_table"],
//            ["u8", "operation"],
//            ["u32", "sl_index"],
//            ["u32", "weight"],
//            ["u8", "n_segments"],
//            ["u8", "segments", 0],
//            {"crc" : "0x66016cd3"}
//        ],
//
type SrPolicyMod struct {
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	FibTable      uint32
	Operation     uint8
	SlIndex       uint32
	Weight        uint32
	NSegments     uint8
	Segments      []byte
}

func (*SrPolicyMod) GetMessageName() string {
	return "sr_policy_mod"
}
func (*SrPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyMod) GetCrcString() string {
	return "66016cd3"
}
func NewSrPolicyMod() api.Message {
	return &SrPolicyMod{}
}

// SrPolicyModReply represents the VPP binary API message 'sr_policy_mod_reply'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 59:
//
//        ["sr_policy_mod_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xe53e8c2a"}
//        ],
//
type SrPolicyModReply struct {
	Retval int32
}

func (*SrPolicyModReply) GetMessageName() string {
	return "sr_policy_mod_reply"
}
func (*SrPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyModReply) GetCrcString() string {
	return "e53e8c2a"
}
func NewSrPolicyModReply() api.Message {
	return &SrPolicyModReply{}
}

// SrPolicyDel represents the VPP binary API message 'sr_policy_del'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 65:
//
//        ["sr_policy_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "bsid_addr", 16],
//            ["u32", "sr_policy_index"],
//            {"crc" : "0xcc9e015e"}
//        ],
//
type SrPolicyDel struct {
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
}

func (*SrPolicyDel) GetMessageName() string {
	return "sr_policy_del"
}
func (*SrPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyDel) GetCrcString() string {
	return "cc9e015e"
}
func NewSrPolicyDel() api.Message {
	return &SrPolicyDel{}
}

// SrPolicyDelReply represents the VPP binary API message 'sr_policy_del_reply'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 73:
//
//        ["sr_policy_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x3315338b"}
//        ],
//
type SrPolicyDelReply struct {
	Retval int32
}

func (*SrPolicyDelReply) GetMessageName() string {
	return "sr_policy_del_reply"
}
func (*SrPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyDelReply) GetCrcString() string {
	return "3315338b"
}
func NewSrPolicyDelReply() api.Message {
	return &SrPolicyDelReply{}
}

// SrSteeringAddDel represents the VPP binary API message 'sr_steering_add_del'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 79:
//
//        ["sr_steering_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_del"],
//            ["u8", "bsid_addr", 16],
//            ["u32", "sr_policy_index"],
//            ["u32", "table_id"],
//            ["u8", "prefix_addr", 16],
//            ["u32", "mask_width"],
//            ["u32", "sw_if_index"],
//            ["u8", "traffic_type"],
//            {"crc" : "0x61da4dae"}
//        ],
//
type SrSteeringAddDel struct {
	IsDel         uint8
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	TableID       uint32
	PrefixAddr    []byte `struc:"[16]byte"`
	MaskWidth     uint32
	SwIfIndex     uint32
	TrafficType   uint8
}

func (*SrSteeringAddDel) GetMessageName() string {
	return "sr_steering_add_del"
}
func (*SrSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrSteeringAddDel) GetCrcString() string {
	return "61da4dae"
}
func NewSrSteeringAddDel() api.Message {
	return &SrSteeringAddDel{}
}

// SrSteeringAddDelReply represents the VPP binary API message 'sr_steering_add_del_reply'.
// Generated from '/usr/share/vpp/api//sr.api.json', line 93:
//
//        ["sr_steering_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x8461b882"}
//        ]
//
type SrSteeringAddDelReply struct {
	Retval int32
}

func (*SrSteeringAddDelReply) GetMessageName() string {
	return "sr_steering_add_del_reply"
}
func (*SrSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrSteeringAddDelReply) GetCrcString() string {
	return "8461b882"
}
func NewSrSteeringAddDelReply() api.Message {
	return &SrSteeringAddDelReply{}
}
