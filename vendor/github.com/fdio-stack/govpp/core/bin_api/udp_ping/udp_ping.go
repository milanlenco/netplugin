// Package udp_ping represents the VPP binary API of the 'udp_ping' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//udp_ping.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package udp_ping

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xbfb052ef

// UDPPingAddDelReq represents the VPP binary API message 'udp_ping_add_del_req'.
// Generated from '/usr/share/vpp/api//udp_ping.api.json', line 6:
//
//        ["udp_ping_add_del_req",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "src_ip_address", 16],
//            ["u8", "dst_ip_address", 16],
//            ["u16", "start_src_port"],
//            ["u16", "end_src_port"],
//            ["u16", "start_dst_port"],
//            ["u16", "end_dst_port"],
//            ["u16", "interval"],
//            ["u8", "is_ipv4"],
//            ["u8", "dis"],
//            ["u8", "fault_det"],
//            ["u8", "reserve", 3],
//            {"crc" : "0xa7280e39"}
//        ],
//
type UDPPingAddDelReq struct {
	SrcIPAddress []byte `struc:"[16]byte"`
	DstIPAddress []byte `struc:"[16]byte"`
	StartSrcPort uint16
	EndSrcPort   uint16
	StartDstPort uint16
	EndDstPort   uint16
	Interval     uint16
	IsIpv4       uint8
	Dis          uint8
	FaultDet     uint8
	Reserve      []byte `struc:"[3]byte"`
}

func (*UDPPingAddDelReq) GetMessageName() string {
	return "udp_ping_add_del_req"
}
func (*UDPPingAddDelReq) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*UDPPingAddDelReq) GetCrcString() string {
	return "a7280e39"
}
func NewUDPPingAddDelReq() api.Message {
	return &UDPPingAddDelReq{}
}

// UDPPingAddDelReply represents the VPP binary API message 'udp_ping_add_del_reply'.
// Generated from '/usr/share/vpp/api//udp_ping.api.json', line 23:
//
//        ["udp_ping_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa08dec44"}
//        ],
//
type UDPPingAddDelReply struct {
	Retval int32
}

func (*UDPPingAddDelReply) GetMessageName() string {
	return "udp_ping_add_del_reply"
}
func (*UDPPingAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*UDPPingAddDelReply) GetCrcString() string {
	return "a08dec44"
}
func NewUDPPingAddDelReply() api.Message {
	return &UDPPingAddDelReply{}
}

// UDPPingExportReq represents the VPP binary API message 'udp_ping_export_req'.
// Generated from '/usr/share/vpp/api//udp_ping.api.json', line 29:
//
//        ["udp_ping_export_req",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable"],
//            {"crc" : "0xe43a0203"}
//        ],
//
type UDPPingExportReq struct {
	Enable uint32
}

func (*UDPPingExportReq) GetMessageName() string {
	return "udp_ping_export_req"
}
func (*UDPPingExportReq) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*UDPPingExportReq) GetCrcString() string {
	return "e43a0203"
}
func NewUDPPingExportReq() api.Message {
	return &UDPPingExportReq{}
}

// UDPPingExportReply represents the VPP binary API message 'udp_ping_export_reply'.
// Generated from '/usr/share/vpp/api//udp_ping.api.json', line 36:
//
//        ["udp_ping_export_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x7f8a6c87"}
//        ]
//
type UDPPingExportReply struct {
	Retval int32
}

func (*UDPPingExportReply) GetMessageName() string {
	return "udp_ping_export_reply"
}
func (*UDPPingExportReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*UDPPingExportReply) GetCrcString() string {
	return "7f8a6c87"
}
func NewUDPPingExportReply() api.Message {
	return &UDPPingExportReply{}
}
