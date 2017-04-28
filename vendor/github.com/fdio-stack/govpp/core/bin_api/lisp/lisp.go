// Package lisp represents the VPP binary API of the 'lisp' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//lisp.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package lisp

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x4da3ff1c

// LocalLocator represents the VPP binary API data type 'local_locator'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 3:
//
//        ["local_locator",
//            ["u32", "sw_if_index"],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            {"crc" : "0x0ad03b93"}
//        ],
//
type LocalLocator struct {
	SwIfIndex uint32
	Priority  uint8
	Weight    uint8
}

func (*LocalLocator) GetTypeName() string {
	return "local_locator"
}
func (*LocalLocator) GetCrcString() string {
	return "0ad03b93"
}

// RemoteLocator represents the VPP binary API data type 'remote_locator'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 9:
//
//        ["remote_locator",
//            ["u8", "is_ip4"],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            ["u8", "addr", 16],
//            {"crc" : "0xf42e0a8f"}
//        ],
//
type RemoteLocator struct {
	IsIP4    uint8
	Priority uint8
	Weight   uint8
	Addr     []byte `struc:"[16]byte"`
}

func (*RemoteLocator) GetTypeName() string {
	return "remote_locator"
}
func (*RemoteLocator) GetCrcString() string {
	return "f42e0a8f"
}

// LispAdjacency represents the VPP binary API data type 'lisp_adjacency'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 16:
//
//        ["lisp_adjacency",
//            ["u8", "eid_type"],
//            ["u8", "reid", 16],
//            ["u8", "leid", 16],
//            ["u8", "reid_prefix_len"],
//            ["u8", "leid_prefix_len"],
//            {"crc" : "0xade34024"}
//        ]
//
type LispAdjacency struct {
	EidType       uint8
	Reid          []byte `struc:"[16]byte"`
	Leid          []byte `struc:"[16]byte"`
	ReidPrefixLen uint8
	LeidPrefixLen uint8
}

func (*LispAdjacency) GetTypeName() string {
	return "lisp_adjacency"
}
func (*LispAdjacency) GetCrcString() string {
	return "ade34024"
}

// LispAddDelLocatorSet represents the VPP binary API message 'lisp_add_del_locator_set'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 26:
//
//        ["lisp_add_del_locator_set",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "locator_set_name", 64],
//            ["u32", "locator_num"],
//            ["vl_api_local_locator_t", "locators", 0, "locator_num"],
//            {"crc" : "0x25774ef4"}
//        ],
//
type LispAddDelLocatorSet struct {
	IsAdd          uint8
	LocatorSetName []byte `struc:"[64]byte"`
	LocatorNum     uint32 `struc:"sizeof=Locators"`
	Locators       []LocalLocator
}

func (*LispAddDelLocatorSet) GetMessageName() string {
	return "lisp_add_del_locator_set"
}
func (*LispAddDelLocatorSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelLocatorSet) GetCrcString() string {
	return "25774ef4"
}
func NewLispAddDelLocatorSet() api.Message {
	return &LispAddDelLocatorSet{}
}

// LispAddDelLocatorSetReply represents the VPP binary API message 'lisp_add_del_locator_set_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 36:
//
//        ["lisp_add_del_locator_set_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "ls_index"],
//            {"crc" : "0xc167cab1"}
//        ],
//
type LispAddDelLocatorSetReply struct {
	Retval  int32
	LsIndex uint32
}

func (*LispAddDelLocatorSetReply) GetMessageName() string {
	return "lisp_add_del_locator_set_reply"
}
func (*LispAddDelLocatorSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelLocatorSetReply) GetCrcString() string {
	return "c167cab1"
}
func NewLispAddDelLocatorSetReply() api.Message {
	return &LispAddDelLocatorSetReply{}
}

// LispAddDelLocator represents the VPP binary API message 'lisp_add_del_locator'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 43:
//
//        ["lisp_add_del_locator",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "locator_set_name", 64],
//            ["u32", "sw_if_index"],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            {"crc" : "0x442b7292"}
//        ],
//
type LispAddDelLocator struct {
	IsAdd          uint8
	LocatorSetName []byte `struc:"[64]byte"`
	SwIfIndex      uint32
	Priority       uint8
	Weight         uint8
}

func (*LispAddDelLocator) GetMessageName() string {
	return "lisp_add_del_locator"
}
func (*LispAddDelLocator) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelLocator) GetCrcString() string {
	return "442b7292"
}
func NewLispAddDelLocator() api.Message {
	return &LispAddDelLocator{}
}

// LispAddDelLocatorReply represents the VPP binary API message 'lisp_add_del_locator_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 54:
//
//        ["lisp_add_del_locator_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x89ce7940"}
//        ],
//
type LispAddDelLocatorReply struct {
	Retval int32
}

func (*LispAddDelLocatorReply) GetMessageName() string {
	return "lisp_add_del_locator_reply"
}
func (*LispAddDelLocatorReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelLocatorReply) GetCrcString() string {
	return "89ce7940"
}
func NewLispAddDelLocatorReply() api.Message {
	return &LispAddDelLocatorReply{}
}

// LispAddDelLocalEid represents the VPP binary API message 'lisp_add_del_local_eid'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 60:
//
//        ["lisp_add_del_local_eid",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "eid_type"],
//            ["u8", "eid", 16],
//            ["u8", "prefix_len"],
//            ["u8", "locator_set_name", 64],
//            ["u32", "vni"],
//            ["u16", "key_id"],
//            ["u8", "key", 64],
//            {"crc" : "0xc0af8d48"}
//        ],
//
type LispAddDelLocalEid struct {
	IsAdd          uint8
	EidType        uint8
	Eid            []byte `struc:"[16]byte"`
	PrefixLen      uint8
	LocatorSetName []byte `struc:"[64]byte"`
	Vni            uint32
	KeyID          uint16
	Key            []byte `struc:"[64]byte"`
}

