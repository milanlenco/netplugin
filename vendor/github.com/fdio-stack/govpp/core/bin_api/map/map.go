// Package map represents the VPP binary API of the 'map' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//map.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package map
import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion =  0x2d1ae300


// MapAddDomain represents the VPP binary API message 'map_add_domain'.
// Generated from '/usr/share/vpp/api//map.api.json', line 6:
//
//        ["map_add_domain",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "ip6_prefix", 16],
//            ["u8", "ip4_prefix", 4],
//            ["u8", "ip6_src", 16],
//            ["u8", "ip6_prefix_len"],
//            ["u8", "ip4_prefix_len"],
//            ["u8", "ip6_src_prefix_len"],
//            ["u8", "ea_bits_len"],
//            ["u8", "psid_offset"],
//            ["u8", "psid_length"],
//            ["u8", "is_translation"],
//            ["u16", "mtu"],
//            {"crc" : "0xe83fafb7"}
//        ],
//
type MapAddDomain struct {
	IP6Prefix []byte	`struc:"[16]byte"`
	IP4Prefix []byte	`struc:"[4]byte"`
	IP6Src []byte	`struc:"[16]byte"`
	IP6PrefixLen uint8
	IP4PrefixLen uint8
	IP6SrcPrefixLen uint8
	EaBitsLen uint8
	PsidOffset uint8
	PsidLength uint8
	IsTranslation uint8
	Mtu uint16
}
func (*MapAddDomain) GetMessageName() string {
	return "map_add_domain"
}
func (*MapAddDomain) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MapAddDomain) GetCrcString() string {
	return "e83fafb7"
}
func NewMapAddDomain() api.Message {
	return &MapAddDomain{}
}

// MapAddDomainReply represents the VPP binary API message 'map_add_domain_reply'.
// Generated from '/usr/share/vpp/api//map.api.json', line 23:
//
//        ["map_add_domain_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "index"],
//            ["i32", "retval"],
//            {"crc" : "0xf6c61861"}
//        ],
//
type MapAddDomainReply struct {
	Index uint32
	Retval int32
}
func (*MapAddDomainReply) GetMessageName() string {
	return "map_add_domain_reply"
}
func (*MapAddDomainReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MapAddDomainReply) GetCrcString() string {
	return "f6c61861"
}
func NewMapAddDomainReply() api.Message {
	return &MapAddDomainReply{}
}

// MapDelDomain represents the VPP binary API message 'map_del_domain'.
// Generated from '/usr/share/vpp/api//map.api.json', line 30:
//
//        ["map_del_domain",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "index"],
//            {"crc" : "0x4c06875c"}
//        ],
//
type MapDelDomain struct {
	Index uint32
}
func (*MapDelDomain) GetMessageName() string {
	return "map_del_domain"
}
func (*MapDelDomain) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MapDelDomain) GetCrcString() string {
	return "4c06875c"
}
func NewMapDelDomain() api.Message {
	return &MapDelDomain{}
}

// MapDelDomainReply represents the VPP binary API message 'map_del_domain_reply'.
// Generated from '/usr/share/vpp/api//map.api.json', line 37:
//
//        ["map_del_domain_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x7c61a925"}
//        ],
//
type MapDelDomainReply struct {
	Retval int32
}
func (*MapDelDomainReply) GetMessageName() string {
	return "map_del_domain_reply"
}
func (*MapDelDomainReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MapDelDomainReply) GetCrcString() string {
	return "7c61a925"
}
func NewMapDelDomainReply() api.Message {
	return &MapDelDomainReply{}
}

// MapAddDelRule represents the VPP binary API message 'map_add_del_rule'.
// Generated from '/usr/share/vpp/api//map.api.json', line 43:
//
//        ["map_add_del_rule",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "index"],
//            ["u32", "is_add"],
//            ["u8", "ip6_dst", 16],
//            ["u16", "psid"],
//            {"crc" : "0xcc9054fe"}
//        ],
//
type MapAddDelRule struct {
	Index uint32
	IsAdd uint32
	IP6Dst []byte	`struc:"[16]byte"`
	Psid uint16
}
func (*MapAddDelRule) GetMessageName() string {
	return "map_add_del_rule"
}
func (*MapAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MapAddDelRule) GetCrcString() string {
	return "cc9054fe"
}
func NewMapAddDelRule() api.Message {
	return &MapAddDelRule{}
}

// MapAddDelRuleReply represents the VPP binary API message 'map_add_del_rule_reply'.
// Generated from '/usr/share/vpp/api//map.api.json', line 53:
//
//        ["map_add_del_rule_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x02c9e765"}
//        ],
//
type MapAddDelRuleReply struct {
	Retval int32
}
func (*MapAddDelRuleReply) GetMessageName() string {
	return "map_add_del_rule_reply"
}
func (*MapAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MapAddDelRuleReply) GetCrcString() string {
	return "02c9e765"
}
func NewMapAddDelRuleReply() api.Message {
	return &MapAddDelRuleReply{}
}

