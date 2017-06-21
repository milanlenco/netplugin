// Package lisp_gpe represents the VPP binary API of the 'lisp_gpe' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//lisp_gpe.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package lisp_gpe

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x3ad3dbd9

// GpeLocator represents the VPP binary API data type 'gpe_locator'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 3:
//
//        ["gpe_locator",
//            ["u8", "is_ip4"],
//            ["u8", "weight"],
//            ["u8", "addr", 16],
//            {"crc" : "0x02b20def"}
//        ],
//
type GpeLocator struct {
	IsIP4  uint8
	Weight uint8
	Addr   []byte `struc:"[16]byte"`
}

func (*GpeLocator) GetTypeName() string {
	return "gpe_locator"
}
func (*GpeLocator) GetCrcString() string {
	return "02b20def"
}

// GpeFwdEntry represents the VPP binary API data type 'gpe_fwd_entry'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 9:
//
//        ["gpe_fwd_entry",
//            ["u32", "fwd_entry_index"],
//            ["u32", "dp_table"],
//            ["u8", "eid_type"],
//            ["u8", "leid_prefix_len"],
//            ["u8", "reid_prefix_len"],
//            ["u8", "leid", 16],
//            ["u8", "reid", 16],
//            {"crc" : "0x3220bf07"}
//        ]
//
type GpeFwdEntry struct {
	FwdEntryIndex uint32
	DpTable       uint32
	EidType       uint8
	LeidPrefixLen uint8
	ReidPrefixLen uint8
	Leid          []byte `struc:"[16]byte"`
	Reid          []byte `struc:"[16]byte"`
}

func (*GpeFwdEntry) GetTypeName() string {
	return "gpe_fwd_entry"
}
func (*GpeFwdEntry) GetCrcString() string {
	return "3220bf07"
}

// GpeAddDelFwdEntry represents the VPP binary API message 'gpe_add_del_fwd_entry'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 21:
//
//        ["gpe_add_del_fwd_entry",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "eid_type"],
//            ["u8", "rmt_eid", 16],
//            ["u8", "lcl_eid", 16],
//            ["u8", "rmt_len"],
//            ["u8", "lcl_len"],
//            ["u32", "vni"],
//            ["u32", "dp_table"],
//            ["u8", "action"],
//            ["u32", "loc_num"],
//            ["vl_api_gpe_locator_t", "locs", 0, "loc_num"],
//            {"crc" : "0xbcb43b69"}
//        ],
//
type GpeAddDelFwdEntry struct {
	IsAdd   uint8
	EidType uint8
	RmtEid  []byte `struc:"[16]byte"`
	LclEid  []byte `struc:"[16]byte"`
	RmtLen  uint8
	LclLen  uint8
	Vni     uint32
	DpTable uint32
	Action  uint8
	LocNum  uint32 `struc:"sizeof=Locs"`
	Locs    []GpeLocator
}

func (*GpeAddDelFwdEntry) GetMessageName() string {
	return "gpe_add_del_fwd_entry"
}
func (*GpeAddDelFwdEntry) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeAddDelFwdEntry) GetCrcString() string {
	return "bcb43b69"
}
func NewGpeAddDelFwdEntry() api.Message {
	return &GpeAddDelFwdEntry{}
}

// GpeAddDelFwdEntryReply represents the VPP binary API message 'gpe_add_del_fwd_entry_reply'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 38:
//
//        ["gpe_add_del_fwd_entry_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd15edd0e"}
//        ],
//
type GpeAddDelFwdEntryReply struct {
	Retval int32
}

func (*GpeAddDelFwdEntryReply) GetMessageName() string {
	return "gpe_add_del_fwd_entry_reply"
}
func (*GpeAddDelFwdEntryReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GpeAddDelFwdEntryReply) GetCrcString() string {
	return "d15edd0e"
}
func NewGpeAddDelFwdEntryReply() api.Message {
	return &GpeAddDelFwdEntryReply{}
}