func (*LispAddDelLocalEid) GetMessageName() string {
	return "lisp_add_del_local_eid"
}
func (*LispAddDelLocalEid) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelLocalEid) GetCrcString() string {
	return "c0af8d48"
}
func NewLispAddDelLocalEid() api.Message {
	return &LispAddDelLocalEid{}
}

// LispAddDelLocalEidReply represents the VPP binary API message 'lisp_add_del_local_eid_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 74:
//
//        ["lisp_add_del_local_eid_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd4860470"}
//        ],
//
type LispAddDelLocalEidReply struct {
	Retval int32
}

func (*LispAddDelLocalEidReply) GetMessageName() string {
	return "lisp_add_del_local_eid_reply"
}
func (*LispAddDelLocalEidReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelLocalEidReply) GetCrcString() string {
	return "d4860470"
}
func NewLispAddDelLocalEidReply() api.Message {
	return &LispAddDelLocalEidReply{}
}

// LispAddDelMapServer represents the VPP binary API message 'lisp_add_del_map_server'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 80:
//
//        ["lisp_add_del_map_server",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0x592b70b3"}
//        ],
//
type LispAddDelMapServer struct {
	IsAdd     uint8
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*LispAddDelMapServer) GetMessageName() string {
	return "lisp_add_del_map_server"
}
func (*LispAddDelMapServer) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelMapServer) GetCrcString() string {
	return "592b70b3"
}
func NewLispAddDelMapServer() api.Message {
	return &LispAddDelMapServer{}
}

// LispAddDelMapServerReply represents the VPP binary API message 'lisp_add_del_map_server_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 89:
//
//        ["lisp_add_del_map_server_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x26f8f732"}
//        ],
//
type LispAddDelMapServerReply struct {
	Retval int32
}

func (*LispAddDelMapServerReply) GetMessageName() string {
	return "lisp_add_del_map_server_reply"
}
func (*LispAddDelMapServerReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelMapServerReply) GetCrcString() string {
	return "26f8f732"
}
func NewLispAddDelMapServerReply() api.Message {
	return &LispAddDelMapServerReply{}
}

// LispAddDelMapResolver represents the VPP binary API message 'lisp_add_del_map_resolver'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 95:
//
//        ["lisp_add_del_map_resolver",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0x1d0303ff"}
//        ],
//
type LispAddDelMapResolver struct {
	IsAdd     uint8
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*LispAddDelMapResolver) GetMessageName() string {
	return "lisp_add_del_map_resolver"
}
func (*LispAddDelMapResolver) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelMapResolver) GetCrcString() string {
	return "1d0303ff"
}
func NewLispAddDelMapResolver() api.Message {
	return &LispAddDelMapResolver{}
}

// LispAddDelMapResolverReply represents the VPP binary API message 'lisp_add_del_map_resolver_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 104:
//
//        ["lisp_add_del_map_resolver_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xda4d72e0"}
//        ],
//
type LispAddDelMapResolverReply struct {
	Retval int32
}

func (*LispAddDelMapResolverReply) GetMessageName() string {
	return "lisp_add_del_map_resolver_reply"
}
func (*LispAddDelMapResolverReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelMapResolverReply) GetCrcString() string {
	return "da4d72e0"
}
func NewLispAddDelMapResolverReply() api.Message {
	return &LispAddDelMapResolverReply{}
}

// LispEnableDisable represents the VPP binary API message 'lisp_enable_disable'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 110:
//
//        ["lisp_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_en"],
//            {"crc" : "0x6c27720f"}
//        ],
//
type LispEnableDisable struct {
	IsEn uint8
}

func (*LispEnableDisable) GetMessageName() string {
	return "lisp_enable_disable"
}
func (*LispEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispEnableDisable) GetCrcString() string {
	return "6c27720f"
}
func NewLispEnableDisable() api.Message {
	return &LispEnableDisable{}
}

// LispEnableDisableReply represents the VPP binary API message 'lisp_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 117:
//
//        ["lisp_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x36406c38"}
//        ],
//
type LispEnableDisableReply struct {
	Retval int32
}

func (*LispEnableDisableReply) GetMessageName() string {
	return "lisp_enable_disable_reply"
}
func (*LispEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispEnableDisableReply) GetCrcString() string {
	return "36406c38"
}
func NewLispEnableDisableReply() api.Message {
	return &LispEnableDisableReply{}
}

// LispPitrSetLocatorSet represents the VPP binary API message 'lisp_pitr_set_locator_set'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 123:
//
//        ["lisp_pitr_set_locator_set",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "ls_name", 64],
//            {"crc" : "0x59fbff25"}
//        ],
//
type LispPitrSetLocatorSet struct {
	IsAdd  uint8
	LsName []byte `struc:"[64]byte"`
}

func (*LispPitrSetLocatorSet) GetMessageName() string {
	return "lisp_pitr_set_locator_set"
}
func (*LispPitrSetLocatorSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispPitrSetLocatorSet) GetCrcString() string {
	return "59fbff25"
}
func NewLispPitrSetLocatorSet() api.Message {
	return &LispPitrSetLocatorSet{}
}

