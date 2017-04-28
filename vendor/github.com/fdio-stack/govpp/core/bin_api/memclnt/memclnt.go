// Package memclnt represents the VPP binary API of the 'memclnt' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//memclnt.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package memclnt

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x6408124c

// MemclntCreate represents the VPP binary API message 'memclnt_create'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 6:
//
//        ["memclnt_create",
//            ["u16", "_vl_msg_id"],
//            ["i32", "ctx_quota"],
//            ["u32", "context"],
//            ["u64", "input_queue"],
//            ["u8", "name", 64],
//            ["u32", "api_versions", 8],
//            {"crc" : "0x2671bcfd"}
//        ],
//
type MemclntCreate struct {
	CtxQuota    int32
	Context     uint32
	InputQueue  uint64
	Name        []byte   `struc:"[64]byte"`
	APIVersions []uint32 `struc:"[8]uint32"`
}

func (*MemclntCreate) GetMessageName() string {
	return "memclnt_create"
}
func (*MemclntCreate) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*MemclntCreate) GetCrcString() string {
	return "2671bcfd"
}
func NewMemclntCreate() api.Message {
	return &MemclntCreate{}
}

// MemclntCreateReply represents the VPP binary API message 'memclnt_create_reply'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 15:
//
//        ["memclnt_create_reply",
//            ["u16", "_vl_msg_id"],
//            ["i32", "response"],
//            ["u64", "handle"],
//            ["u32", "index"],
//            ["u32", "context"],
//            ["u64", "message_table"],
//            {"crc" : "0xf71a8a1a"}
//        ],
//
type MemclntCreateReply struct {
	Response     int32
	Handle       uint64
	Index        uint32
	Context      uint32
	MessageTable uint64
}

func (*MemclntCreateReply) GetMessageName() string {
	return "memclnt_create_reply"
}
func (*MemclntCreateReply) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*MemclntCreateReply) GetCrcString() string {
	return "f71a8a1a"
}
func NewMemclntCreateReply() api.Message {
	return &MemclntCreateReply{}
}

// MemclntDelete represents the VPP binary API message 'memclnt_delete'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 24:
//
//        ["memclnt_delete",
//            ["u16", "_vl_msg_id"],
//            ["u32", "index"],
//            ["u64", "handle"],
//            {"crc" : "0x602f4d82"}
//        ],
//
type MemclntDelete struct {
	Index  uint32
	Handle uint64
}

func (*MemclntDelete) GetMessageName() string {
	return "memclnt_delete"
}
func (*MemclntDelete) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*MemclntDelete) GetCrcString() string {
	return "602f4d82"
}
func NewMemclntDelete() api.Message {
	return &MemclntDelete{}
}

// MemclntDeleteReply represents the VPP binary API message 'memclnt_delete_reply'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 30:
//
//        ["memclnt_delete_reply",
//            ["u16", "_vl_msg_id"],
//            ["i32", "response"],
//            ["u64", "handle"],
//            {"crc" : "0x587855a7"}
//        ],
//
type MemclntDeleteReply struct {
	Response int32
	Handle   uint64
}

func (*MemclntDeleteReply) GetMessageName() string {
	return "memclnt_delete_reply"
}
func (*MemclntDeleteReply) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*MemclntDeleteReply) GetCrcString() string {
	return "587855a7"
}
func NewMemclntDeleteReply() api.Message {
	return &MemclntDeleteReply{}
}

// RxThreadExit represents the VPP binary API message 'rx_thread_exit'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 36:
//
//        ["rx_thread_exit",
//            ["u16", "_vl_msg_id"],
//            ["u8", "dummy"],
//            {"crc" : "0x6110e464"}
//        ],
//
type RxThreadExit struct {
	Dummy uint8
}

func (*RxThreadExit) GetMessageName() string {
	return "rx_thread_exit"
}
func (*RxThreadExit) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*RxThreadExit) GetCrcString() string {
	return "6110e464"
}
func NewRxThreadExit() api.Message {
	return &RxThreadExit{}
}

// MemclntRxThreadSuspend represents the VPP binary API message 'memclnt_rx_thread_suspend'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 41:
//
//        ["memclnt_rx_thread_suspend",
//            ["u16", "_vl_msg_id"],
//            ["u8", "dummy"],
//            {"crc" : "0x98c139f3"}
//        ],
//
type MemclntRxThreadSuspend struct {
	Dummy uint8
}

