// Package classify represents the VPP binary API of the 'classify' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//classify.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package classify

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x13320cb7

// ClassifyAddDelTable represents the VPP binary API message 'classify_add_del_table'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 6:
//
//        ["classify_add_del_table",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "del_chain"],
//            ["u32", "table_index"],
//            ["u32", "nbuckets"],
//            ["u32", "memory_size"],
//            ["u32", "skip_n_vectors"],
//            ["u32", "match_n_vectors"],
//            ["u32", "next_table_index"],
//            ["u32", "miss_next_index"],
//            ["u32", "current_data_flag"],
//            ["i32", "current_data_offset"],
//            ["u8", "mask", 0],
//            {"crc" : "0x1120f35d"}
//        ],
//
type ClassifyAddDelTable struct {
	IsAdd             uint8
	DelChain          uint8
	TableIndex        uint32
	Nbuckets          uint32
	MemorySize        uint32
	SkipNVectors      uint32
	MatchNVectors     uint32
	NextTableIndex    uint32
	MissNextIndex     uint32
	CurrentDataFlag   uint32
	CurrentDataOffset int32
	Mask              []byte
}

func (*ClassifyAddDelTable) GetMessageName() string {
	return "classify_add_del_table"
}
func (*ClassifyAddDelTable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifyAddDelTable) GetCrcString() string {
	return "1120f35d"
}
func NewClassifyAddDelTable() api.Message {
	return &ClassifyAddDelTable{}
}

// ClassifyAddDelTableReply represents the VPP binary API message 'classify_add_del_table_reply'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 24:
//
//        ["classify_add_del_table_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "new_table_index"],
//            ["u32", "skip_n_vectors"],
//            ["u32", "match_n_vectors"],
//            {"crc" : "0xd4e63320"}
//        ],
//
type ClassifyAddDelTableReply struct {
	Retval        int32
	NewTableIndex uint32
	SkipNVectors  uint32
	MatchNVectors uint32
}

func (*ClassifyAddDelTableReply) GetMessageName() string {
	return "classify_add_del_table_reply"
}
func (*ClassifyAddDelTableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifyAddDelTableReply) GetCrcString() string {
	return "d4e63320"
}
func NewClassifyAddDelTableReply() api.Message {
	return &ClassifyAddDelTableReply{}
}

// ClassifyAddDelSession represents the VPP binary API message 'classify_add_del_session'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 33:
//
//        ["classify_add_del_session",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "table_index"],
//            ["u32", "hit_next_index"],
//            ["u32", "opaque_index"],
//            ["i32", "advance"],
//            ["u8", "action"],
//            ["u32", "metadata"],
//            ["u8", "match", 0],
//            {"crc" : "0x25a952f5"}
//        ],
//
type ClassifyAddDelSession struct {
	IsAdd        uint8
	TableIndex   uint32
	HitNextIndex uint32
	OpaqueIndex  uint32
	Advance      int32
	Action       uint8
	Metadata     uint32
	Match        []byte
}

func (*ClassifyAddDelSession) GetMessageName() string {
	return "classify_add_del_session"
}
func (*ClassifyAddDelSession) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifyAddDelSession) GetCrcString() string {
	return "25a952f5"
}
func NewClassifyAddDelSession() api.Message {
	return &ClassifyAddDelSession{}
}

// ClassifyAddDelSessionReply represents the VPP binary API message 'classify_add_del_session_reply'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 47:
//
//        ["classify_add_del_session_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xdd4aa9ac"}
//        ],
//
type ClassifyAddDelSessionReply struct {
	Retval int32
}

func (*ClassifyAddDelSessionReply) GetMessageName() string {
	return "classify_add_del_session_reply"
}
func (*ClassifyAddDelSessionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifyAddDelSessionReply) GetCrcString() string {
	return "dd4aa9ac"
}
func NewClassifyAddDelSessionReply() api.Message {
	return &ClassifyAddDelSessionReply{}
}