// LispPitrSetLocatorSetReply represents the VPP binary API message 'lisp_pitr_set_locator_set_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 131:
//
//        ["lisp_pitr_set_locator_set_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x46d996c2"}
//        ],
//
type LispPitrSetLocatorSetReply struct {
	Retval int32
}

func (*LispPitrSetLocatorSetReply) GetMessageName() string {
	return "lisp_pitr_set_locator_set_reply"
}
func (*LispPitrSetLocatorSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispPitrSetLocatorSetReply) GetCrcString() string {
	return "46d996c2"
}
func NewLispPitrSetLocatorSetReply() api.Message {
	return &LispPitrSetLocatorSetReply{}
}

// LispUsePetr represents the VPP binary API message 'lisp_use_petr'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 137:
//
//        ["lisp_use_petr",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "address", 16],
//            ["u8", "is_add"],
//            {"crc" : "0x5b3f74b8"}
//        ],
//
type LispUsePetr struct {
	IsIP4   uint8
	Address []byte `struc:"[16]byte"`
	IsAdd   uint8
}

func (*LispUsePetr) GetMessageName() string {
	return "lisp_use_petr"
}
func (*LispUsePetr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispUsePetr) GetCrcString() string {
	return "5b3f74b8"
}
func NewLispUsePetr() api.Message {
	return &LispUsePetr{}
}

// LispUsePetrReply represents the VPP binary API message 'lisp_use_petr_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 146:
//
//        ["lisp_use_petr_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x752b0f68"}
//        ],
//
type LispUsePetrReply struct {
	Retval int32
}

func (*LispUsePetrReply) GetMessageName() string {
	return "lisp_use_petr_reply"
}
func (*LispUsePetrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispUsePetrReply) GetCrcString() string {
	return "752b0f68"
}
func NewLispUsePetrReply() api.Message {
	return &LispUsePetrReply{}
}

// ShowLispUsePetr represents the VPP binary API message 'show_lisp_use_petr'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 152:
//
//        ["show_lisp_use_petr",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x69cf8a34"}
//        ],
//
type ShowLispUsePetr struct {
}

func (*ShowLispUsePetr) GetMessageName() string {
	return "show_lisp_use_petr"
}
func (*ShowLispUsePetr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowLispUsePetr) GetCrcString() string {
	return "69cf8a34"
}
func NewShowLispUsePetr() api.Message {
	return &ShowLispUsePetr{}
}

// ShowLispUsePetrReply represents the VPP binary API message 'show_lisp_use_petr_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 158:
//
//        ["show_lisp_use_petr_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "status"],
//            ["u8", "is_ip4"],
//            ["u8", "address", 16],
//            {"crc" : "0x8e9360a8"}
//        ],
//
type ShowLispUsePetrReply struct {
	Retval  int32
	Status  uint8
	IsIP4   uint8
	Address []byte `struc:"[16]byte"`
}

func (*ShowLispUsePetrReply) GetMessageName() string {
	return "show_lisp_use_petr_reply"
}
func (*ShowLispUsePetrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowLispUsePetrReply) GetCrcString() string {
	return "8e9360a8"
}
func NewShowLispUsePetrReply() api.Message {
	return &ShowLispUsePetrReply{}
}

// ShowLispRlocProbeState represents the VPP binary API message 'show_lisp_rloc_probe_state'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 167:
//
//        ["show_lisp_rloc_probe_state",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xc0b0f08b"}
//        ],
//
type ShowLispRlocProbeState struct {
}

func (*ShowLispRlocProbeState) GetMessageName() string {
	return "show_lisp_rloc_probe_state"
}
func (*ShowLispRlocProbeState) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowLispRlocProbeState) GetCrcString() string {
	return "c0b0f08b"
}
func NewShowLispRlocProbeState() api.Message {
	return &ShowLispRlocProbeState{}
}

// ShowLispRlocProbeStateReply represents the VPP binary API message 'show_lisp_rloc_probe_state_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 173:
//
//        ["show_lisp_rloc_probe_state_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "is_enabled"],
//            {"crc" : "0x23a1e712"}
//        ],
//
type ShowLispRlocProbeStateReply struct {
	Retval    int32
	IsEnabled uint8
}

func (*ShowLispRlocProbeStateReply) GetMessageName() string {
	return "show_lisp_rloc_probe_state_reply"
}
func (*ShowLispRlocProbeStateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowLispRlocProbeStateReply) GetCrcString() string {
	return "23a1e712"
}
func NewShowLispRlocProbeStateReply() api.Message {
	return &ShowLispRlocProbeStateReply{}
}

// LispRlocProbeEnableDisable represents the VPP binary API message 'lisp_rloc_probe_enable_disable'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 180:
//
//        ["lisp_rloc_probe_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_enabled"],
//            {"crc" : "0x32c270ac"}
//        ],
//
type LispRlocProbeEnableDisable struct {
	IsEnabled uint8
}

func (*LispRlocProbeEnableDisable) GetMessageName() string {
	return "lisp_rloc_probe_enable_disable"
}
func (*LispRlocProbeEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispRlocProbeEnableDisable) GetCrcString() string {
	return "32c270ac"
}
func NewLispRlocProbeEnableDisable() api.Message {
	return &LispRlocProbeEnableDisable{}
}

// LispRlocProbeEnableDisableReply represents the VPP binary API message 'lisp_rloc_probe_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 187:
//
//        ["lisp_rloc_probe_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x97d05bc4"}
//        ],
//
type LispRlocProbeEnableDisableReply struct {
	Retval int32
}

