// Package one represents the VPP binary API of the 'one' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//one.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package one

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xcf44a7fd

// OneLocalLocator represents the VPP binary API data type 'one_local_locator'.
// Generated from '/usr/share/vpp/api//one.api.json', line 3:
//
//        ["one_local_locator",
//            ["u32", "sw_if_index"],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            {"crc" : "0x3227bcd5"}
//        ],
//
type OneLocalLocator struct {
	SwIfIndex uint32
	Priority  uint8
	Weight    uint8
}

func (*OneLocalLocator) GetTypeName() string {
	return "one_local_locator"
}
func (*OneLocalLocator) GetCrcString() string {
	return "3227bcd5"
}

// OneRemoteLocator represents the VPP binary API data type 'one_remote_locator'.
// Generated from '/usr/share/vpp/api//one.api.json', line 9:
//
//        ["one_remote_locator",
//            ["u8", "is_ip4"],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            ["u8", "addr", 16],
//            {"crc" : "0xce284680"}
//        ],
//
type OneRemoteLocator struct {
	IsIP4    uint8
	Priority uint8
	Weight   uint8
	Addr     []byte `struc:"[16]byte"`
}

func (*OneRemoteLocator) GetTypeName() string {
	return "one_remote_locator"
}
func (*OneRemoteLocator) GetCrcString() string {
	return "ce284680"
}

// OneAdjacency represents the VPP binary API data type 'one_adjacency'.
// Generated from '/usr/share/vpp/api//one.api.json', line 16:
//
//        ["one_adjacency",
//            ["u8", "eid_type"],
//            ["u8", "reid", 16],
//            ["u8", "leid", 16],
//            ["u8", "reid_prefix_len"],
//            ["u8", "leid_prefix_len"],
//            {"crc" : "0x035c4f6e"}
//        ]
//
type OneAdjacency struct {
	EidType       uint8
	Reid          []byte `struc:"[16]byte"`
	Leid          []byte `struc:"[16]byte"`
	ReidPrefixLen uint8
	LeidPrefixLen uint8
}

func (*OneAdjacency) GetTypeName() string {
	return "one_adjacency"
}
func (*OneAdjacency) GetCrcString() string {
	return "035c4f6e"
}

// OneAddDelLocatorSet represents the VPP binary API message 'one_add_del_locator_set'.
// Generated from '/usr/share/vpp/api//one.api.json', line 26:
//
//        ["one_add_del_locator_set",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "locator_set_name", 64],
//            ["u32", "locator_num"],
//            ["vl_api_one_local_locator_t", "locators", 0, "locator_num"],
//            {"crc" : "0x57cc6026"}
//        ],
//
type OneAddDelLocatorSet struct {
	IsAdd          uint8
	LocatorSetName []byte `struc:"[64]byte"`
	LocatorNum     uint32 `struc:"sizeof=Locators"`
	Locators       []OneLocalLocator
}

func (*OneAddDelLocatorSet) GetMessageName() string {
	return "one_add_del_locator_set"
}
func (*OneAddDelLocatorSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelLocatorSet) GetCrcString() string {
	return "57cc6026"
}
func NewOneAddDelLocatorSet() api.Message {
	return &OneAddDelLocatorSet{}
}

// OneAddDelLocatorSetReply represents the VPP binary API message 'one_add_del_locator_set_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 36:
//
//        ["one_add_del_locator_set_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "ls_index"],
//            {"crc" : "0xc6af12ec"}
//        ],
//
type OneAddDelLocatorSetReply struct {
	Retval  int32
	LsIndex uint32
}

func (*OneAddDelLocatorSetReply) GetMessageName() string {
	return "one_add_del_locator_set_reply"
}
func (*OneAddDelLocatorSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelLocatorSetReply) GetCrcString() string {
	return "c6af12ec"
}
func NewOneAddDelLocatorSetReply() api.Message {
	return &OneAddDelLocatorSetReply{}
}

// OneAddDelLocator represents the VPP binary API message 'one_add_del_locator'.
// Generated from '/usr/share/vpp/api//one.api.json', line 43:
//
//        ["one_add_del_locator",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "locator_set_name", 64],
//            ["u32", "sw_if_index"],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            {"crc" : "0xfc816f2d"}
//        ],
//
type OneAddDelLocator struct {
	IsAdd          uint8
	LocatorSetName []byte `struc:"[64]byte"`
	SwIfIndex      uint32
	Priority       uint8
	Weight         uint8
}

func (*OneAddDelLocator) GetMessageName() string {
	return "one_add_del_locator"
}
func (*OneAddDelLocator) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelLocator) GetCrcString() string {
	return "fc816f2d"
}
func NewOneAddDelLocator() api.Message {
	return &OneAddDelLocator{}
}

// OneAddDelLocatorReply represents the VPP binary API message 'one_add_del_locator_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 54:
//
//        ["one_add_del_locator_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x72b93e34"}
//        ],
//
type OneAddDelLocatorReply struct {
	Retval int32
}

func (*OneAddDelLocatorReply) GetMessageName() string {
	return "one_add_del_locator_reply"
}
func (*OneAddDelLocatorReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelLocatorReply) GetCrcString() string {
	return "72b93e34"
}
func NewOneAddDelLocatorReply() api.Message {
	return &OneAddDelLocatorReply{}
}

// OneAddDelLocalEid represents the VPP binary API message 'one_add_del_local_eid'.
// Generated from '/usr/share/vpp/api//one.api.json', line 60:
//
//        ["one_add_del_local_eid",
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
//            {"crc" : "0x9fa2e1b7"}
//        ],
//
type OneAddDelLocalEid struct {
	IsAdd          uint8
	EidType        uint8
	Eid            []byte `struc:"[16]byte"`
	PrefixLen      uint8
	LocatorSetName []byte `struc:"[64]byte"`
	Vni            uint32
	KeyID          uint16
	Key            []byte `struc:"[64]byte"`
}

