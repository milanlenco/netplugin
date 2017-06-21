// Package flowperpkt represents the VPP binary API of the 'flowperpkt' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//flowperpkt.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package flowperpkt

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x8418509d

// FlowperpktTxInterfaceAddDel represents the VPP binary API message 'flowperpkt_tx_interface_add_del'.
// Generated from '/usr/share/vpp/api//flowperpkt.api.json', line 6:
//
//        ["flowperpkt_tx_interface_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "which"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xef38c58c"}
//        ],
//
type FlowperpktTxInterfaceAddDel struct {
	IsAdd     uint8
	Which     uint8
	SwIfIndex uint32
}

func (*FlowperpktTxInterfaceAddDel) GetMessageName() string {
	return "flowperpkt_tx_interface_add_del"
}
func (*FlowperpktTxInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*FlowperpktTxInterfaceAddDel) GetCrcString() string {
	return "ef38c58c"
}
func NewFlowperpktTxInterfaceAddDel() api.Message {
	return &FlowperpktTxInterfaceAddDel{}
}

// FlowperpktTxInterfaceAddDelReply represents the VPP binary API message 'flowperpkt_tx_interface_add_del_reply'.
// Generated from '/usr/share/vpp/api//flowperpkt.api.json', line 15:
//
//        ["flowperpkt_tx_interface_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd47e6e0b"}
//        ],
//
type FlowperpktTxInterfaceAddDelReply struct {
	Retval int32
}

func (*FlowperpktTxInterfaceAddDelReply) GetMessageName() string {
	return "flowperpkt_tx_interface_add_del_reply"
}
func (*FlowperpktTxInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*FlowperpktTxInterfaceAddDelReply) GetCrcString() string {
	return "d47e6e0b"
}
func NewFlowperpktTxInterfaceAddDelReply() api.Message {
	return &FlowperpktTxInterfaceAddDelReply{}
}

// FlowperpktParams represents the VPP binary API message 'flowperpkt_params'.
// Generated from '/usr/share/vpp/api//flowperpkt.api.json', line 21:
//
//        ["flowperpkt_params",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "record_l2"],
//            ["u8", "record_l3"],
//            ["u8", "record_l4"],
//            ["u32", "active_timer"],
//            ["u32", "passive_timer"],
//            {"crc" : "0x738b8459"}
//        ],
//
type FlowperpktParams struct {
	RecordL2     uint8
	RecordL3     uint8
	RecordL4     uint8
	ActiveTimer  uint32
	PassiveTimer uint32
}

func (*FlowperpktParams) GetMessageName() string {
	return "flowperpkt_params"
}
func (*FlowperpktParams) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*FlowperpktParams) GetCrcString() string {
	return "738b8459"
}
func NewFlowperpktParams() api.Message {
	return &FlowperpktParams{}
}

// FlowperpktParamsReply represents the VPP binary API message 'flowperpkt_params_reply'.
// Generated from '/usr/share/vpp/api//flowperpkt.api.json', line 32:
//
//        ["flowperpkt_params_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x547ad509"}
//        ]
//
type FlowperpktParamsReply struct {
	Retval int32
}

func (*FlowperpktParamsReply) GetMessageName() string {
	return "flowperpkt_params_reply"
}
func (*FlowperpktParamsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*FlowperpktParamsReply) GetCrcString() string {
	return "547ad509"
}
func NewFlowperpktParamsReply() api.Message {
	return &FlowperpktParamsReply{}
}
