// Package vxlan_gpe_ioam_export represents the VPP binary API of the 'vxlan_gpe_ioam_export' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//vxlan_gpe_ioam_export.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package vxlan_gpe_ioam_export

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xd0221e6a

// VxlanGpeIoamExportEnableDisable represents the VPP binary API message 'vxlan_gpe_ioam_export_enable_disable'.
// Generated from '/usr/share/vpp/api//vxlan_gpe_ioam_export.api.json', line 6:
//
//        ["vxlan_gpe_ioam_export_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_disable"],
//            ["u8", "collector_address", 4],
//            ["u8", "src_address", 4],
//            {"crc" : "0x20586df7"}
//        ],
//
type VxlanGpeIoamExportEnableDisable struct {
	IsDisable        uint8
	CollectorAddress []byte `struc:"[4]byte"`
	SrcAddress       []byte `struc:"[4]byte"`
}

func (*VxlanGpeIoamExportEnableDisable) GetMessageName() string {
	return "vxlan_gpe_ioam_export_enable_disable"
}
func (*VxlanGpeIoamExportEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeIoamExportEnableDisable) GetCrcString() string {
	return "20586df7"
}
func NewVxlanGpeIoamExportEnableDisable() api.Message {
	return &VxlanGpeIoamExportEnableDisable{}
}

// VxlanGpeIoamExportEnableDisableReply represents the VPP binary API message 'vxlan_gpe_ioam_export_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//vxlan_gpe_ioam_export.api.json', line 15:
//
//        ["vxlan_gpe_ioam_export_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x2baa825a"}
//        ]
//
type VxlanGpeIoamExportEnableDisableReply struct {
	Retval int32
}

func (*VxlanGpeIoamExportEnableDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_export_enable_disable_reply"
}
func (*VxlanGpeIoamExportEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeIoamExportEnableDisableReply) GetCrcString() string {
	return "2baa825a"
}
func NewVxlanGpeIoamExportEnableDisableReply() api.Message {
	return &VxlanGpeIoamExportEnableDisableReply{}
}