func (*OneAddDelLocalEid) GetMessageName() string {
	return "one_add_del_local_eid"
}
func (*OneAddDelLocalEid) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelLocalEid) GetCrcString() string {
	return "9fa2e1b7"
}
func NewOneAddDelLocalEid() api.Message {
	return &OneAddDelLocalEid{}
}

// OneAddDelLocalEidReply represents the VPP binary API message 'one_add_del_local_eid_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 74:
//
//        ["one_add_del_local_eid_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb0e4330b"}
//        ],
//
type OneAddDelLocalEidReply struct {
	Retval int32
}

func (*OneAddDelLocalEidReply) GetMessageName() string {
	return "one_add_del_local_eid_reply"
}
func (*OneAddDelLocalEidReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelLocalEidReply) GetCrcString() string {
	return "b0e4330b"
}
func NewOneAddDelLocalEidReply() api.Message {
	return &OneAddDelLocalEidReply{}
}

// OneAddDelMapServer represents the VPP binary API message 'one_add_del_map_server'.
// Generated from '/usr/share/vpp/api//one.api.json', line 80:
//
//        ["one_add_del_map_server",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0x731cfa38"}
//        ],
//
type OneAddDelMapServer struct {
	IsAdd     uint8
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*OneAddDelMapServer) GetMessageName() string {
	return "one_add_del_map_server"
}
func (*OneAddDelMapServer) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelMapServer) GetCrcString() string {
	return "731cfa38"
}
func NewOneAddDelMapServer() api.Message {
	return &OneAddDelMapServer{}
}

// OneAddDelMapServerReply represents the VPP binary API message 'one_add_del_map_server_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 89:
//
//        ["one_add_del_map_server_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xce8d6a33"}
//        ],
//
type OneAddDelMapServerReply struct {
	Retval int32
}

func (*OneAddDelMapServerReply) GetMessageName() string {
	return "one_add_del_map_server_reply"
}
func (*OneAddDelMapServerReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelMapServerReply) GetCrcString() string {
	return "ce8d6a33"
}
func NewOneAddDelMapServerReply() api.Message {
	return &OneAddDelMapServerReply{}
}

// OneAddDelMapResolver represents the VPP binary API message 'one_add_del_map_resolver'.
// Generated from '/usr/share/vpp/api//one.api.json', line 95:
//
//        ["one_add_del_map_resolver",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0xa627ab50"}
//        ],
//
type OneAddDelMapResolver struct {
	IsAdd     uint8
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*OneAddDelMapResolver) GetMessageName() string {
	return "one_add_del_map_resolver"
}
func (*OneAddDelMapResolver) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelMapResolver) GetCrcString() string {
	return "a627ab50"
}
func NewOneAddDelMapResolver() api.Message {
	return &OneAddDelMapResolver{}
}

// OneAddDelMapResolverReply represents the VPP binary API message 'one_add_del_map_resolver_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 104:
//
//        ["one_add_del_map_resolver_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xe46fb226"}
//        ],
//
type OneAddDelMapResolverReply struct {
	Retval int32
}

func (*OneAddDelMapResolverReply) GetMessageName() string {
	return "one_add_del_map_resolver_reply"
}
func (*OneAddDelMapResolverReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelMapResolverReply) GetCrcString() string {
	return "e46fb226"
}
func NewOneAddDelMapResolverReply() api.Message {
	return &OneAddDelMapResolverReply{}
}

// OneEnableDisable represents the VPP binary API message 'one_enable_disable'.
// Generated from '/usr/share/vpp/api//one.api.json', line 110:
//
//        ["one_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_en"],
//            {"crc" : "0x3486b168"}
//        ],
//
type OneEnableDisable struct {
	IsEn uint8
}

func (*OneEnableDisable) GetMessageName() string {
	return "one_enable_disable"
}
func (*OneEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneEnableDisable) GetCrcString() string {
	return "3486b168"
}
func NewOneEnableDisable() api.Message {
	return &OneEnableDisable{}
}

// OneEnableDisableReply represents the VPP binary API message 'one_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 117:
//
//        ["one_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x000bda40"}
//        ],
//
type OneEnableDisableReply struct {
	Retval int32
}

func (*OneEnableDisableReply) GetMessageName() string {
	return "one_enable_disable_reply"
}
func (*OneEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneEnableDisableReply) GetCrcString() string {
	return "000bda40"
}
func NewOneEnableDisableReply() api.Message {
	return &OneEnableDisableReply{}
}

// OnePitrSetLocatorSet represents the VPP binary API message 'one_pitr_set_locator_set'.
// Generated from '/usr/share/vpp/api//one.api.json', line 123:
//
//        ["one_pitr_set_locator_set",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "ls_name", 64],
//            {"crc" : "0x4b46f7e0"}
//        ],
//
type OnePitrSetLocatorSet struct {
	IsAdd  uint8
	LsName []byte `struc:"[64]byte"`
}

func (*OnePitrSetLocatorSet) GetMessageName() string {
	return "one_pitr_set_locator_set"
}
func (*OnePitrSetLocatorSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OnePitrSetLocatorSet) GetCrcString() string {
	return "4b46f7e0"
}
func NewOnePitrSetLocatorSet() api.Message {
	return &OnePitrSetLocatorSet{}
}

// OnePitrSetLocatorSetReply represents the VPP binary API message 'one_pitr_set_locator_set_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 131:
//
//        ["one_pitr_set_locator_set_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x78fb5604"}
//        ],
//
type OnePitrSetLocatorSetReply struct {
	Retval int32
}

func (*OnePitrSetLocatorSetReply) GetMessageName() string {
	return "one_pitr_set_locator_set_reply"
}
func (*OnePitrSetLocatorSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OnePitrSetLocatorSetReply) GetCrcString() string {
	return "78fb5604"
}
func NewOnePitrSetLocatorSetReply() api.Message {
	return &OnePitrSetLocatorSetReply{}
}

// OneUsePetr represents the VPP binary API message 'one_use_petr'.
// Generated from '/usr/share/vpp/api//one.api.json', line 137:
//
//        ["one_use_petr",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "address", 16],
//            ["u8", "is_add"],
//            {"crc" : "0x5d87b0c3"}
//        ],
//
type OneUsePetr struct {
	IsIP4   uint8
	Address []byte `struc:"[16]byte"`
	IsAdd   uint8
}