func (*LispRlocProbeEnableDisableReply) GetMessageName() string {
	return "lisp_rloc_probe_enable_disable_reply"
}
func (*LispRlocProbeEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispRlocProbeEnableDisableReply) GetCrcString() string {
	return "97d05bc4"
}
func NewLispRlocProbeEnableDisableReply() api.Message {
	return &LispRlocProbeEnableDisableReply{}
}

// LispMapRegisterEnableDisable represents the VPP binary API message 'lisp_map_register_enable_disable'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 193:
//
//        ["lisp_map_register_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_enabled"],
//            {"crc" : "0x8d0a81ca"}
//        ],
//
type LispMapRegisterEnableDisable struct {
	IsEnabled uint8
}

func (*LispMapRegisterEnableDisable) GetMessageName() string {
	return "lisp_map_register_enable_disable"
}
func (*LispMapRegisterEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispMapRegisterEnableDisable) GetCrcString() string {
	return "8d0a81ca"
}
func NewLispMapRegisterEnableDisable() api.Message {
	return &LispMapRegisterEnableDisable{}
}

// LispMapRegisterEnableDisableReply represents the VPP binary API message 'lisp_map_register_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 200:
//
//        ["lisp_map_register_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x99e7a700"}
//        ],
//
type LispMapRegisterEnableDisableReply struct {
	Retval int32
}

func (*LispMapRegisterEnableDisableReply) GetMessageName() string {
	return "lisp_map_register_enable_disable_reply"
}
func (*LispMapRegisterEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispMapRegisterEnableDisableReply) GetCrcString() string {
	return "99e7a700"
}
func NewLispMapRegisterEnableDisableReply() api.Message {
	return &LispMapRegisterEnableDisableReply{}
}

// ShowLispMapRegisterState represents the VPP binary API message 'show_lisp_map_register_state'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 206:
//
//        ["show_lisp_map_register_state",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x55fc9581"}
//        ],
//
type ShowLispMapRegisterState struct {
}

func (*ShowLispMapRegisterState) GetMessageName() string {
	return "show_lisp_map_register_state"
}
func (*ShowLispMapRegisterState) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowLispMapRegisterState) GetCrcString() string {
	return "55fc9581"
}
func NewShowLispMapRegisterState() api.Message {
	return &ShowLispMapRegisterState{}
}

// ShowLispMapRegisterStateReply represents the VPP binary API message 'show_lisp_map_register_state_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 212:
//
//        ["show_lisp_map_register_state_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "is_enabled"],
//            {"crc" : "0x8d04052e"}
//        ],
//
type ShowLispMapRegisterStateReply struct {
	Retval    int32
	IsEnabled uint8
}

func (*ShowLispMapRegisterStateReply) GetMessageName() string {
	return "show_lisp_map_register_state_reply"
}
func (*ShowLispMapRegisterStateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowLispMapRegisterStateReply) GetCrcString() string {
	return "8d04052e"
}
func NewShowLispMapRegisterStateReply() api.Message {
	return &ShowLispMapRegisterStateReply{}
}

// LispMapRequestMode represents the VPP binary API message 'lisp_map_request_mode'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 219:
//
//        ["lisp_map_request_mode",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "mode"],
//            {"crc" : "0xd204de7c"}
//        ],
//
type LispMapRequestMode struct {
	Mode uint8
}

func (*LispMapRequestMode) GetMessageName() string {
	return "lisp_map_request_mode"
}
func (*LispMapRequestMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispMapRequestMode) GetCrcString() string {
	return "d204de7c"
}
func NewLispMapRequestMode() api.Message {
	return &LispMapRequestMode{}
}

// LispMapRequestModeReply represents the VPP binary API message 'lisp_map_request_mode_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 226:
//
//        ["lisp_map_request_mode_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x930edf9f"}
//        ],
//
type LispMapRequestModeReply struct {
	Retval int32
}

func (*LispMapRequestModeReply) GetMessageName() string {
	return "lisp_map_request_mode_reply"
}
func (*LispMapRequestModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispMapRequestModeReply) GetCrcString() string {
	return "930edf9f"
}
func NewLispMapRequestModeReply() api.Message {
	return &LispMapRequestModeReply{}
}

// ShowLispMapRequestMode represents the VPP binary API message 'show_lisp_map_request_mode'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 232:
//
//        ["show_lisp_map_request_mode",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xda78d18a"}
//        ],
//
type ShowLispMapRequestMode struct {
}

func (*ShowLispMapRequestMode) GetMessageName() string {
	return "show_lisp_map_request_mode"
}
func (*ShowLispMapRequestMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowLispMapRequestMode) GetCrcString() string {
	return "da78d18a"
}
func NewShowLispMapRequestMode() api.Message {
	return &ShowLispMapRequestMode{}
}

// ShowLispMapRequestModeReply represents the VPP binary API message 'show_lisp_map_request_mode_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 238:
//
//        ["show_lisp_map_request_mode_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "mode"],
//            {"crc" : "0xed3aadef"}
//        ],
//
type ShowLispMapRequestModeReply struct {
	Retval int32
	Mode   uint8
}

func (*ShowLispMapRequestModeReply) GetMessageName() string {
	return "show_lisp_map_request_mode_reply"
}
func (*ShowLispMapRequestModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowLispMapRequestModeReply) GetCrcString() string {
	return "ed3aadef"
}
func NewShowLispMapRequestModeReply() api.Message {
	return &ShowLispMapRequestModeReply{}
}

