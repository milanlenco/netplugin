// Package ioam_export represents the VPP binary API of the 'ioam_export' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//ioam_export.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package ioam_export

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x230fca5f

// IoamExportIP6EnableDisable represents the VPP binary API message 'ioam_export_ip6_enable_disable'.
// Generated from '/usr/share/vpp/api//ioam_export.api.json', line 6:
//
//        ["ioam_export_ip6_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_disable"],
//            ["u8", "collector_address", 4],
//            ["u8", "src_address", 4],
//            {"crc" : "0x8b988966"}
//        ],
//
type IoamExportIP6EnableDisable struct {
	IsDisable        uint8
	CollectorAddress []byte `struc:"[4]byte"`
	SrcAddress       []byte `struc:"[4]byte"`
}

func (*IoamExportIP6EnableDisable) GetMessageName() string {
	return "ioam_export_ip6_enable_disable"
}
func (*IoamExportIP6EnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IoamExportIP6EnableDisable) GetCrcString() string {
	return "8b988966"
}
func NewIoamExportIP6EnableDisable() api.Message {
	return &IoamExportIP6EnableDisable{}
}

// IoamExportIP6EnableDisableReply represents the VPP binary API message 'ioam_export_ip6_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//ioam_export.api.json', line 15:
//
//        ["ioam_export_ip6_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa49763f0"}
//        ]
//
type IoamExportIP6EnableDisableReply struct {
	Retval int32
}

func (*IoamExportIP6EnableDisableReply) GetMessageName() string {
	return "ioam_export_ip6_enable_disable_reply"
}
func (*IoamExportIP6EnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IoamExportIP6EnableDisableReply) GetCrcString() string {
	return "a49763f0"
}
func NewIoamExportIP6EnableDisableReply() api.Message {
	return &IoamExportIP6EnableDisableReply{}
}