// PolicerClassifySetInterface represents the VPP binary API message 'policer_classify_set_interface'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 53:
//
//        ["policer_classify_set_interface",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "ip4_table_index"],
//            ["u32", "ip6_table_index"],
//            ["u32", "l2_table_index"],
//            ["u8", "is_add"],
//            {"crc" : "0x4ad6b5a8"}
//        ],
//
type PolicerClassifySetInterface struct {
	SwIfIndex     uint32
	IP4TableIndex uint32
	IP6TableIndex uint32
	L2TableIndex  uint32
	IsAdd         uint8
}

func (*PolicerClassifySetInterface) GetMessageName() string {
	return "policer_classify_set_interface"
}
func (*PolicerClassifySetInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PolicerClassifySetInterface) GetCrcString() string {
	return "4ad6b5a8"
}
func NewPolicerClassifySetInterface() api.Message {
	return &PolicerClassifySetInterface{}
}

// PolicerClassifySetInterfaceReply represents the VPP binary API message 'policer_classify_set_interface_reply'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 64:
//
//        ["policer_classify_set_interface_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x003dd29f"}
//        ],
//
type PolicerClassifySetInterfaceReply struct {
	Retval int32
}

func (*PolicerClassifySetInterfaceReply) GetMessageName() string {
	return "policer_classify_set_interface_reply"
}
func (*PolicerClassifySetInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PolicerClassifySetInterfaceReply) GetCrcString() string {
	return "003dd29f"
}
func NewPolicerClassifySetInterfaceReply() api.Message {
	return &PolicerClassifySetInterfaceReply{}
}

// PolicerClassifyDump represents the VPP binary API message 'policer_classify_dump'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 70:
//
//        ["policer_classify_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "type"],
//            {"crc" : "0x593ab73c"}
//        ],
//
type PolicerClassifyDump struct {
	Type uint8
}

func (*PolicerClassifyDump) GetMessageName() string {
	return "policer_classify_dump"
}
func (*PolicerClassifyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PolicerClassifyDump) GetCrcString() string {
	return "593ab73c"
}
func NewPolicerClassifyDump() api.Message {
	return &PolicerClassifyDump{}
}

// PolicerClassifyDetails represents the VPP binary API message 'policer_classify_details'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 77:
//
//        ["policer_classify_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "table_index"],
//            {"crc" : "0xe3439be8"}
//        ],
//
type PolicerClassifyDetails struct {
	SwIfIndex  uint32
	TableIndex uint32
}

func (*PolicerClassifyDetails) GetMessageName() string {
	return "policer_classify_details"
}
func (*PolicerClassifyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PolicerClassifyDetails) GetCrcString() string {
	return "e3439be8"
}
func NewPolicerClassifyDetails() api.Message {
	return &PolicerClassifyDetails{}
}

// ClassifyTableIds represents the VPP binary API message 'classify_table_ids'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 84:
//
//        ["classify_table_ids",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xfea5bc0b"}
//        ],
//
type ClassifyTableIds struct {
}

func (*ClassifyTableIds) GetMessageName() string {
	return "classify_table_ids"
}
func (*ClassifyTableIds) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifyTableIds) GetCrcString() string {
	return "fea5bc0b"
}
func NewClassifyTableIds() api.Message {
	return &ClassifyTableIds{}
}

// ClassifyTableIdsReply represents the VPP binary API message 'classify_table_ids_reply'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 90:
//
//        ["classify_table_ids_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "count"],
//            ["u32", "ids", 0, "count"],
//            {"crc" : "0x6eacee4e"}
//        ],
//
type ClassifyTableIdsReply struct {
	Retval int32
	Count  uint32 `struc:"sizeof=Ids"`
	Ids    []uint32
}