// LispAddDelRemoteMapping represents the VPP binary API message 'lisp_add_del_remote_mapping'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 245:
//
//        ["lisp_add_del_remote_mapping",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_src_dst"],
//            ["u8", "del_all"],
//            ["u32", "vni"],
//            ["u8", "action"],
//            ["u8", "eid_type"],
//            ["u8", "eid", 16],
//            ["u8", "eid_len"],
//            ["u8", "seid", 16],
//            ["u8", "seid_len"],
//            ["u32", "rloc_num"],
//            ["vl_api_remote_locator_t", "rlocs", 0, "rloc_num"],
//            {"crc" : "0x833c53f9"}
//        ],
//
type LispAddDelRemoteMapping struct {
	IsAdd    uint8
	IsSrcDst uint8
	DelAll   uint8
	Vni      uint32
	Action   uint8
	EidType  uint8
	Eid      []byte `struc:"[16]byte"`
	EidLen   uint8
	Seid     []byte `struc:"[16]byte"`
	SeidLen  uint8
	RlocNum  uint32 `struc:"sizeof=Rlocs"`
	Rlocs    []RemoteLocator
}

func (*LispAddDelRemoteMapping) GetMessageName() string {
	return "lisp_add_del_remote_mapping"
}
func (*LispAddDelRemoteMapping) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelRemoteMapping) GetCrcString() string {
	return "833c53f9"
}
func NewLispAddDelRemoteMapping() api.Message {
	return &LispAddDelRemoteMapping{}
}

// LispAddDelRemoteMappingReply represents the VPP binary API message 'lisp_add_del_remote_mapping_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 263:
//
//        ["lisp_add_del_remote_mapping_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x4cae72c9"}
//        ],
//
type LispAddDelRemoteMappingReply struct {
	Retval int32
}

func (*LispAddDelRemoteMappingReply) GetMessageName() string {
	return "lisp_add_del_remote_mapping_reply"
}
func (*LispAddDelRemoteMappingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelRemoteMappingReply) GetCrcString() string {
	return "4cae72c9"
}
func NewLispAddDelRemoteMappingReply() api.Message {
	return &LispAddDelRemoteMappingReply{}
}

// LispAddDelAdjacency represents the VPP binary API message 'lisp_add_del_adjacency'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 269:
//
//        ["lisp_add_del_adjacency",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "vni"],
//            ["u8", "eid_type"],
//            ["u8", "reid", 16],
//            ["u8", "leid", 16],
//            ["u8", "reid_len"],
//            ["u8", "leid_len"],
//            {"crc" : "0x2bbefa02"}
//        ],
//
type LispAddDelAdjacency struct {
	IsAdd   uint8
	Vni     uint32
	EidType uint8
	Reid    []byte `struc:"[16]byte"`
	Leid    []byte `struc:"[16]byte"`
	ReidLen uint8
	LeidLen uint8
}

func (*LispAddDelAdjacency) GetMessageName() string {
	return "lisp_add_del_adjacency"
}
func (*LispAddDelAdjacency) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelAdjacency) GetCrcString() string {
	return "2bbefa02"
}
func NewLispAddDelAdjacency() api.Message {
	return &LispAddDelAdjacency{}
}

// LispAddDelAdjacencyReply represents the VPP binary API message 'lisp_add_del_adjacency_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 282:
//
//        ["lisp_add_del_adjacency_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x4628e1a8"}
//        ],
//
type LispAddDelAdjacencyReply struct {
	Retval int32
}

func (*LispAddDelAdjacencyReply) GetMessageName() string {
	return "lisp_add_del_adjacency_reply"
}
func (*LispAddDelAdjacencyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelAdjacencyReply) GetCrcString() string {
	return "4628e1a8"
}
func NewLispAddDelAdjacencyReply() api.Message {
	return &LispAddDelAdjacencyReply{}
}

// LispAddDelMapRequestItrRlocs represents the VPP binary API message 'lisp_add_del_map_request_itr_rlocs'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 288:
//
//        ["lisp_add_del_map_request_itr_rlocs",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "locator_set_name", 64],
//            {"crc" : "0x3376a927"}
//        ],
//
type LispAddDelMapRequestItrRlocs struct {
	IsAdd          uint8
	LocatorSetName []byte `struc:"[64]byte"`
}

func (*LispAddDelMapRequestItrRlocs) GetMessageName() string {
	return "lisp_add_del_map_request_itr_rlocs"
}
func (*LispAddDelMapRequestItrRlocs) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAddDelMapRequestItrRlocs) GetCrcString() string {
	return "3376a927"
}
func NewLispAddDelMapRequestItrRlocs() api.Message {
	return &LispAddDelMapRequestItrRlocs{}
}

// LispAddDelMapRequestItrRlocsReply represents the VPP binary API message 'lisp_add_del_map_request_itr_rlocs_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 296:
//
//        ["lisp_add_del_map_request_itr_rlocs_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x712498b4"}
//        ],
//
type LispAddDelMapRequestItrRlocsReply struct {
	Retval int32
}

func (*LispAddDelMapRequestItrRlocsReply) GetMessageName() string {
	return "lisp_add_del_map_request_itr_rlocs_reply"
}
func (*LispAddDelMapRequestItrRlocsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAddDelMapRequestItrRlocsReply) GetCrcString() string {
	return "712498b4"
}
func NewLispAddDelMapRequestItrRlocsReply() api.Message {
	return &LispAddDelMapRequestItrRlocsReply{}
}