func (*OneUsePetr) GetMessageName() string {
	return "one_use_petr"
}
func (*OneUsePetr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneUsePetr) GetCrcString() string {
	return "5d87b0c3"
}
func NewOneUsePetr() api.Message {
	return &OneUsePetr{}
}

// OneUsePetrReply represents the VPP binary API message 'one_use_petr_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 146:
//
//        ["one_use_petr_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x20d789f4"}
//        ],
//
type OneUsePetrReply struct {
	Retval int32
}

func (*OneUsePetrReply) GetMessageName() string {
	return "one_use_petr_reply"
}
func (*OneUsePetrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneUsePetrReply) GetCrcString() string {
	return "20d789f4"
}
func NewOneUsePetrReply() api.Message {
	return &OneUsePetrReply{}
}

// ShowOneUsePetr represents the VPP binary API message 'show_one_use_petr'.
// Generated from '/usr/share/vpp/api//one.api.json', line 152:
//
//        ["show_one_use_petr",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x762fcd06"}
//        ],
//
type ShowOneUsePetr struct {
}

func (*ShowOneUsePetr) GetMessageName() string {
	return "show_one_use_petr"
}
func (*ShowOneUsePetr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowOneUsePetr) GetCrcString() string {
	return "762fcd06"
}
func NewShowOneUsePetr() api.Message {
	return &ShowOneUsePetr{}
}

// ShowOneUsePetrReply represents the VPP binary API message 'show_one_use_petr_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 158:
//
//        ["show_one_use_petr_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "status"],
//            ["u8", "is_ip4"],
//            ["u8", "address", 16],
//            {"crc" : "0xdf47b6f3"}
//        ],
//
type ShowOneUsePetrReply struct {
	Retval  int32
	Status  uint8
	IsIP4   uint8
	Address []byte `struc:"[16]byte"`
}

func (*ShowOneUsePetrReply) GetMessageName() string {
	return "show_one_use_petr_reply"
}
func (*ShowOneUsePetrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowOneUsePetrReply) GetCrcString() string {
	return "df47b6f3"
}
func NewShowOneUsePetrReply() api.Message {
	return &ShowOneUsePetrReply{}
}

// ShowOneRlocProbeState represents the VPP binary API message 'show_one_rloc_probe_state'.
// Generated from '/usr/share/vpp/api//one.api.json', line 167:
//
//        ["show_one_rloc_probe_state",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xc9274fad"}
//        ],
//
type ShowOneRlocProbeState struct {
}

func (*ShowOneRlocProbeState) GetMessageName() string {
	return "show_one_rloc_probe_state"
}
func (*ShowOneRlocProbeState) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowOneRlocProbeState) GetCrcString() string {
	return "c9274fad"
}
func NewShowOneRlocProbeState() api.Message {
	return &ShowOneRlocProbeState{}
}

// ShowOneRlocProbeStateReply represents the VPP binary API message 'show_one_rloc_probe_state_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 173:
//
//        ["show_one_rloc_probe_state_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "is_enabled"],
//            {"crc" : "0x0bb0ebf4"}
//        ],
//
type ShowOneRlocProbeStateReply struct {
	Retval    int32
	IsEnabled uint8
}

func (*ShowOneRlocProbeStateReply) GetMessageName() string {
	return "show_one_rloc_probe_state_reply"
}
func (*ShowOneRlocProbeStateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowOneRlocProbeStateReply) GetCrcString() string {
	return "0bb0ebf4"
}
func NewShowOneRlocProbeStateReply() api.Message {
	return &ShowOneRlocProbeStateReply{}
}

// OneRlocProbeEnableDisable represents the VPP binary API message 'one_rloc_probe_enable_disable'.
// Generated from '/usr/share/vpp/api//one.api.json', line 180:
//
//        ["one_rloc_probe_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_enabled"],
//            {"crc" : "0x6072518f"}
//        ],
//
type OneRlocProbeEnableDisable struct {
	IsEnabled uint8
}

func (*OneRlocProbeEnableDisable) GetMessageName() string {
	return "one_rloc_probe_enable_disable"
}
func (*OneRlocProbeEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneRlocProbeEnableDisable) GetCrcString() string {
	return "6072518f"
}
func NewOneRlocProbeEnableDisable() api.Message {
	return &OneRlocProbeEnableDisable{}
}

// OneRlocProbeEnableDisableReply represents the VPP binary API message 'one_rloc_probe_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 187:
//
//        ["one_rloc_probe_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xbdfc3a16"}
//        ],
//
type OneRlocProbeEnableDisableReply struct {
	Retval int32
}

func (*OneRlocProbeEnableDisableReply) GetMessageName() string {
	return "one_rloc_probe_enable_disable_reply"
}
func (*OneRlocProbeEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneRlocProbeEnableDisableReply) GetCrcString() string {
	return "bdfc3a16"
}
func NewOneRlocProbeEnableDisableReply() api.Message {
	return &OneRlocProbeEnableDisableReply{}
}

// OneMapRegisterEnableDisable represents the VPP binary API message 'one_map_register_enable_disable'.
// Generated from '/usr/share/vpp/api//one.api.json', line 193:
//
//        ["one_map_register_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_enabled"],
//            {"crc" : "0x14b0953c"}
//        ],
//
type OneMapRegisterEnableDisable struct {
	IsEnabled uint8
}

func (*OneMapRegisterEnableDisable) GetMessageName() string {
	return "one_map_register_enable_disable"
}
func (*OneMapRegisterEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneMapRegisterEnableDisable) GetCrcString() string {
	return "14b0953c"
}
func NewOneMapRegisterEnableDisable() api.Message {
	return &OneMapRegisterEnableDisable{}
}

// OneMapRegisterEnableDisableReply represents the VPP binary API message 'one_map_register_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 200:
//
//        ["one_map_register_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x161e60df"}
//        ],
//
type OneMapRegisterEnableDisableReply struct {
	Retval int32
}

