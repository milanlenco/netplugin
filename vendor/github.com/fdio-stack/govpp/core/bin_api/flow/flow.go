// Package flow represents the VPP binary API of the 'flow' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//flow.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package flow

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xbc0a0bf1

// SetIpfixExporter represents the VPP binary API message 'set_ipfix_exporter'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 6:
//
//        ["set_ipfix_exporter",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "collector_address", 16],
//            ["u16", "collector_port"],
//            ["u8", "src_address", 16],
//            ["u32", "vrf_id"],
//            ["u32", "path_mtu"],
//            ["u32", "template_interval"],
//            ["u8", "udp_checksum"],
//            {"crc" : "0xd4c80f5c"}
//        ],
//
type SetIpfixExporter struct {
	CollectorAddress []byte `struc:"[16]byte"`
	CollectorPort    uint16
	SrcAddress       []byte `struc:"[16]byte"`
	VrfID            uint32
	PathMtu          uint32
	TemplateInterval uint32
	UDPChecksum      uint8
}

func (*SetIpfixExporter) GetMessageName() string {
	return "set_ipfix_exporter"
}
func (*SetIpfixExporter) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SetIpfixExporter) GetCrcString() string {
	return "d4c80f5c"
}
func NewSetIpfixExporter() api.Message {
	return &SetIpfixExporter{}
}

// SetIpfixExporterReply represents the VPP binary API message 'set_ipfix_exporter_reply'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 19:
//
//        ["set_ipfix_exporter_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x40b502e9"}
//        ],
//
type SetIpfixExporterReply struct {
	Retval int32
}

func (*SetIpfixExporterReply) GetMessageName() string {
	return "set_ipfix_exporter_reply"
}
func (*SetIpfixExporterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SetIpfixExporterReply) GetCrcString() string {
	return "40b502e9"
}
func NewSetIpfixExporterReply() api.Message {
	return &SetIpfixExporterReply{}
}

// IpfixExporterDump represents the VPP binary API message 'ipfix_exporter_dump'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 25:
//
//        ["ipfix_exporter_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x81a51716"}
//        ],
//
type IpfixExporterDump struct {
}

func (*IpfixExporterDump) GetMessageName() string {
	return "ipfix_exporter_dump"
}
func (*IpfixExporterDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpfixExporterDump) GetCrcString() string {
	return "81a51716"
}
func NewIpfixExporterDump() api.Message {
	return &IpfixExporterDump{}
}

// IpfixExporterDetails represents the VPP binary API message 'ipfix_exporter_details'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 31:
//
//        ["ipfix_exporter_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "collector_address", 16],
//            ["u16", "collector_port"],
//            ["u8", "src_address", 16],
//            ["u32", "vrf_id"],
//            ["u32", "path_mtu"],
//            ["u32", "template_interval"],
//            ["u8", "udp_checksum"],
//            {"crc" : "0xe7cc717b"}
//        ],
//
type IpfixExporterDetails struct {
	CollectorAddress []byte `struc:"[16]byte"`
	CollectorPort    uint16
	SrcAddress       []byte `struc:"[16]byte"`
	VrfID            uint32
	PathMtu          uint32
	TemplateInterval uint32
	UDPChecksum      uint8
}

func (*IpfixExporterDetails) GetMessageName() string {
	return "ipfix_exporter_details"
}
func (*IpfixExporterDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpfixExporterDetails) GetCrcString() string {
	return "e7cc717b"
}
func NewIpfixExporterDetails() api.Message {
	return &IpfixExporterDetails{}
}

// SetIpfixClassifyStream represents the VPP binary API message 'set_ipfix_classify_stream'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 43:
//
//        ["set_ipfix_classify_stream",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "domain_id"],
//            ["u16", "src_port"],
//            {"crc" : "0x7ee60f3a"}
//        ],
//
type SetIpfixClassifyStream struct {
	DomainID uint32
	SrcPort  uint16
}

func (*SetIpfixClassifyStream) GetMessageName() string {
	return "set_ipfix_classify_stream"
}
func (*SetIpfixClassifyStream) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SetIpfixClassifyStream) GetCrcString() string {
	return "7ee60f3a"
}
func NewSetIpfixClassifyStream() api.Message {
	return &SetIpfixClassifyStream{}
}

// SetIpfixClassifyStreamReply represents the VPP binary API message 'set_ipfix_classify_stream_reply'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 51:
//
//        ["set_ipfix_classify_stream_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa4d2d102"}
//        ],
//
type SetIpfixClassifyStreamReply struct {
	Retval int32
}