// LispEidTableAddDelMap represents the VPP binary API message 'lisp_eid_table_add_del_map'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 302:
//
//        ["lisp_eid_table_add_del_map",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "vni"],
//            ["u32", "dp_table"],
//            ["u8", "is_l2"],
//            {"crc" : "0x598a91ce"}
//        ],
//
type LispEidTableAddDelMap struct {
	IsAdd   uint8
	Vni     uint32
	DpTable uint32
	IsL2    uint8
}

func (*LispEidTableAddDelMap) GetMessageName() string {
	return "lisp_eid_table_add_del_map"
}
func (*LispEidTableAddDelMap) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispEidTableAddDelMap) GetCrcString() string {
	return "598a91ce"
}
func NewLispEidTableAddDelMap() api.Message {
	return &LispEidTableAddDelMap{}
}

// LispEidTableAddDelMapReply represents the VPP binary API message 'lisp_eid_table_add_del_map_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 312:
//
//        ["lisp_eid_table_add_del_map_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x9c948155"}
//        ],
//
type LispEidTableAddDelMapReply struct {
	Retval int32
}

func (*LispEidTableAddDelMapReply) GetMessageName() string {
	return "lisp_eid_table_add_del_map_reply"
}
func (*LispEidTableAddDelMapReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispEidTableAddDelMapReply) GetCrcString() string {
	return "9c948155"
}
func NewLispEidTableAddDelMapReply() api.Message {
	return &LispEidTableAddDelMapReply{}
}

// LispLocatorDump represents the VPP binary API message 'lisp_locator_dump'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 318:
//
//        ["lisp_locator_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "ls_index"],
//            ["u8", "ls_name", 64],
//            ["u8", "is_index_set"],
//            {"crc" : "0x35176bc8"}
//        ],
//
type LispLocatorDump struct {
	LsIndex    uint32
	LsName     []byte `struc:"[64]byte"`
	IsIndexSet uint8
}

func (*LispLocatorDump) GetMessageName() string {
	return "lisp_locator_dump"
}
func (*LispLocatorDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispLocatorDump) GetCrcString() string {
	return "35176bc8"
}
func NewLispLocatorDump() api.Message {
	return &LispLocatorDump{}
}

// LispLocatorDetails represents the VPP binary API message 'lisp_locator_details'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 327:
//
//        ["lisp_locator_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "local"],
//            ["u32", "sw_if_index"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            {"crc" : "0xcb282c00"}
//        ],
//
type LispLocatorDetails struct {
	Local     uint8
	SwIfIndex uint32
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
	Priority  uint8
	Weight    uint8
}

func (*LispLocatorDetails) GetMessageName() string {
	return "lisp_locator_details"
}
func (*LispLocatorDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispLocatorDetails) GetCrcString() string {
	return "cb282c00"
}
func NewLispLocatorDetails() api.Message {
	return &LispLocatorDetails{}
}

// LispLocatorSetDetails represents the VPP binary API message 'lisp_locator_set_details'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 338:
//
//        ["lisp_locator_set_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "ls_index"],
//            ["u8", "ls_name", 64],
//            {"crc" : "0x4ab2d4cf"}
//        ],
//
type LispLocatorSetDetails struct {
	LsIndex uint32
	LsName  []byte `struc:"[64]byte"`
}

func (*LispLocatorSetDetails) GetMessageName() string {
	return "lisp_locator_set_details"
}
func (*LispLocatorSetDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispLocatorSetDetails) GetCrcString() string {
	return "4ab2d4cf"
}
func NewLispLocatorSetDetails() api.Message {
	return &LispLocatorSetDetails{}
}

// LispLocatorSetDump represents the VPP binary API message 'lisp_locator_set_dump'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 345:
//
//        ["lisp_locator_set_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "filter"],
//            {"crc" : "0x0f3d315b"}
//        ],
//
type LispLocatorSetDump struct {
	Filter uint8
}

func (*LispLocatorSetDump) GetMessageName() string {
	return "lisp_locator_set_dump"
}
func (*LispLocatorSetDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispLocatorSetDump) GetCrcString() string {
	return "0f3d315b"
}
func NewLispLocatorSetDump() api.Message {
	return &LispLocatorSetDump{}
}

// LispEidTableDetails represents the VPP binary API message 'lisp_eid_table_details'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 352:
//
//        ["lisp_eid_table_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "locator_set_index"],
//            ["u8", "action"],
//            ["u8", "is_local"],
//            ["u8", "eid_type"],
//            ["u8", "is_src_dst"],
//            ["u32", "vni"],
//            ["u8", "eid", 16],
//            ["u8", "eid_prefix_len"],
//            ["u8", "seid", 16],
//            ["u8", "seid_prefix_len"],
//            ["u32", "ttl"],
//            ["u8", "authoritative"],
//            ["u16", "key_id"],
//            ["u8", "key", 64],
//            {"crc" : "0xb93cde6b"}
//        ],
//
type LispEidTableDetails struct {
	LocatorSetIndex uint32
	Action          uint8
	IsLocal         uint8
	EidType         uint8
	IsSrcDst        uint8
	Vni             uint32
	Eid             []byte `struc:"[16]byte"`
	EidPrefixLen    uint8
	Seid            []byte `struc:"[16]byte"`
	SeidPrefixLen   uint8
	TTL             uint32
	Authoritative   uint8
	KeyID           uint16
	Key             []byte `struc:"[64]byte"`
}

func (*LispEidTableDetails) GetMessageName() string {
	return "lisp_eid_table_details"
}
func (*LispEidTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispEidTableDetails) GetCrcString() string {
	return "b93cde6b"
}
func NewLispEidTableDetails() api.Message {
	return &LispEidTableDetails{}
}