func (*OneMapRegisterEnableDisableReply) GetMessageName() string {
	return "one_map_register_enable_disable_reply"
}
func (*OneMapRegisterEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneMapRegisterEnableDisableReply) GetCrcString() string {
	return "161e60df"
}
func NewOneMapRegisterEnableDisableReply() api.Message {
	return &OneMapRegisterEnableDisableReply{}
}

// ShowOneMapRegisterState represents the VPP binary API message 'show_one_map_register_state'.
// Generated from '/usr/share/vpp/api//one.api.json', line 206:
//
//        ["show_one_map_register_state",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xafbe60cb"}
//        ],
//
type ShowOneMapRegisterState struct {
}

func (*ShowOneMapRegisterState) GetMessageName() string {
	return "show_one_map_register_state"
}
func (*ShowOneMapRegisterState) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowOneMapRegisterState) GetCrcString() string {
	return "afbe60cb"
}
func NewShowOneMapRegisterState() api.Message {
	return &ShowOneMapRegisterState{}
}

// ShowOneMapRegisterStateReply represents the VPP binary API message 'show_one_map_register_state_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 212:
//
//        ["show_one_map_register_state_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "is_enabled"],
//            {"crc" : "0xd2d76c22"}
//        ],
//
type ShowOneMapRegisterStateReply struct {
	Retval    int32
	IsEnabled uint8
}

func (*ShowOneMapRegisterStateReply) GetMessageName() string {
	return "show_one_map_register_state_reply"
}
func (*ShowOneMapRegisterStateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowOneMapRegisterStateReply) GetCrcString() string {
	return "d2d76c22"
}
func NewShowOneMapRegisterStateReply() api.Message {
	return &ShowOneMapRegisterStateReply{}
}

// OneMapRequestMode represents the VPP binary API message 'one_map_request_mode'.
// Generated from '/usr/share/vpp/api//one.api.json', line 219:
//
//        ["one_map_request_mode",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "mode"],
//            {"crc" : "0x67508836"}
//        ],
//
type OneMapRequestMode struct {
	Mode uint8
}

func (*OneMapRequestMode) GetMessageName() string {
	return "one_map_request_mode"
}
func (*OneMapRequestMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneMapRequestMode) GetCrcString() string {
	return "67508836"
}
func NewOneMapRequestMode() api.Message {
	return &OneMapRequestMode{}
}

// OneMapRequestModeReply represents the VPP binary API message 'one_map_request_mode_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 226:
//
//        ["one_map_request_mode_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x25f76bca"}
//        ],
//
type OneMapRequestModeReply struct {
	Retval int32
}

func (*OneMapRequestModeReply) GetMessageName() string {
	return "one_map_request_mode_reply"
}
func (*OneMapRequestModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneMapRequestModeReply) GetCrcString() string {
	return "25f76bca"
}
func NewOneMapRequestModeReply() api.Message {
	return &OneMapRequestModeReply{}
}

// ShowOneMapRequestMode represents the VPP binary API message 'show_one_map_request_mode'.
// Generated from '/usr/share/vpp/api//one.api.json', line 232:
//
//        ["show_one_map_request_mode",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xd3ef6eac"}
//        ],
//
type ShowOneMapRequestMode struct {
}

func (*ShowOneMapRequestMode) GetMessageName() string {
	return "show_one_map_request_mode"
}
func (*ShowOneMapRequestMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowOneMapRequestMode) GetCrcString() string {
	return "d3ef6eac"
}
func NewShowOneMapRequestMode() api.Message {
	return &ShowOneMapRequestMode{}
}

// ShowOneMapRequestModeReply represents the VPP binary API message 'show_one_map_request_mode_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 238:
//
//        ["show_one_map_request_mode_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "mode"],
//            {"crc" : "0xf3c29518"}
//        ],
//
type ShowOneMapRequestModeReply struct {
	Retval int32
	Mode   uint8
}

func (*ShowOneMapRequestModeReply) GetMessageName() string {
	return "show_one_map_request_mode_reply"
}
func (*ShowOneMapRequestModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowOneMapRequestModeReply) GetCrcString() string {
	return "f3c29518"
}
func NewShowOneMapRequestModeReply() api.Message {
	return &ShowOneMapRequestModeReply{}
}

// OneAddDelRemoteMapping represents the VPP binary API message 'one_add_del_remote_mapping'.
// Generated from '/usr/share/vpp/api//one.api.json', line 245:
//
//        ["one_add_del_remote_mapping",
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
//            ["vl_api_one_remote_locator_t", "rlocs", 0, "rloc_num"],
//            {"crc" : "0x02550000"}
//        ],
//
type OneAddDelRemoteMapping struct {
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
	Rlocs    []OneRemoteLocator
}

func (*OneAddDelRemoteMapping) GetMessageName() string {
	return "one_add_del_remote_mapping"
}
func (*OneAddDelRemoteMapping) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelRemoteMapping) GetCrcString() string {
	return "02550000"
}
func NewOneAddDelRemoteMapping() api.Message {
	return &OneAddDelRemoteMapping{}
}

// OneAddDelRemoteMappingReply represents the VPP binary API message 'one_add_del_remote_mapping_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 263:
//
//        ["one_add_del_remote_mapping_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xe23807e0"}
//        ],
//
type OneAddDelRemoteMappingReply struct {
	Retval int32
}

func (*OneAddDelRemoteMappingReply) GetMessageName() string {
	return "one_add_del_remote_mapping_reply"
}
func (*OneAddDelRemoteMappingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelRemoteMappingReply) GetCrcString() string {
	return "e23807e0"
}
func NewOneAddDelRemoteMappingReply() api.Message {
	return &OneAddDelRemoteMappingReply{}
}