// MapDomainDump represents the VPP binary API message 'map_domain_dump'.
// Generated from '/usr/share/vpp/api//map.api.json', line 59:
//
//        ["map_domain_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xb64dff5d"}
//        ],
//
type MapDomainDump struct {
}
func (*MapDomainDump) GetMessageName() string {
	return "map_domain_dump"
}
func (*MapDomainDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MapDomainDump) GetCrcString() string {
	return "b64dff5d"
}
func NewMapDomainDump() api.Message {
	return &MapDomainDump{}
}

// MapDomainDetails represents the VPP binary API message 'map_domain_details'.
// Generated from '/usr/share/vpp/api//map.api.json', line 65:
//
//        ["map_domain_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "domain_index"],
//            ["u8", "ip6_prefix", 16],
//            ["u8", "ip4_prefix", 4],
//            ["u8", "ip6_src", 16],
//            ["u8", "ip6_prefix_len"],
//            ["u8", "ip4_prefix_len"],
//            ["u8", "ip6_src_len"],
//            ["u8", "ea_bits_len"],
//            ["u8", "psid_offset"],
//            ["u8", "psid_length"],
//            ["u8", "flags"],
//            ["u16", "mtu"],
//            ["u8", "is_translation"],
//            {"crc" : "0xcaa121eb"}
//        ],
//
type MapDomainDetails struct {
	DomainIndex uint32
	IP6Prefix []byte	`struc:"[16]byte"`
	IP4Prefix []byte	`struc:"[4]byte"`
	IP6Src []byte	`struc:"[16]byte"`
	IP6PrefixLen uint8
	IP4PrefixLen uint8
	IP6SrcLen uint8
	EaBitsLen uint8
	PsidOffset uint8
	PsidLength uint8
	Flags uint8
	Mtu uint16
	IsTranslation uint8
}
func (*MapDomainDetails) GetMessageName() string {
	return "map_domain_details"
}
func (*MapDomainDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MapDomainDetails) GetCrcString() string {
	return "caa121eb"
}
func NewMapDomainDetails() api.Message {
	return &MapDomainDetails{}
}

// MapRuleDump represents the VPP binary API message 'map_rule_dump'.
// Generated from '/usr/share/vpp/api//map.api.json', line 83:
//
//        ["map_rule_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "domain_index"],
//            {"crc" : "0x11d0e36b"}
//        ],
//
type MapRuleDump struct {
	DomainIndex uint32
}
func (*MapRuleDump) GetMessageName() string {
	return "map_rule_dump"
}
func (*MapRuleDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MapRuleDump) GetCrcString() string {
	return "11d0e36b"
}
func NewMapRuleDump() api.Message {
	return &MapRuleDump{}
}

// MapRuleDetails represents the VPP binary API message 'map_rule_details'.
// Generated from '/usr/share/vpp/api//map.api.json', line 90:
//
//        ["map_rule_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "ip6_dst", 16],
//            ["u16", "psid"],
//            {"crc" : "0x7ef2b612"}
//        ],
//
type MapRuleDetails struct {
	IP6Dst []byte	`struc:"[16]byte"`
	Psid uint16
}
func (*MapRuleDetails) GetMessageName() string {
	return "map_rule_details"
}
func (*MapRuleDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MapRuleDetails) GetCrcString() string {
	return "7ef2b612"
}
func NewMapRuleDetails() api.Message {
	return &MapRuleDetails{}
}

// MapSummaryStats represents the VPP binary API message 'map_summary_stats'.
// Generated from '/usr/share/vpp/api//map.api.json', line 97:
//
//        ["map_summary_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xf6a12980"}
//        ],
//
type MapSummaryStats struct {
}
func (*MapSummaryStats) GetMessageName() string {
	return "map_summary_stats"
}
func (*MapSummaryStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MapSummaryStats) GetCrcString() string {
	return "f6a12980"
}
func NewMapSummaryStats() api.Message {
	return &MapSummaryStats{}
}

// MapSummaryStatsReply represents the VPP binary API message 'map_summary_stats_reply'.
// Generated from '/usr/share/vpp/api//map.api.json', line 103:
//
//        ["map_summary_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u64", "total_bindings"],
//            ["u64", "total_pkts", 2],
//            ["u64", "total_bytes", 2],
//            ["u64", "total_ip4_fragments"],
//            ["u64", "total_security_check", 2],
//            {"crc" : "0x700b3e21"}
//        ]
//
type MapSummaryStatsReply struct {
	Retval int32
	TotalBindings uint64
	TotalPkts []uint64	`struc:"[2]uint64"`
	TotalBytes []uint64	`struc:"[2]uint64"`
	TotalIP4Fragments uint64
	TotalSecurityCheck []uint64	`struc:"[2]uint64"`
}
func (*MapSummaryStatsReply) GetMessageName() string {
	return "map_summary_stats_reply"
}
func (*MapSummaryStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MapSummaryStatsReply) GetCrcString() string {
	return "700b3e21"
}
func NewMapSummaryStatsReply() api.Message {
	return &MapSummaryStatsReply{}
}