// LispEidTableDump represents the VPP binary API message 'lisp_eid_table_dump'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 371:
//
//        ["lisp_eid_table_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "eid_set"],
//            ["u8", "prefix_length"],
//            ["u32", "vni"],
//            ["u8", "eid_type"],
//            ["u8", "eid", 16],
//            ["u8", "filter"],
//            {"crc" : "0x354e0c1a"}
//        ],
//
type LispEidTableDump struct {
	EidSet       uint8
	PrefixLength uint8
	Vni          uint32
	EidType      uint8
	Eid          []byte `struc:"[16]byte"`
	Filter       uint8
}

func (*LispEidTableDump) GetMessageName() string {
	return "lisp_eid_table_dump"
}
func (*LispEidTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispEidTableDump) GetCrcString() string {
	return "354e0c1a"
}
func NewLispEidTableDump() api.Message {
	return &LispEidTableDump{}
}

// LispAdjacenciesGetReply represents the VPP binary API message 'lisp_adjacencies_get_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 383:
//
//        ["lisp_adjacencies_get_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "count"],
//            ["vl_api_lisp_adjacency_t", "adjacencies", 0, "count"],
//            {"crc" : "0x00dcfe1d"}
//        ],
//
type LispAdjacenciesGetReply struct {
	Retval      int32
	Count       uint32 `struc:"sizeof=Adjacencies"`
	Adjacencies []LispAdjacency
}

func (*LispAdjacenciesGetReply) GetMessageName() string {
	return "lisp_adjacencies_get_reply"
}
func (*LispAdjacenciesGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispAdjacenciesGetReply) GetCrcString() string {
	return "00dcfe1d"
}
func NewLispAdjacenciesGetReply() api.Message {
	return &LispAdjacenciesGetReply{}
}

// LispAdjacenciesGet represents the VPP binary API message 'lisp_adjacencies_get'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 391:
//
//        ["lisp_adjacencies_get",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            {"crc" : "0xf0252c92"}
//        ],
//
type LispAdjacenciesGet struct {
	Vni uint32
}

func (*LispAdjacenciesGet) GetMessageName() string {
	return "lisp_adjacencies_get"
}
func (*LispAdjacenciesGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispAdjacenciesGet) GetCrcString() string {
	return "f0252c92"
}
func NewLispAdjacenciesGet() api.Message {
	return &LispAdjacenciesGet{}
}

// LispEidTableMapDetails represents the VPP binary API message 'lisp_eid_table_map_details'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 398:
//
//        ["lisp_eid_table_map_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            ["u32", "dp_table"],
//            {"crc" : "0xc5f081e9"}
//        ],
//
type LispEidTableMapDetails struct {
	Vni     uint32
	DpTable uint32
}

func (*LispEidTableMapDetails) GetMessageName() string {
	return "lisp_eid_table_map_details"
}
func (*LispEidTableMapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispEidTableMapDetails) GetCrcString() string {
	return "c5f081e9"
}
func NewLispEidTableMapDetails() api.Message {
	return &LispEidTableMapDetails{}
}

// LispEidTableMapDump represents the VPP binary API message 'lisp_eid_table_map_dump'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 405:
//
//        ["lisp_eid_table_map_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_l2"],
//            {"crc" : "0xb0704823"}
//        ],
//
type LispEidTableMapDump struct {
	IsL2 uint8
}

func (*LispEidTableMapDump) GetMessageName() string {
	return "lisp_eid_table_map_dump"
}
func (*LispEidTableMapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispEidTableMapDump) GetCrcString() string {
	return "b0704823"
}
func NewLispEidTableMapDump() api.Message {
	return &LispEidTableMapDump{}
}

// LispEidTableVniDump represents the VPP binary API message 'lisp_eid_table_vni_dump'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 412:
//
//        ["lisp_eid_table_vni_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x3456e06a"}
//        ],
//
type LispEidTableVniDump struct {
}

func (*LispEidTableVniDump) GetMessageName() string {
	return "lisp_eid_table_vni_dump"
}
func (*LispEidTableVniDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispEidTableVniDump) GetCrcString() string {
	return "3456e06a"
}
func NewLispEidTableVniDump() api.Message {
	return &LispEidTableVniDump{}
}

// LispEidTableVniDetails represents the VPP binary API message 'lisp_eid_table_vni_details'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 418:
//
//        ["lisp_eid_table_vni_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            {"crc" : "0xe2f8a8b9"}
//        ],
//
type LispEidTableVniDetails struct {
	Vni uint32
}

func (*LispEidTableVniDetails) GetMessageName() string {
	return "lisp_eid_table_vni_details"
}
func (*LispEidTableVniDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispEidTableVniDetails) GetCrcString() string {
	return "e2f8a8b9"
}
func NewLispEidTableVniDetails() api.Message {
	return &LispEidTableVniDetails{}
}

// LispMapResolverDetails represents the VPP binary API message 'lisp_map_resolver_details'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 425:
//
//        ["lisp_map_resolver_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0xe8c68ebd"}
//        ],
//
type LispMapResolverDetails struct {
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*LispMapResolverDetails) GetMessageName() string {
	return "lisp_map_resolver_details"
}
func (*LispMapResolverDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispMapResolverDetails) GetCrcString() string {
	return "e8c68ebd"
}
func NewLispMapResolverDetails() api.Message {
	return &LispMapResolverDetails{}
}