// OneAddDelAdjacency represents the VPP binary API message 'one_add_del_adjacency'.
// Generated from '/usr/share/vpp/api//one.api.json', line 269:
//
//        ["one_add_del_adjacency",
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
//            {"crc" : "0xebba0b82"}
//        ],
//
type OneAddDelAdjacency struct {
	IsAdd   uint8
	Vni     uint32
	EidType uint8
	Reid    []byte `struc:"[16]byte"`
	Leid    []byte `struc:"[16]byte"`
	ReidLen uint8
	LeidLen uint8
}

func (*OneAddDelAdjacency) GetMessageName() string {
	return "one_add_del_adjacency"
}
func (*OneAddDelAdjacency) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelAdjacency) GetCrcString() string {
	return "ebba0b82"
}
func NewOneAddDelAdjacency() api.Message {
	return &OneAddDelAdjacency{}
}

// OneAddDelAdjacencyReply represents the VPP binary API message 'one_add_del_adjacency_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 282:
//
//        ["one_add_del_adjacency_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x224ad6d3"}
//        ],
//
type OneAddDelAdjacencyReply struct {
	Retval int32
}

func (*OneAddDelAdjacencyReply) GetMessageName() string {
	return "one_add_del_adjacency_reply"
}
func (*OneAddDelAdjacencyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelAdjacencyReply) GetCrcString() string {
	return "224ad6d3"
}
func NewOneAddDelAdjacencyReply() api.Message {
	return &OneAddDelAdjacencyReply{}
}

// OneAddDelMapRequestItrRlocs represents the VPP binary API message 'one_add_del_map_request_itr_rlocs'.
// Generated from '/usr/share/vpp/api//one.api.json', line 288:
//
//        ["one_add_del_map_request_itr_rlocs",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "locator_set_name", 64],
//            {"crc" : "0x11184601"}
//        ],
//
type OneAddDelMapRequestItrRlocs struct {
	IsAdd          uint8
	LocatorSetName []byte `struc:"[64]byte"`
}

func (*OneAddDelMapRequestItrRlocs) GetMessageName() string {
	return "one_add_del_map_request_itr_rlocs"
}
func (*OneAddDelMapRequestItrRlocs) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAddDelMapRequestItrRlocs) GetCrcString() string {
	return "11184601"
}
func NewOneAddDelMapRequestItrRlocs() api.Message {
	return &OneAddDelMapRequestItrRlocs{}
}

// OneAddDelMapRequestItrRlocsReply represents the VPP binary API message 'one_add_del_map_request_itr_rlocs_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 296:
//
//        ["one_add_del_map_request_itr_rlocs_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xba0b92e3"}
//        ],
//
type OneAddDelMapRequestItrRlocsReply struct {
	Retval int32
}

func (*OneAddDelMapRequestItrRlocsReply) GetMessageName() string {
	return "one_add_del_map_request_itr_rlocs_reply"
}
func (*OneAddDelMapRequestItrRlocsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAddDelMapRequestItrRlocsReply) GetCrcString() string {
	return "ba0b92e3"
}
func NewOneAddDelMapRequestItrRlocsReply() api.Message {
	return &OneAddDelMapRequestItrRlocsReply{}
}

// OneEidTableAddDelMap represents the VPP binary API message 'one_eid_table_add_del_map'.
// Generated from '/usr/share/vpp/api//one.api.json', line 302:
//
//        ["one_eid_table_add_del_map",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "vni"],
//            ["u32", "dp_table"],
//            ["u8", "is_l2"],
//            {"crc" : "0xe2ae3961"}
//        ],
//
type OneEidTableAddDelMap struct {
	IsAdd   uint8
	Vni     uint32
	DpTable uint32
	IsL2    uint8
}

func (*OneEidTableAddDelMap) GetMessageName() string {
	return "one_eid_table_add_del_map"
}
func (*OneEidTableAddDelMap) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneEidTableAddDelMap) GetCrcString() string {
	return "e2ae3961"
}
func NewOneEidTableAddDelMap() api.Message {
	return &OneEidTableAddDelMap{}
}

// OneEidTableAddDelMapReply represents the VPP binary API message 'one_eid_table_add_del_map_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 312:
//
//        ["one_eid_table_add_del_map_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x798662b9"}
//        ],
//
type OneEidTableAddDelMapReply struct {
	Retval int32
}

func (*OneEidTableAddDelMapReply) GetMessageName() string {
	return "one_eid_table_add_del_map_reply"
}
func (*OneEidTableAddDelMapReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneEidTableAddDelMapReply) GetCrcString() string {
	return "798662b9"
}
func NewOneEidTableAddDelMapReply() api.Message {
	return &OneEidTableAddDelMapReply{}
}

// OneLocatorDump represents the VPP binary API message 'one_locator_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 318:
//
//        ["one_locator_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "ls_index"],
//            ["u8", "ls_name", 64],
//            ["u8", "is_index_set"],
//            {"crc" : "0x9ba86482"}
//        ],
//
type OneLocatorDump struct {
	LsIndex    uint32
	LsName     []byte `struc:"[64]byte"`
	IsIndexSet uint8
}

func (*OneLocatorDump) GetMessageName() string {
	return "one_locator_dump"
}
func (*OneLocatorDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneLocatorDump) GetCrcString() string {
	return "9ba86482"
}
func NewOneLocatorDump() api.Message {
	return &OneLocatorDump{}
}

// OneLocatorDetails represents the VPP binary API message 'one_locator_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 327:
//
//        ["one_locator_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "local"],
//            ["u32", "sw_if_index"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            ["u8", "priority"],
//            ["u8", "weight"],
//            {"crc" : "0x46aff68e"}
//        ],
//
type OneLocatorDetails struct {
	Local     uint8
	SwIfIndex uint32
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
	Priority  uint8
	Weight    uint8
}

func (*OneLocatorDetails) GetMessageName() string {
	return "one_locator_details"
}
func (*OneLocatorDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneLocatorDetails) GetCrcString() string {
	return "46aff68e"
}
func NewOneLocatorDetails() api.Message {
	return &OneLocatorDetails{}
}