func (*MemclntRxThreadSuspend) GetMessageName() string {
	return "memclnt_rx_thread_suspend"
}
func (*MemclntRxThreadSuspend) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*MemclntRxThreadSuspend) GetCrcString() string {
	return "98c139f3"
}
func NewMemclntRxThreadSuspend() api.Message {
	return &MemclntRxThreadSuspend{}
}

// MemclntReadTimeout represents the VPP binary API message 'memclnt_read_timeout'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 46:
//
//        ["memclnt_read_timeout",
//            ["u16", "_vl_msg_id"],
//            ["u8", "dummy"],
//            {"crc" : "0x8161e828"}
//        ],
//
type MemclntReadTimeout struct {
	Dummy uint8
}

func (*MemclntReadTimeout) GetMessageName() string {
	return "memclnt_read_timeout"
}
func (*MemclntReadTimeout) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*MemclntReadTimeout) GetCrcString() string {
	return "8161e828"
}
func NewMemclntReadTimeout() api.Message {
	return &MemclntReadTimeout{}
}

// RPCCall represents the VPP binary API message 'rpc_call'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 51:
//
//        ["rpc_call",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u64", "function"],
//            ["u8", "multicast"],
//            ["u8", "need_barrier_sync"],
//            ["u8", "send_reply"],
//            ["u8", "data", 0],
//            {"crc" : "0xe17d6c23"}
//        ],
//
type RPCCall struct {
	Function        uint64
	Multicast       uint8
	NeedBarrierSync uint8
	SendReply       uint8
	Data            []byte
}

func (*RPCCall) GetMessageName() string {
	return "rpc_call"
}
func (*RPCCall) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*RPCCall) GetCrcString() string {
	return "e17d6c23"
}
func NewRPCCall() api.Message {
	return &RPCCall{}
}

// RPCCallReply represents the VPP binary API message 'rpc_call_reply'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 62:
//
//        ["rpc_call_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x45621c6a"}
//        ],
//
type RPCCallReply struct {
	Retval int32
}

func (*RPCCallReply) GetMessageName() string {
	return "rpc_call_reply"
}
func (*RPCCallReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*RPCCallReply) GetCrcString() string {
	return "45621c6a"
}
func NewRPCCallReply() api.Message {
	return &RPCCallReply{}
}

// GetFirstMsgID represents the VPP binary API message 'get_first_msg_id'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 68:
//
//        ["get_first_msg_id",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            {"crc" : "0x56f7fd40"}
//        ],
//
type GetFirstMsgID struct {
	Name []byte `struc:"[64]byte"`
}

func (*GetFirstMsgID) GetMessageName() string {
	return "get_first_msg_id"
}
func (*GetFirstMsgID) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GetFirstMsgID) GetCrcString() string {
	return "56f7fd40"
}
func NewGetFirstMsgID() api.Message {
	return &GetFirstMsgID{}
}

// GetFirstMsgIDReply represents the VPP binary API message 'get_first_msg_id_reply'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 75:
//
//        ["get_first_msg_id_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u16", "first_msg_id"],
//            {"crc" : "0x3c6931c6"}
//        ],
//
type GetFirstMsgIDReply struct {
	Retval     int32
	FirstMsgID uint16
}

func (*GetFirstMsgIDReply) GetMessageName() string {
	return "get_first_msg_id_reply"
}
func (*GetFirstMsgIDReply) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GetFirstMsgIDReply) GetCrcString() string {
	return "3c6931c6"
}
func NewGetFirstMsgIDReply() api.Message {
	return &GetFirstMsgIDReply{}
}

// TracePluginMsgIds represents the VPP binary API message 'trace_plugin_msg_ids'.
// Generated from '/usr/share/vpp/api//memclnt.api.json', line 83:
//
//        ["trace_plugin_msg_ids",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "plugin_name", 128],
//            ["u16", "first_msg_id"],
//            ["u16", "last_msg_id"],
//            {"crc" : "0x12ce6ba5"}
//        ]
//
type TracePluginMsgIds struct {
	PluginName []byte `struc:"[128]byte"`
	FirstMsgID uint16
	LastMsgID  uint16
}

func (*TracePluginMsgIds) GetMessageName() string {
	return "trace_plugin_msg_ids"
}
func (*TracePluginMsgIds) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TracePluginMsgIds) GetCrcString() string {
	return "12ce6ba5"
}
func NewTracePluginMsgIds() api.Message {
	return &TracePluginMsgIds{}
}
