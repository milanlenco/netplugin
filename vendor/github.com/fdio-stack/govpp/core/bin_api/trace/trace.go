// Package trace represents the VPP binary API of the 'trace' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//trace.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package trace

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xe5d0d932

// TraceProfileAdd represents the VPP binary API message 'trace_profile_add'.
// Generated from '/usr/share/vpp/api//trace.api.json', line 6:
//
//        ["trace_profile_add",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "trace_type"],
//            ["u8", "num_elts"],
//            ["u8", "trace_tsp"],
//            ["u32", "node_id"],
//            ["u32", "app_data"],
//            {"crc" : "0xc6300bd4"}
//        ],
//
type TraceProfileAdd struct {
	TraceType uint8
	NumElts   uint8
	TraceTsp  uint8
	NodeID    uint32
	AppData   uint32
}

func (*TraceProfileAdd) GetMessageName() string {
	return "trace_profile_add"
}
func (*TraceProfileAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TraceProfileAdd) GetCrcString() string {
	return "c6300bd4"
}
func NewTraceProfileAdd() api.Message {
	return &TraceProfileAdd{}
}

// TraceProfileAddReply represents the VPP binary API message 'trace_profile_add_reply'.
// Generated from '/usr/share/vpp/api//trace.api.json', line 17:
//
//        ["trace_profile_add_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xfcf91946"}
//        ],
//
type TraceProfileAddReply struct {
	Retval int32
}

func (*TraceProfileAddReply) GetMessageName() string {
	return "trace_profile_add_reply"
}
func (*TraceProfileAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TraceProfileAddReply) GetCrcString() string {
	return "fcf91946"
}
func NewTraceProfileAddReply() api.Message {
	return &TraceProfileAddReply{}
}

// TraceProfileDel represents the VPP binary API message 'trace_profile_del'.
// Generated from '/usr/share/vpp/api//trace.api.json', line 23:
//
//        ["trace_profile_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x94b1447f"}
//        ],
//
type TraceProfileDel struct {
}

func (*TraceProfileDel) GetMessageName() string {
	return "trace_profile_del"
}
func (*TraceProfileDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TraceProfileDel) GetCrcString() string {
	return "94b1447f"
}
func NewTraceProfileDel() api.Message {
	return &TraceProfileDel{}
}

// TraceProfileDelReply represents the VPP binary API message 'trace_profile_del_reply'.
// Generated from '/usr/share/vpp/api//trace.api.json', line 29:
//
//        ["trace_profile_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd04895a4"}
//        ],
//
type TraceProfileDelReply struct {
	Retval int32
}

func (*TraceProfileDelReply) GetMessageName() string {
	return "trace_profile_del_reply"
}
func (*TraceProfileDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TraceProfileDelReply) GetCrcString() string {
	return "d04895a4"
}
func NewTraceProfileDelReply() api.Message {
	return &TraceProfileDelReply{}
}

// TraceProfileShowConfig represents the VPP binary API message 'trace_profile_show_config'.
// Generated from '/usr/share/vpp/api//trace.api.json', line 35:
//
//        ["trace_profile_show_config",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x983f7b0c"}
//        ],
//
type TraceProfileShowConfig struct {
}

func (*TraceProfileShowConfig) GetMessageName() string {
	return "trace_profile_show_config"
}
func (*TraceProfileShowConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TraceProfileShowConfig) GetCrcString() string {
	return "983f7b0c"
}
func NewTraceProfileShowConfig() api.Message {
	return &TraceProfileShowConfig{}
}

// TraceProfileShowConfigReply represents the VPP binary API message 'trace_profile_show_config_reply'.
// Generated from '/usr/share/vpp/api//trace.api.json', line 41:
//
//        ["trace_profile_show_config_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "trace_type"],
//            ["u8", "num_elts"],
//            ["u8", "trace_tsp"],
//            ["u32", "node_id"],
//            ["u32", "app_data"],
//            {"crc" : "0x4f6f9bdd"}
//        ]
//
type TraceProfileShowConfigReply struct {
	Retval    int32
	TraceType uint8
	NumElts   uint8
	TraceTsp  uint8
	NodeID    uint32
	AppData   uint32
}

func (*TraceProfileShowConfigReply) GetMessageName() string {
	return "trace_profile_show_config_reply"
}
func (*TraceProfileShowConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TraceProfileShowConfigReply) GetCrcString() string {
	return "4f6f9bdd"
}
func NewTraceProfileShowConfigReply() api.Message {
	return &TraceProfileShowConfigReply{}
}