// OneLocatorSetDetails represents the VPP binary API message 'one_locator_set_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 338:
//
//        ["one_locator_set_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "ls_index"],
//            ["u8", "ls_name", 64],
//            {"crc" : "0x4d7a0c92"}
//        ],
//
type OneLocatorSetDetails struct {
	LsIndex uint32
	LsName  []byte `struc:"[64]byte"`
}

func (*OneLocatorSetDetails) GetMessageName() string {
	return "one_locator_set_details"
}
func (*OneLocatorSetDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneLocatorSetDetails) GetCrcString() string {
	return "4d7a0c92"
}
func NewOneLocatorSetDetails() api.Message {
	return &OneLocatorSetDetails{}
}

// OneLocatorSetDump represents the VPP binary API message 'one_locator_set_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 345:
//
//        ["one_locator_set_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "filter"],
//            {"crc" : "0x3dff8c66"}
//        ],
//
type OneLocatorSetDump struct {
	Filter uint8
}

func (*OneLocatorSetDump) GetMessageName() string {
	return "one_locator_set_dump"
}
func (*OneLocatorSetDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneLocatorSetDump) GetCrcString() string {
	return "3dff8c66"
}
func NewOneLocatorSetDump() api.Message {
	return &OneLocatorSetDump{}
}

// OneEidTableDetails represents the VPP binary API message 'one_eid_table_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 352:
//
//        ["one_eid_table_details",
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
//            {"crc" : "0xc059f5a4"}
//        ],
//
type OneEidTableDetails struct {
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

func (*OneEidTableDetails) GetMessageName() string {
	return "one_eid_table_details"
}
func (*OneEidTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneEidTableDetails) GetCrcString() string {
	return "c059f5a4"
}
func NewOneEidTableDetails() api.Message {
	return &OneEidTableDetails{}
}

// OneEidTableDump represents the VPP binary API message 'one_eid_table_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 371:
//
//        ["one_eid_table_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "eid_set"],
//            ["u8", "prefix_length"],
//            ["u32", "vni"],
//            ["u8", "eid_type"],
//            ["u8", "eid", 16],
//            ["u8", "filter"],
//            {"crc" : "0xab7f0498"}
//        ],
//
type OneEidTableDump struct {
	EidSet       uint8
	PrefixLength uint8
	Vni          uint32
	EidType      uint8
	Eid          []byte `struc:"[16]byte"`
	Filter       uint8
}

func (*OneEidTableDump) GetMessageName() string {
	return "one_eid_table_dump"
}
func (*OneEidTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneEidTableDump) GetCrcString() string {
	return "ab7f0498"
}
func NewOneEidTableDump() api.Message {
	return &OneEidTableDump{}
}

// OneAdjacenciesGetReply represents the VPP binary API message 'one_adjacencies_get_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 383:
//
//        ["one_adjacencies_get_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "count"],
//            ["vl_api_one_adjacency_t", "adjacencies", 0, "count"],
//            {"crc" : "0x73a0a09b"}
//        ],
//
type OneAdjacenciesGetReply struct {
	Retval      int32
	Count       uint32 `struc:"sizeof=Adjacencies"`
	Adjacencies []OneAdjacency
}

func (*OneAdjacenciesGetReply) GetMessageName() string {
	return "one_adjacencies_get_reply"
}
func (*OneAdjacenciesGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneAdjacenciesGetReply) GetCrcString() string {
	return "73a0a09b"
}
func NewOneAdjacenciesGetReply() api.Message {
	return &OneAdjacenciesGetReply{}
}

// OneAdjacenciesGet represents the VPP binary API message 'one_adjacencies_get'.
// Generated from '/usr/share/vpp/api//one.api.json', line 391:
//
//        ["one_adjacencies_get",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            {"crc" : "0xa884eff5"}
//        ],
//
type OneAdjacenciesGet struct {
	Vni uint32
}

func (*OneAdjacenciesGet) GetMessageName() string {
	return "one_adjacencies_get"
}
func (*OneAdjacenciesGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneAdjacenciesGet) GetCrcString() string {
	return "a884eff5"
}
func NewOneAdjacenciesGet() api.Message {
	return &OneAdjacenciesGet{}
}

// OneEidTableMapDetails represents the VPP binary API message 'one_eid_table_map_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 398:
//
//        ["one_eid_table_map_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            ["u32", "dp_table"],
//            {"crc" : "0xefdce03b"}
//        ],
//
type OneEidTableMapDetails struct {
	Vni     uint32
	DpTable uint32
}

func (*OneEidTableMapDetails) GetMessageName() string {
	return "one_eid_table_map_details"
}
func (*OneEidTableMapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneEidTableMapDetails) GetCrcString() string {
	return "efdce03b"
}
func NewOneEidTableMapDetails() api.Message {
	return &OneEidTableMapDetails{}
}

// OneEidTableMapDump represents the VPP binary API message 'one_eid_table_map_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 405:
//
//        ["one_eid_table_map_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_l2"],
//            {"crc" : "0x3f898ffc"}
//        ],
//
type OneEidTableMapDump struct {
	IsL2 uint8
}

func (*OneEidTableMapDump) GetMessageName() string {
	return "one_eid_table_map_dump"
}
func (*OneEidTableMapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneEidTableMapDump) GetCrcString() string {
	return "3f898ffc"
}
func NewOneEidTableMapDump() api.Message {
	return &OneEidTableMapDump{}
}

// OneEidTableVniDump represents the VPP binary API message 'one_eid_table_vni_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 412:
//
//        ["one_eid_table_vni_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xdc237d6b"}
//        ],
//
type OneEidTableVniDump struct {
}

func (*OneEidTableVniDump) GetMessageName() string {
	return "one_eid_table_vni_dump"
}
func (*OneEidTableVniDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneEidTableVniDump) GetCrcString() string {
	return "dc237d6b"
}
func NewOneEidTableVniDump() api.Message {
	return &OneEidTableVniDump{}
}

// OneEidTableVniDetails represents the VPP binary API message 'one_eid_table_vni_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 418:
//
//        ["one_eid_table_vni_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            {"crc" : "0x29d7a2ee"}
//        ],
//
type OneEidTableVniDetails struct {
	Vni uint32
}

func (*OneEidTableVniDetails) GetMessageName() string {
	return "one_eid_table_vni_details"
}
func (*OneEidTableVniDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneEidTableVniDetails) GetCrcString() string {
	return "29d7a2ee"
}
func NewOneEidTableVniDetails() api.Message {
	return &OneEidTableVniDetails{}
}

// OneMapResolverDetails represents the VPP binary API message 'one_map_resolver_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 425:
//
//        ["one_map_resolver_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0xb1223d87"}
//        ],
//
type OneMapResolverDetails struct {
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*OneMapResolverDetails) GetMessageName() string {
	return "one_map_resolver_details"
}
func (*OneMapResolverDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneMapResolverDetails) GetCrcString() string {
	return "b1223d87"
}
func NewOneMapResolverDetails() api.Message {
	return &OneMapResolverDetails{}
}

// OneMapResolverDump represents the VPP binary API message 'one_map_resolver_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 432:
//
//        ["one_map_resolver_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x2a3c1778"}
//        ],
//
type OneMapResolverDump struct {
}

func (*OneMapResolverDump) GetMessageName() string {
	return "one_map_resolver_dump"
}
func (*OneMapResolverDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneMapResolverDump) GetCrcString() string {
	return "2a3c1778"
}
func NewOneMapResolverDump() api.Message {
	return &OneMapResolverDump{}
}

// OneMapServerDetails represents the VPP binary API message 'one_map_server_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 438:
//
//        ["one_map_server_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            {"crc" : "0x493b5607"}
//        ],
//
type OneMapServerDetails struct {
	IsIpv6    uint8
	IPAddress []byte `struc:"[16]byte"`
}

func (*OneMapServerDetails) GetMessageName() string {
	return "one_map_server_details"
}
func (*OneMapServerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneMapServerDetails) GetCrcString() string {
	return "493b5607"
}
func NewOneMapServerDetails() api.Message {
	return &OneMapServerDetails{}
}

// OneMapServerDump represents the VPP binary API message 'one_map_server_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 445:
//
//        ["one_map_server_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xd05edf96"}
//        ],
//
type OneMapServerDump struct {
}

func (*OneMapServerDump) GetMessageName() string {
	return "one_map_server_dump"
}
func (*OneMapServerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneMapServerDump) GetCrcString() string {
	return "d05edf96"
}
func NewOneMapServerDump() api.Message {
	return &OneMapServerDump{}
}

// ShowOneStatus represents the VPP binary API message 'show_one_status'.
// Generated from '/usr/share/vpp/api//one.api.json', line 451:
//
//        ["show_one_status",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x34b8c69d"}
//        ],
//
type ShowOneStatus struct {
}

func (*ShowOneStatus) GetMessageName() string {
	return "show_one_status"
}
func (*ShowOneStatus) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowOneStatus) GetCrcString() string {
	return "34b8c69d"
}
func NewShowOneStatus() api.Message {
	return &ShowOneStatus{}
}