func (*SetIpfixClassifyStreamReply) GetMessageName() string {
	return "set_ipfix_classify_stream_reply"
}
func (*SetIpfixClassifyStreamReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SetIpfixClassifyStreamReply) GetCrcString() string {
	return "a4d2d102"
}
func NewSetIpfixClassifyStreamReply() api.Message {
	return &SetIpfixClassifyStreamReply{}
}

// IpfixClassifyStreamDump represents the VPP binary API message 'ipfix_classify_stream_dump'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 57:
//
//        ["ipfix_classify_stream_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x81842294"}
//        ],
//
type IpfixClassifyStreamDump struct {
}

func (*IpfixClassifyStreamDump) GetMessageName() string {
	return "ipfix_classify_stream_dump"
}
func (*IpfixClassifyStreamDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpfixClassifyStreamDump) GetCrcString() string {
	return "81842294"
}
func NewIpfixClassifyStreamDump() api.Message {
	return &IpfixClassifyStreamDump{}
}

// IpfixClassifyStreamDetails represents the VPP binary API message 'ipfix_classify_stream_details'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 63:
//
//        ["ipfix_classify_stream_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "domain_id"],
//            ["u16", "src_port"],
//            {"crc" : "0x6b9383aa"}
//        ],
//
type IpfixClassifyStreamDetails struct {
	DomainID uint32
	SrcPort  uint16
}

func (*IpfixClassifyStreamDetails) GetMessageName() string {
	return "ipfix_classify_stream_details"
}
func (*IpfixClassifyStreamDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpfixClassifyStreamDetails) GetCrcString() string {
	return "6b9383aa"
}
func NewIpfixClassifyStreamDetails() api.Message {
	return &IpfixClassifyStreamDetails{}
}

// IpfixClassifyTableAddDel represents the VPP binary API message 'ipfix_classify_table_add_del'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 70:
//
//        ["ipfix_classify_table_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "table_id"],
//            ["u8", "ip_version"],
//            ["u8", "transport_protocol"],
//            ["u8", "is_add"],
//            {"crc" : "0x52cc2ed9"}
//        ],
//
type IpfixClassifyTableAddDel struct {
	TableID           uint32
	IPVersion         uint8
	TransportProtocol uint8
	IsAdd             uint8
}

func (*IpfixClassifyTableAddDel) GetMessageName() string {
	return "ipfix_classify_table_add_del"
}
func (*IpfixClassifyTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpfixClassifyTableAddDel) GetCrcString() string {
	return "52cc2ed9"
}
func NewIpfixClassifyTableAddDel() api.Message {
	return &IpfixClassifyTableAddDel{}
}

// IpfixClassifyTableAddDelReply represents the VPP binary API message 'ipfix_classify_table_add_del_reply'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 80:
//
//        ["ipfix_classify_table_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x3116af60"}
//        ],
//
type IpfixClassifyTableAddDelReply struct {
	Retval int32
}

func (*IpfixClassifyTableAddDelReply) GetMessageName() string {
	return "ipfix_classify_table_add_del_reply"
}
func (*IpfixClassifyTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpfixClassifyTableAddDelReply) GetCrcString() string {
	return "3116af60"
}
func NewIpfixClassifyTableAddDelReply() api.Message {
	return &IpfixClassifyTableAddDelReply{}
}

// IpfixClassifyTableDump represents the VPP binary API message 'ipfix_classify_table_dump'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 86:
//
//        ["ipfix_classify_table_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xb2ce9db1"}
//        ],
//
type IpfixClassifyTableDump struct {
}

func (*IpfixClassifyTableDump) GetMessageName() string {
	return "ipfix_classify_table_dump"
}
func (*IpfixClassifyTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpfixClassifyTableDump) GetCrcString() string {
	return "b2ce9db1"
}
func NewIpfixClassifyTableDump() api.Message {
	return &IpfixClassifyTableDump{}
}

// IpfixClassifyTableDetails represents the VPP binary API message 'ipfix_classify_table_details'.
// Generated from '/usr/share/vpp/api//flow.api.json', line 92:
//
//        ["ipfix_classify_table_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "table_id"],
//            ["u8", "ip_version"],
//            ["u8", "transport_protocol"],
//            {"crc" : "0xd0ec861f"}
//        ]
//
type IpfixClassifyTableDetails struct {
	TableID           uint32
	IPVersion         uint8
	TransportProtocol uint8
}

func (*IpfixClassifyTableDetails) GetMessageName() string {
	return "ipfix_classify_table_details"
}
func (*IpfixClassifyTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpfixClassifyTableDetails) GetCrcString() string {
	return "d0ec861f"
}
func NewIpfixClassifyTableDetails() api.Message {
	return &IpfixClassifyTableDetails{}
}
