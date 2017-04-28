// Package span represents the VPP binary API of the 'span' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//span.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package span

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x566d5000

// SwInterfaceSpanEnableDisable represents the VPP binary API message 'sw_interface_span_enable_disable'.
// Generated from '/usr/share/vpp/api//span.api.json', line 6:
//
//        ["sw_interface_span_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index_from"],
//            ["u32", "sw_if_index_to"],
//            ["u8", "state"],
//            {"crc" : "0x15a0b8ab"}
//        ],
//
type SwInterfaceSpanEnableDisable struct {
	SwIfIndexFrom uint32
	SwIfIndexTo   uint32
	State         uint8
}

func (*SwInterfaceSpanEnableDisable) GetMessageName() string {
	return "sw_interface_span_enable_disable"
}
func (*SwInterfaceSpanEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSpanEnableDisable) GetCrcString() string {
	return "15a0b8ab"
}
func NewSwInterfaceSpanEnableDisable() api.Message {
	return &SwInterfaceSpanEnableDisable{}
}

// SwInterfaceSpanEnableDisableReply represents the VPP binary API message 'sw_interface_span_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//span.api.json', line 15:
//
//        ["sw_interface_span_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5f317ff8"}
//        ],
//
type SwInterfaceSpanEnableDisableReply struct {
	Retval int32
}

func (*SwInterfaceSpanEnableDisableReply) GetMessageName() string {
	return "sw_interface_span_enable_disable_reply"
}
func (*SwInterfaceSpanEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSpanEnableDisableReply) GetCrcString() string {
	return "5f317ff8"
}
func NewSwInterfaceSpanEnableDisableReply() api.Message {
	return &SwInterfaceSpanEnableDisableReply{}
}

// SwInterfaceSpanDump represents the VPP binary API message 'sw_interface_span_dump'.
// Generated from '/usr/share/vpp/api//span.api.json', line 21:
//
//        ["sw_interface_span_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x63b587fe"}
//        ],
//
type SwInterfaceSpanDump struct {
}

func (*SwInterfaceSpanDump) GetMessageName() string {
	return "sw_interface_span_dump"
}
func (*SwInterfaceSpanDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSpanDump) GetCrcString() string {
	return "63b587fe"
}
func NewSwInterfaceSpanDump() api.Message {
	return &SwInterfaceSpanDump{}
}

// SwInterfaceSpanDetails represents the VPP binary API message 'sw_interface_span_details'.
// Generated from '/usr/share/vpp/api//span.api.json', line 27:
//
//        ["sw_interface_span_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index_from"],
//            ["u32", "sw_if_index_to"],
//            ["u8", "state"],
//            {"crc" : "0x7ebe522b"}
//        ]
//
type SwInterfaceSpanDetails struct {
	SwIfIndexFrom uint32
	SwIfIndexTo   uint32
	State         uint8
}

func (*SwInterfaceSpanDetails) GetMessageName() string {
	return "sw_interface_span_details"
}
func (*SwInterfaceSpanDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSpanDetails) GetCrcString() string {
	return "7ebe522b"
}
func NewSwInterfaceSpanDetails() api.Message {
	return &SwInterfaceSpanDetails{}
}