// ShowOneStatusReply represents the VPP binary API message 'show_one_status_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 457:
//
//        ["show_one_status_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "feature_status"],
//            ["u8", "gpe_status"],
//            {"crc" : "0xe0f96f92"}
//        ],
//
type ShowOneStatusReply struct {
	Retval        int32
	FeatureStatus uint8
	GpeStatus     uint8
}

func (*ShowOneStatusReply) GetMessageName() string {
	return "show_one_status_reply"
}
func (*ShowOneStatusReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowOneStatusReply) GetCrcString() string {
	return "e0f96f92"
}
func NewShowOneStatusReply() api.Message {
	return &ShowOneStatusReply{}
}

// OneGetMapRequestItrRlocs represents the VPP binary API message 'one_get_map_request_itr_rlocs'.
// Generated from '/usr/share/vpp/api//one.api.json', line 465:
//
//        ["one_get_map_request_itr_rlocs",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x34014276"}
//        ],
//
type OneGetMapRequestItrRlocs struct {
}

func (*OneGetMapRequestItrRlocs) GetMessageName() string {
	return "one_get_map_request_itr_rlocs"
}
func (*OneGetMapRequestItrRlocs) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneGetMapRequestItrRlocs) GetCrcString() string {
	return "34014276"
}
func NewOneGetMapRequestItrRlocs() api.Message {
	return &OneGetMapRequestItrRlocs{}
}

// OneGetMapRequestItrRlocsReply represents the VPP binary API message 'one_get_map_request_itr_rlocs_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 471:
//
//        ["one_get_map_request_itr_rlocs_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "locator_set_name", 64],
//            {"crc" : "0xd45b9cde"}
//        ],
//
type OneGetMapRequestItrRlocsReply struct {
	Retval         int32
	LocatorSetName []byte `struc:"[64]byte"`
}

func (*OneGetMapRequestItrRlocsReply) GetMessageName() string {
	return "one_get_map_request_itr_rlocs_reply"
}
func (*OneGetMapRequestItrRlocsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneGetMapRequestItrRlocsReply) GetCrcString() string {
	return "d45b9cde"
}
func NewOneGetMapRequestItrRlocsReply() api.Message {
	return &OneGetMapRequestItrRlocsReply{}
}

// ShowOnePitr represents the VPP binary API message 'show_one_pitr'.
// Generated from '/usr/share/vpp/api//one.api.json', line 478:
//
//        ["show_one_pitr",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x110fb2d5"}
//        ],
//
type ShowOnePitr struct {
}

func (*ShowOnePitr) GetMessageName() string {
	return "show_one_pitr"
}
func (*ShowOnePitr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowOnePitr) GetCrcString() string {
	return "110fb2d5"
}
func NewShowOnePitr() api.Message {
	return &ShowOnePitr{}
}

// ShowOnePitrReply represents the VPP binary API message 'show_one_pitr_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 484:
//
//        ["show_one_pitr_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "status"],
//            ["u8", "locator_set_name", 64],
//            {"crc" : "0xd60a3bf4"}
//        ],
//
type ShowOnePitrReply struct {
	Retval         int32
	Status         uint8
	LocatorSetName []byte `struc:"[64]byte"`
}

