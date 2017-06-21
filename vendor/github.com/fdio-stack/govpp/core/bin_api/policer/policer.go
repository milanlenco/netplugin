// Package policer represents the VPP binary API of the 'policer' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//policer.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package policer

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x791375b8

// PolicerAddDel represents the VPP binary API message 'policer_add_del'.
// Generated from '/usr/share/vpp/api//policer.api.json', line 6:
//
//        ["policer_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "name", 64],
//            ["u32", "cir"],
//            ["u32", "eir"],
//            ["u64", "cb"],
//            ["u64", "eb"],
//            ["u8", "rate_type"],
//            ["u8", "round_type"],
//            ["u8", "type"],
//            ["u8", "color_aware"],
//            ["u8", "conform_action_type"],
//            ["u8", "conform_dscp"],
//            ["u8", "exceed_action_type"],
//            ["u8", "exceed_dscp"],
//            ["u8", "violate_action_type"],
//            ["u8", "violate_dscp"],
//            {"crc" : "0xe1bba755"}
//        ],
//
type PolicerAddDel struct {
	IsAdd             uint8
	Name              []byte `struc:"[64]byte"`
	Cir               uint32
	Eir               uint32
	Cb                uint64
	Eb                uint64
	RateType          uint8
	RoundType         uint8
	Type              uint8
	ColorAware        uint8
	ConformActionType uint8
	ConformDscp       uint8
	ExceedActionType  uint8
	ExceedDscp        uint8
	ViolateActionType uint8
	ViolateDscp       uint8
}

func (*PolicerAddDel) GetMessageName() string {
	return "policer_add_del"
}
func (*PolicerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PolicerAddDel) GetCrcString() string {
	return "e1bba755"
}
func NewPolicerAddDel() api.Message {
	return &PolicerAddDel{}
}

// PolicerAddDelReply represents the VPP binary API message 'policer_add_del_reply'.
// Generated from '/usr/share/vpp/api//policer.api.json', line 28:
//
//        ["policer_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "policer_index"],
//            {"crc" : "0xddb244b0"}
//        ],
//
type PolicerAddDelReply struct {
	Retval       int32
	PolicerIndex uint32
}

func (*PolicerAddDelReply) GetMessageName() string {
	return "policer_add_del_reply"
}
func (*PolicerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PolicerAddDelReply) GetCrcString() string {
	return "ddb244b0"
}
func NewPolicerAddDelReply() api.Message {
	return &PolicerAddDelReply{}
}

// PolicerDump represents the VPP binary API message 'policer_dump'.
// Generated from '/usr/share/vpp/api//policer.api.json', line 35:
//
//        ["policer_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "match_name_valid"],
//            ["u8", "match_name", 64],
//            {"crc" : "0x526b205b"}
//        ],
//
type PolicerDump struct {
	MatchNameValid uint8
	MatchName      []byte `struc:"[64]byte"`
}

func (*PolicerDump) GetMessageName() string {
	return "policer_dump"
}
func (*PolicerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PolicerDump) GetCrcString() string {
	return "526b205b"
}
func NewPolicerDump() api.Message {
	return &PolicerDump{}
}

// PolicerDetails represents the VPP binary API message 'policer_details'.
// Generated from '/usr/share/vpp/api//policer.api.json', line 43:
//
//        ["policer_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u32", "cir"],
//            ["u32", "eir"],
//            ["u64", "cb"],
//            ["u64", "eb"],
//            ["u8", "rate_type"],
//            ["u8", "round_type"],
//            ["u8", "type"],
//            ["u8", "conform_action_type"],
//            ["u8", "conform_dscp"],
//            ["u8", "exceed_action_type"],
//            ["u8", "exceed_dscp"],
//            ["u8", "violate_action_type"],
//            ["u8", "violate_dscp"],
//            ["u8", "single_rate"],
//            ["u8", "color_aware"],
//            ["u32", "scale"],
//            ["u32", "cir_tokens_per_period"],
//            ["u32", "pir_tokens_per_period"],
//            ["u32", "current_limit"],
//            ["u32", "current_bucket"],
//            ["u32", "extended_limit"],
//            ["u32", "extended_bucket"],
//            ["u64", "last_update_time"],
//            {"crc" : "0xa9729913"}
//        ]
//
type PolicerDetails struct {
	Name               []byte `struc:"[64]byte"`
	Cir                uint32
	Eir                uint32
	Cb                 uint64
	Eb                 uint64
	RateType           uint8
	RoundType          uint8
	Type               uint8
	ConformActionType  uint8
	ConformDscp        uint8
	ExceedActionType   uint8
	ExceedDscp         uint8
	ViolateActionType  uint8
	ViolateDscp        uint8
	SingleRate         uint8
	ColorAware         uint8
	Scale              uint32
	CirTokensPerPeriod uint32
	PirTokensPerPeriod uint32
	CurrentLimit       uint32
	CurrentBucket      uint32
	ExtendedLimit      uint32
	ExtendedBucket     uint32
	LastUpdateTime     uint64
}

func (*PolicerDetails) GetMessageName() string {
	return "policer_details"
}
func (*PolicerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PolicerDetails) GetCrcString() string {
	return "a9729913"
}
func NewPolicerDetails() api.Message {
	return &PolicerDetails{}
}