// GpeEnableDisable represents the VPP binary API message 'gpe_enable_disable'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 44:
//
//        ["gpe_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_en"],
//            {"crc" : "0x05e310d8"}
//        ],
//
type GpeEnableDisable struct {
	IsEn uint8
}

func (*GpeEnableDisable) GetMessageName() string {
	return "gpe_enable_disable"
}
func (*GpeEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeEnableDisable) GetCrcString() string {
	return "05e310d8"
}
func NewGpeEnableDisable() api.Message {
	return &GpeEnableDisable{}
}

// GpeEnableDisableReply represents the VPP binary API message 'gpe_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 51:
//
//        ["gpe_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd021cf76"}
//        ],
//
type GpeEnableDisableReply struct {
	Retval int32
}

func (*GpeEnableDisableReply) GetMessageName() string {
	return "gpe_enable_disable_reply"
}
func (*GpeEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GpeEnableDisableReply) GetCrcString() string {
	return "d021cf76"
}
func NewGpeEnableDisableReply() api.Message {
	return &GpeEnableDisableReply{}
}

// GpeAddDelIface represents the VPP binary API message 'gpe_add_del_iface'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 57:
//
//        ["gpe_add_del_iface",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_l2"],
//            ["u32", "dp_table"],
//            ["u32", "vni"],
//            {"crc" : "0xe4870f5a"}
//        ],
//
type GpeAddDelIface struct {
	IsAdd   uint8
	IsL2    uint8
	DpTable uint32
	Vni     uint32
}

func (*GpeAddDelIface) GetMessageName() string {
	return "gpe_add_del_iface"
}
func (*GpeAddDelIface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeAddDelIface) GetCrcString() string {
	return "e4870f5a"
}
func NewGpeAddDelIface() api.Message {
	return &GpeAddDelIface{}
}

// GpeAddDelIfaceReply represents the VPP binary API message 'gpe_add_del_iface_reply'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 67:
//
//        ["gpe_add_del_iface_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x92a5cddb"}
//        ],
//
type GpeAddDelIfaceReply struct {
	Retval int32
}

func (*GpeAddDelIfaceReply) GetMessageName() string {
	return "gpe_add_del_iface_reply"
}
func (*GpeAddDelIfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GpeAddDelIfaceReply) GetCrcString() string {
	return "92a5cddb"
}
func NewGpeAddDelIfaceReply() api.Message {
	return &GpeAddDelIfaceReply{}
}

// GpeFwdEntriesGet represents the VPP binary API message 'gpe_fwd_entries_get'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 73:
//
//        ["gpe_fwd_entries_get",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            {"crc" : "0x113b9a39"}
//        ],
//
type GpeFwdEntriesGet struct {
	Vni uint32
}

func (*GpeFwdEntriesGet) GetMessageName() string {
	return "gpe_fwd_entries_get"
}
func (*GpeFwdEntriesGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeFwdEntriesGet) GetCrcString() string {
	return "113b9a39"
}
func NewGpeFwdEntriesGet() api.Message {
	return &GpeFwdEntriesGet{}
}

// GpeFwdEntriesGetReply represents the VPP binary API message 'gpe_fwd_entries_get_reply'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 80:
//
//        ["gpe_fwd_entries_get_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "count"],
//            ["vl_api_gpe_fwd_entry_t", "entries", 0, "count"],
//            {"crc" : "0x484da49c"}
//        ],
//
type GpeFwdEntriesGetReply struct {
	Retval  int32
	Count   uint32 `struc:"sizeof=Entries"`
	Entries []GpeFwdEntry
}

func (*GpeFwdEntriesGetReply) GetMessageName() string {
	return "gpe_fwd_entries_get_reply"
}
func (*GpeFwdEntriesGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GpeFwdEntriesGetReply) GetCrcString() string {
	return "484da49c"
}
func NewGpeFwdEntriesGetReply() api.Message {
	return &GpeFwdEntriesGetReply{}
}