func (*ShowOnePitrReply) GetMessageName() string {
	return "show_one_pitr_reply"
}
func (*ShowOnePitrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowOnePitrReply) GetCrcString() string {
	return "d60a3bf4"
}
func NewShowOnePitrReply() api.Message {
	return &ShowOnePitrReply{}
}

// OneStatsDump represents the VPP binary API message 'one_stats_dump'.
// Generated from '/usr/share/vpp/api//one.api.json', line 492:
//
//        ["one_stats_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x8c9b17a1"}
//        ],
//
type OneStatsDump struct {
}

func (*OneStatsDump) GetMessageName() string {
	return "one_stats_dump"
}
func (*OneStatsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneStatsDump) GetCrcString() string {
	return "8c9b17a1"
}
func NewOneStatsDump() api.Message {
	return &OneStatsDump{}
}

// OneStatsDetails represents the VPP binary API message 'one_stats_details'.
// Generated from '/usr/share/vpp/api//one.api.json', line 498:
//
//        ["one_stats_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            ["u8", "eid_type"],
//            ["u8", "deid", 16],
//            ["u8", "seid", 16],
//            ["u8", "deid_pref_len"],
//            ["u8", "seid_pref_len"],
//            ["u8", "is_ip4"],
//            ["u8", "rloc", 16],
//            ["u8", "lloc", 16],
//            ["u32", "pkt_count"],
//            ["u32", "bytes"],
//            {"crc" : "0x759f7b5d"}
//        ],
//
type OneStatsDetails struct {
	Vni         uint32
	EidType     uint8
	Deid        []byte `struc:"[16]byte"`
	Seid        []byte `struc:"[16]byte"`
	DeidPrefLen uint8
	SeidPrefLen uint8
	IsIP4       uint8
	Rloc        []byte `struc:"[16]byte"`
	Lloc        []byte `struc:"[16]byte"`
	PktCount    uint32
	Bytes       uint32
}

func (*OneStatsDetails) GetMessageName() string {
	return "one_stats_details"
}
func (*OneStatsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneStatsDetails) GetCrcString() string {
	return "759f7b5d"
}
func NewOneStatsDetails() api.Message {
	return &OneStatsDetails{}
}

// OneStatsFlush represents the VPP binary API message 'one_stats_flush'.
// Generated from '/usr/share/vpp/api//one.api.json', line 514:
//
//        ["one_stats_flush",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x8472634b"}
//        ],
//
type OneStatsFlush struct {
}

func (*OneStatsFlush) GetMessageName() string {
	return "one_stats_flush"
}
func (*OneStatsFlush) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneStatsFlush) GetCrcString() string {
	return "8472634b"
}
func NewOneStatsFlush() api.Message {
	return &OneStatsFlush{}
}

// OneStatsFlushReply represents the VPP binary API message 'one_stats_flush_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 520:
//
//        ["one_stats_flush_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xc08bb290"}
//        ],
//
type OneStatsFlushReply struct {
	Retval int32
}

func (*OneStatsFlushReply) GetMessageName() string {
	return "one_stats_flush_reply"
}
func (*OneStatsFlushReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneStatsFlushReply) GetCrcString() string {
	return "c08bb290"
}
func NewOneStatsFlushReply() api.Message {
	return &OneStatsFlushReply{}
}

// OneStatsEnableDisable represents the VPP binary API message 'one_stats_enable_disable'.
// Generated from '/usr/share/vpp/api//one.api.json', line 526:
//
//        ["one_stats_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_en"],
//            {"crc" : "0x5586aa83"}
//        ],
//
type OneStatsEnableDisable struct {
	IsEn uint8
}

func (*OneStatsEnableDisable) GetMessageName() string {
	return "one_stats_enable_disable"
}
func (*OneStatsEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OneStatsEnableDisable) GetCrcString() string {
	return "5586aa83"
}
func NewOneStatsEnableDisable() api.Message {
	return &OneStatsEnableDisable{}
}

// OneStatsEnableDisableReply represents the VPP binary API message 'one_stats_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 533:
//
//        ["one_stats_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x1c5d305c"}
//        ],
//
type OneStatsEnableDisableReply struct {
	Retval int32
}

func (*OneStatsEnableDisableReply) GetMessageName() string {
	return "one_stats_enable_disable_reply"
}
func (*OneStatsEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OneStatsEnableDisableReply) GetCrcString() string {
	return "1c5d305c"
}
func NewOneStatsEnableDisableReply() api.Message {
	return &OneStatsEnableDisableReply{}
}

// ShowOneStatsEnableDisable represents the VPP binary API message 'show_one_stats_enable_disable'.
// Generated from '/usr/share/vpp/api//one.api.json', line 539:
//
//        ["show_one_stats_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xfd8b1e1f"}
//        ],
//
type ShowOneStatsEnableDisable struct {
}

func (*ShowOneStatsEnableDisable) GetMessageName() string {
	return "show_one_stats_enable_disable"
}
func (*ShowOneStatsEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowOneStatsEnableDisable) GetCrcString() string {
	return "fd8b1e1f"
}
func NewShowOneStatsEnableDisable() api.Message {
	return &ShowOneStatsEnableDisable{}
}

// ShowOneStatsEnableDisableReply represents the VPP binary API message 'show_one_stats_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//one.api.json', line 545:
//
//        ["show_one_stats_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "is_en"],
//            {"crc" : "0x7f197887"}
//        ]
//
type ShowOneStatsEnableDisableReply struct {
	Retval int32
	IsEn   uint8
}

func (*ShowOneStatsEnableDisableReply) GetMessageName() string {
	return "show_one_stats_enable_disable_reply"
}
func (*ShowOneStatsEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowOneStatsEnableDisableReply) GetCrcString() string {
	return "7f197887"
}
func NewShowOneStatsEnableDisableReply() api.Message {
	return &ShowOneStatsEnableDisableReply{}
}