func (*ClassifyTableIdsReply) GetMessageName() string {
	return "classify_table_ids_reply"
}
func (*ClassifyTableIdsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifyTableIdsReply) GetCrcString() string {
	return "6eacee4e"
}
func NewClassifyTableIdsReply() api.Message {
	return &ClassifyTableIdsReply{}
}

// ClassifyTableByInterface represents the VPP binary API message 'classify_table_by_interface'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 98:
//
//        ["classify_table_by_interface",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x2dc2ff38"}
//        ],
//
type ClassifyTableByInterface struct {
	SwIfIndex uint32
}

func (*ClassifyTableByInterface) GetMessageName() string {
	return "classify_table_by_interface"
}
func (*ClassifyTableByInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifyTableByInterface) GetCrcString() string {
	return "2dc2ff38"
}
func NewClassifyTableByInterface() api.Message {
	return &ClassifyTableByInterface{}
}

// ClassifyTableByInterfaceReply represents the VPP binary API message 'classify_table_by_interface_reply'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 105:
//
//        ["classify_table_by_interface_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            ["u32", "l2_table_id"],
//            ["u32", "ip4_table_id"],
//            ["u32", "ip6_table_id"],
//            {"crc" : "0x7a9ac873"}
//        ],
//
type ClassifyTableByInterfaceReply struct {
	Retval     int32
	SwIfIndex  uint32
	L2TableID  uint32
	IP4TableID uint32
	IP6TableID uint32
}

func (*ClassifyTableByInterfaceReply) GetMessageName() string {
	return "classify_table_by_interface_reply"
}
func (*ClassifyTableByInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifyTableByInterfaceReply) GetCrcString() string {
	return "7a9ac873"
}
func NewClassifyTableByInterfaceReply() api.Message {
	return &ClassifyTableByInterfaceReply{}
}

// ClassifyTableInfo represents the VPP binary API message 'classify_table_info'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 115:
//
//        ["classify_table_info",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "table_id"],
//            {"crc" : "0x33caf8c6"}
//        ],
//
type ClassifyTableInfo struct {
	TableID uint32
}

func (*ClassifyTableInfo) GetMessageName() string {
	return "classify_table_info"
}
func (*ClassifyTableInfo) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifyTableInfo) GetCrcString() string {
	return "33caf8c6"
}
func NewClassifyTableInfo() api.Message {
	return &ClassifyTableInfo{}
}

// ClassifyTableInfoReply represents the VPP binary API message 'classify_table_info_reply'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 122:
//
//        ["classify_table_info_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "table_id"],
//            ["u32", "nbuckets"],
//            ["u32", "match_n_vectors"],
//            ["u32", "skip_n_vectors"],
//            ["u32", "active_sessions"],
//            ["u32", "next_table_index"],
//            ["u32", "miss_next_index"],
//            ["u32", "mask_length"],
//            ["u8", "mask", 0, "mask_length"],
//            {"crc" : "0x60312b83"}
//        ],
//
type ClassifyTableInfoReply struct {
	Retval         int32
	TableID        uint32
	Nbuckets       uint32
	MatchNVectors  uint32
	SkipNVectors   uint32
	ActiveSessions uint32
	NextTableIndex uint32
	MissNextIndex  uint32
	MaskLength     uint32 `struc:"sizeof=Mask"`
	Mask           []byte
}

func (*ClassifyTableInfoReply) GetMessageName() string {
	return "classify_table_info_reply"
}
func (*ClassifyTableInfoReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifyTableInfoReply) GetCrcString() string {
	return "60312b83"
}
func NewClassifyTableInfoReply() api.Message {
	return &ClassifyTableInfoReply{}
}

// ClassifySessionDump represents the VPP binary API message 'classify_session_dump'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 137:
//
//        ["classify_session_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "table_id"],
//            {"crc" : "0x87d2ca2b"}
//        ],
//
type ClassifySessionDump struct {
	TableID uint32
}