// GpeFwdEntryPathDump represents the VPP binary API message 'gpe_fwd_entry_path_dump'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 88:
//
//        ["gpe_fwd_entry_path_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "fwd_entry_index"],
//            {"crc" : "0x30b320f6"}
//        ],
//
type GpeFwdEntryPathDump struct {
	FwdEntryIndex uint32
}

func (*GpeFwdEntryPathDump) GetMessageName() string {
	return "gpe_fwd_entry_path_dump"
}
func (*GpeFwdEntryPathDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeFwdEntryPathDump) GetCrcString() string {
	return "30b320f6"
}
func NewGpeFwdEntryPathDump() api.Message {
	return &GpeFwdEntryPathDump{}
}

// GpeFwdEntryPathDetails represents the VPP binary API message 'gpe_fwd_entry_path_details'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 95:
//
//        ["gpe_fwd_entry_path_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["vl_api_gpe_locator_t", "lcl_loc"],
//            ["vl_api_gpe_locator_t", "rmt_loc"],
//            {"crc" : "0xa9187578"}
//        ],
//
type GpeFwdEntryPathDetails struct {
	LclLoc GpeLocator
	RmtLoc GpeLocator
}

func (*GpeFwdEntryPathDetails) GetMessageName() string {
	return "gpe_fwd_entry_path_details"
}
func (*GpeFwdEntryPathDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeFwdEntryPathDetails) GetCrcString() string {
	return "a9187578"
}
func NewGpeFwdEntryPathDetails() api.Message {
	return &GpeFwdEntryPathDetails{}
}

// GpeSetEncapMode represents the VPP binary API message 'gpe_set_encap_mode'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 103:
//
//        ["gpe_set_encap_mode",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "mode"],
//            {"crc" : "0x6b0767d1"}
//        ],
//
type GpeSetEncapMode struct {
	Mode uint8
}

func (*GpeSetEncapMode) GetMessageName() string {
	return "gpe_set_encap_mode"
}
func (*GpeSetEncapMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeSetEncapMode) GetCrcString() string {
	return "6b0767d1"
}
func NewGpeSetEncapMode() api.Message {
	return &GpeSetEncapMode{}
}

// GpeSetEncapModeReply represents the VPP binary API message 'gpe_set_encap_mode_reply'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 110:
//
//        ["gpe_set_encap_mode_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xdcadb033"}
//        ],
//
type GpeSetEncapModeReply struct {
	Retval int32
}

func (*GpeSetEncapModeReply) GetMessageName() string {
	return "gpe_set_encap_mode_reply"
}
func (*GpeSetEncapModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GpeSetEncapModeReply) GetCrcString() string {
	return "dcadb033"
}
func NewGpeSetEncapModeReply() api.Message {
	return &GpeSetEncapModeReply{}
}

// GpeGetEncapMode represents the VPP binary API message 'gpe_get_encap_mode'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 116:
//
//        ["gpe_get_encap_mode",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xc8c38e9f"}
//        ],
//
type GpeGetEncapMode struct {
}

func (*GpeGetEncapMode) GetMessageName() string {
	return "gpe_get_encap_mode"
}
func (*GpeGetEncapMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GpeGetEncapMode) GetCrcString() string {
	return "c8c38e9f"
}
func NewGpeGetEncapMode() api.Message {
	return &GpeGetEncapMode{}
}

// GpeGetEncapModeReply represents the VPP binary API message 'gpe_get_encap_mode_reply'.
// Generated from '/usr/share/vpp/api//lisp_gpe.api.json', line 122:
//
//        ["gpe_get_encap_mode_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "encap_mode"],
//            {"crc" : "0x90637a31"}
//        ]
//
type GpeGetEncapModeReply struct {
	Retval    int32
	EncapMode uint8
}

func (*GpeGetEncapModeReply) GetMessageName() string {
	return "gpe_get_encap_mode_reply"
}
func (*GpeGetEncapModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GpeGetEncapModeReply) GetCrcString() string {
	return "90637a31"
}
func NewGpeGetEncapModeReply() api.Message {
	return &GpeGetEncapModeReply{}
}