// LispMapResolverDump represents the VPP binary API message 'lisp_map_resolver_dump'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 432:
//
//        ["lisp_map_resolver_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x4e5e2003"}
//        ],
//
type LispMapResolverDump struct {
}

func (*LispMapResolverDump) GetMessageName() string {
	return "lisp_map_resolver_dump"
}
func (*LispMapResolverDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispMapResolverDump) GetCrcString() string {
	return "4e5e2003"
}
func NewLispMapResolverDump() api.Message {
	return &LispMapResolverDump{}
}

// LispMapServerDetails represents the VPP binary API message 'lisp_map_server_details'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 438:
//
//        ["lisp_map_server_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0x4ef38e5a"}
//        ],
//
type LispMapServerDetails struct {
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*LispMapServerDetails) GetMessageName() string {
	return "lisp_map_server_details"
}
func (*LispMapServerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispMapServerDetails) GetCrcString() string {
	return "4ef38e5a"
}
func NewLispMapServerDetails() api.Message {
	return &LispMapServerDetails{}
}

// LispMapServerDump represents the VPP binary API message 'lisp_map_server_dump'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 445:
//
//        ["lisp_map_server_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x2b2998e2"}
//        ],
//
type LispMapServerDump struct {
}

func (*LispMapServerDump) GetMessageName() string {
	return "lisp_map_server_dump"
}
func (*LispMapServerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispMapServerDump) GetCrcString() string {
	return "2b2998e2"
}
func NewLispMapServerDump() api.Message {
	return &LispMapServerDump{}
}

// ShowLispStatus represents the VPP binary API message 'show_lisp_status'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 451:
//
//        ["show_lisp_status",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x8092ab77"}
//        ],
//
type ShowLispStatus struct {
}

func (*ShowLispStatus) GetMessageName() string {
	return "show_lisp_status"
}
func (*ShowLispStatus) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowLispStatus) GetCrcString() string {
	return "8092ab77"
}
func NewShowLispStatus() api.Message {
	return &ShowLispStatus{}
}

// ShowLispStatusReply represents the VPP binary API message 'show_lisp_status_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 457:
//
//        ["show_lisp_status_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "feature_status"],
//            ["u8", "gpe_status"],
//            {"crc" : "0x6aa3f21d"}
//        ],
//
type ShowLispStatusReply struct {
	Retval        int32
	FeatureStatus uint8
	GpeStatus     uint8
}

func (*ShowLispStatusReply) GetMessageName() string {
	return "show_lisp_status_reply"
}
func (*ShowLispStatusReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowLispStatusReply) GetCrcString() string {
	return "6aa3f21d"
}
func NewShowLispStatusReply() api.Message {
	return &ShowLispStatusReply{}
}

// LispGetMapRequestItrRlocs represents the VPP binary API message 'lisp_get_map_request_itr_rlocs'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 465:
//
//        ["lisp_get_map_request_itr_rlocs",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x1e2d23a4"}
//        ],
//
type LispGetMapRequestItrRlocs struct {
}

func (*LispGetMapRequestItrRlocs) GetMessageName() string {
	return "lisp_get_map_request_itr_rlocs"
}
func (*LispGetMapRequestItrRlocs) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*LispGetMapRequestItrRlocs) GetCrcString() string {
	return "1e2d23a4"
}
func NewLispGetMapRequestItrRlocs() api.Message {
	return &LispGetMapRequestItrRlocs{}
}

// LispGetMapRequestItrRlocsReply represents the VPP binary API message 'lisp_get_map_request_itr_rlocs_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 471:
//
//        ["lisp_get_map_request_itr_rlocs_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "locator_set_name", 64],
//            {"crc" : "0x39bfca79"}
//        ],
//
type LispGetMapRequestItrRlocsReply struct {
	Retval         int32
	LocatorSetName []byte `struc:"[64]byte"`
}

func (*LispGetMapRequestItrRlocsReply) GetMessageName() string {
	return "lisp_get_map_request_itr_rlocs_reply"
}
func (*LispGetMapRequestItrRlocsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*LispGetMapRequestItrRlocsReply) GetCrcString() string {
	return "39bfca79"
}
func NewLispGetMapRequestItrRlocsReply() api.Message {
	return &LispGetMapRequestItrRlocsReply{}
}

// ShowLispPitr represents the VPP binary API message 'show_lisp_pitr'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 478:
//
//        ["show_lisp_pitr",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xd4a061e6"}
//        ],
//
type ShowLispPitr struct {
}

func (*ShowLispPitr) GetMessageName() string {
	return "show_lisp_pitr"
}
func (*ShowLispPitr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowLispPitr) GetCrcString() string {
	return "d4a061e6"
}
func NewShowLispPitr() api.Message {
	return &ShowLispPitr{}
}

// ShowLispPitrReply represents the VPP binary API message 'show_lisp_pitr_reply'.
// Generated from '/usr/share/vpp/api//lisp.api.json', line 484:
//
//        ["show_lisp_pitr_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "status"],
//            ["u8", "locator_set_name", 64],
//            {"crc" : "0xe730f16e"}
//        ]
//
type ShowLispPitrReply struct {
	Retval         int32
	Status         uint8
	LocatorSetName []byte `struc:"[64]byte"`
}

func (*ShowLispPitrReply) GetMessageName() string {
	return "show_lisp_pitr_reply"
}
func (*ShowLispPitrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowLispPitrReply) GetCrcString() string {
	return "e730f16e"
}
func NewShowLispPitrReply() api.Message {
	return &ShowLispPitrReply{}
}