func (*ClassifySessionDump) GetMessageName() string {
	return "classify_session_dump"
}
func (*ClassifySessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifySessionDump) GetCrcString() string {
	return "87d2ca2b"
}
func NewClassifySessionDump() api.Message {
	return &ClassifySessionDump{}
}

// ClassifySessionDetails represents the VPP binary API message 'classify_session_details'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 144:
//
//        ["classify_session_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "table_id"],
//            ["u32", "hit_next_index"],
//            ["i32", "advance"],
//            ["u32", "opaque_index"],
//            ["u32", "match_length"],
//            ["u8", "match", 0, "match_length"],
//            {"crc" : "0x95efd073"}
//        ],
//
type ClassifySessionDetails struct {
	Retval       int32
	TableID      uint32
	HitNextIndex uint32
	Advance      int32
	OpaqueIndex  uint32
	MatchLength  uint32 `struc:"sizeof=Match"`
	Match        []byte
}

func (*ClassifySessionDetails) GetMessageName() string {
	return "classify_session_details"
}
func (*ClassifySessionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifySessionDetails) GetCrcString() string {
	return "95efd073"
}
func NewClassifySessionDetails() api.Message {
	return &ClassifySessionDetails{}
}

// FlowClassifySetInterface represents the VPP binary API message 'flow_classify_set_interface'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 156:
//
//        ["flow_classify_set_interface",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "ip4_table_index"],
//            ["u32", "ip6_table_index"],
//            ["u8", "is_add"],
//            {"crc" : "0x6e0a565e"}
//        ],
//
type FlowClassifySetInterface struct {
	SwIfIndex     uint32
	IP4TableIndex uint32
	IP6TableIndex uint32
	IsAdd         uint8
}

func (*FlowClassifySetInterface) GetMessageName() string {
	return "flow_classify_set_interface"
}
func (*FlowClassifySetInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*FlowClassifySetInterface) GetCrcString() string {
	return "6e0a565e"
}
func NewFlowClassifySetInterface() api.Message {
	return &FlowClassifySetInterface{}
}

// FlowClassifySetInterfaceReply represents the VPP binary API message 'flow_classify_set_interface_reply'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 166:
//
//        ["flow_classify_set_interface_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x3407e7c3"}
//        ],
//
type FlowClassifySetInterfaceReply struct {
	Retval int32
}

func (*FlowClassifySetInterfaceReply) GetMessageName() string {
	return "flow_classify_set_interface_reply"
}
func (*FlowClassifySetInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*FlowClassifySetInterfaceReply) GetCrcString() string {
	return "3407e7c3"
}
func NewFlowClassifySetInterfaceReply() api.Message {
	return &FlowClassifySetInterfaceReply{}
}

// FlowClassifyDump represents the VPP binary API message 'flow_classify_dump'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 172:
//
//        ["flow_classify_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "type"],
//            {"crc" : "0x97f781c8"}
//        ],
//
type FlowClassifyDump struct {
	Type uint8
}

func (*FlowClassifyDump) GetMessageName() string {
	return "flow_classify_dump"
}
func (*FlowClassifyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*FlowClassifyDump) GetCrcString() string {
	return "97f781c8"
}
func NewFlowClassifyDump() api.Message {
	return &FlowClassifyDump{}
}

// FlowClassifyDetails represents the VPP binary API message 'flow_classify_details'.
// Generated from '/usr/share/vpp/api//classify.api.json', line 179:
//
//        ["flow_classify_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "table_index"],
//            {"crc" : "0x08475e65"}
//        ]
//
type FlowClassifyDetails struct {
	SwIfIndex  uint32
	TableIndex uint32
}

func (*FlowClassifyDetails) GetMessageName() string {
	return "flow_classify_details"
}
func (*FlowClassifyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*FlowClassifyDetails) GetCrcString() string {
	return "08475e65"
}
func NewFlowClassifyDetails() api.Message {
	return &FlowClassifyDetails{}
}
