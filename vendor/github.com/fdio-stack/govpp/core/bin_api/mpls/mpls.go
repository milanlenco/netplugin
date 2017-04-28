// Package mpls represents the VPP binary API of the 'mpls' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//mpls.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package mpls

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xdb72b0b0

// FibPath2 represents the VPP binary API data type 'fib_path2'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 3:
//
//        ["fib_path2",
//            ["u32", "sw_if_index"],
//            ["u32", "weight"],
//            ["u8", "is_local"],
//            ["u8", "is_drop"],
//            ["u8", "is_unreach"],
//            ["u8", "is_prohibit"],
//            ["u8", "afi"],
//            ["u8", "next_hop", 16],
//            ["u32", "labels", 16],
//            {"crc" : "0x0aa98130"}
//        ]
//
type FibPath2 struct {
	SwIfIndex  uint32
	Weight     uint32
	IsLocal    uint8
	IsDrop     uint8
	IsUnreach  uint8
	IsProhibit uint8
	Afi        uint8
	NextHop    []byte   `struc:"[16]byte"`
	Labels     []uint32 `struc:"[16]uint32"`
}

func (*FibPath2) GetTypeName() string {
	return "fib_path2"
}
func (*FibPath2) GetCrcString() string {
	return "0aa98130"
}

// MplsIPBindUnbind represents the VPP binary API message 'mpls_ip_bind_unbind'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 17:
//
//        ["mpls_ip_bind_unbind",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "mb_mpls_table_id"],
//            ["u32", "mb_label"],
//            ["u32", "mb_ip_table_id"],
//            ["u8", "mb_create_table_if_needed"],
//            ["u8", "mb_is_bind"],
//            ["u8", "mb_is_ip4"],
//            ["u8", "mb_address_length"],
//            ["u8", "mb_address", 16],
//            {"crc" : "0x167f24bd"}
//        ],
//
type MplsIPBindUnbind struct {
	MbMplsTableID         uint32
	MbLabel               uint32
	MbIPTableID           uint32
	MbCreateTableIfNeeded uint8
	MbIsBind              uint8
	MbIsIP4               uint8
	MbAddressLength       uint8
	MbAddress             []byte `struc:"[16]byte"`
}

func (*MplsIPBindUnbind) GetMessageName() string {
	return "mpls_ip_bind_unbind"
}
func (*MplsIPBindUnbind) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MplsIPBindUnbind) GetCrcString() string {
	return "167f24bd"
}
func NewMplsIPBindUnbind() api.Message {
	return &MplsIPBindUnbind{}
}

// MplsIPBindUnbindReply represents the VPP binary API message 'mpls_ip_bind_unbind_reply'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 31:
//
//        ["mpls_ip_bind_unbind_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5753d1ed"}
//        ],
//
type MplsIPBindUnbindReply struct {
	Retval int32
}

func (*MplsIPBindUnbindReply) GetMessageName() string {
	return "mpls_ip_bind_unbind_reply"
}
func (*MplsIPBindUnbindReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MplsIPBindUnbindReply) GetCrcString() string {
	return "5753d1ed"
}
func NewMplsIPBindUnbindReply() api.Message {
	return &MplsIPBindUnbindReply{}
}

// MplsTunnelAddDel represents the VPP binary API message 'mpls_tunnel_add_del'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 37:
//
//        ["mpls_tunnel_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "mt_sw_if_index"],
//            ["u8", "mt_is_add"],
//            ["u8", "mt_l2_only"],
//            ["u8", "mt_is_multicast"],
//            ["u8", "mt_next_hop_proto_is_ip4"],
//            ["u8", "mt_next_hop_weight"],
//            ["u8", "mt_next_hop", 16],
//            ["u8", "mt_next_hop_n_out_labels"],
//            ["u32", "mt_next_hop_sw_if_index"],
//            ["u32", "mt_next_hop_table_id"],
//            ["u32", "mt_next_hop_out_label_stack", 0, "mt_next_hop_n_out_labels"],
//            {"crc" : "0x09b7931c"}
//        ],
//
type MplsTunnelAddDel struct {
	MtSwIfIndex            uint32
	MtIsAdd                uint8
	MtL2Only               uint8
	MtIsMulticast          uint8
	MtNextHopProtoIsIP4    uint8
	MtNextHopWeight        uint8
	MtNextHop              []byte `struc:"[16]byte"`
	MtNextHopNOutLabels    uint8  `struc:"sizeof=MtNextHopOutLabelStack"`
	MtNextHopSwIfIndex     uint32
	MtNextHopTableID       uint32
	MtNextHopOutLabelStack []uint32
}

func (*MplsTunnelAddDel) GetMessageName() string {
	return "mpls_tunnel_add_del"
}
func (*MplsTunnelAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MplsTunnelAddDel) GetCrcString() string {
	return "09b7931c"
}
func NewMplsTunnelAddDel() api.Message {
	return &MplsTunnelAddDel{}
}

// MplsTunnelAddDelReply represents the VPP binary API message 'mpls_tunnel_add_del_reply'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 54:
//
//        ["mpls_tunnel_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xbb483273"}
//        ],
//
type MplsTunnelAddDelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*MplsTunnelAddDelReply) GetMessageName() string {
	return "mpls_tunnel_add_del_reply"
}
func (*MplsTunnelAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MplsTunnelAddDelReply) GetCrcString() string {
	return "bb483273"
}
func NewMplsTunnelAddDelReply() api.Message {
	return &MplsTunnelAddDelReply{}
}

// MplsTunnelDump represents the VPP binary API message 'mpls_tunnel_dump'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 61:
//
//        ["mpls_tunnel_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["i32", "tunnel_index"],
//            {"crc" : "0xbe9ada9c"}
//        ],
//
type MplsTunnelDump struct {
	TunnelIndex int32
}

func (*MplsTunnelDump) GetMessageName() string {
	return "mpls_tunnel_dump"
}
func (*MplsTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MplsTunnelDump) GetCrcString() string {
	return "be9ada9c"
}
func NewMplsTunnelDump() api.Message {
	return &MplsTunnelDump{}
}

// MplsTunnelDetails represents the VPP binary API message 'mpls_tunnel_details'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 68:
//
//        ["mpls_tunnel_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "mt_sw_if_index"],
//            ["u8", "mt_tunnel_index"],
//            ["u8", "mt_l2_only"],
//            ["u8", "mt_is_multicast"],
//            ["u32", "mt_count"],
//            ["vl_api_fib_path2_t", "mt_paths", 0, "mt_count"],
//            {"crc" : "0xf43de2ca"}
//        ],
//
type MplsTunnelDetails struct {
	MtSwIfIndex   uint8
	MtTunnelIndex uint8
	MtL2Only      uint8
	MtIsMulticast uint8
	MtCount       uint32 `struc:"sizeof=MtPaths"`
	MtPaths       []FibPath2
}

func (*MplsTunnelDetails) GetMessageName() string {
	return "mpls_tunnel_details"
}
func (*MplsTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MplsTunnelDetails) GetCrcString() string {
	return "f43de2ca"
}
func NewMplsTunnelDetails() api.Message {
	return &MplsTunnelDetails{}
}

// MplsRouteAddDel represents the VPP binary API message 'mpls_route_add_del'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 79:
//
//        ["mpls_route_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "mr_label"],
//            ["u8", "mr_eos"],
//            ["u32", "mr_table_id"],
//            ["u32", "mr_classify_table_index"],
//            ["u8", "mr_create_table_if_needed"],
//            ["u8", "mr_is_add"],
//            ["u8", "mr_is_classify"],
//            ["u8", "mr_is_multicast"],
//            ["u8", "mr_is_multipath"],
//            ["u8", "mr_is_resolve_host"],
//            ["u8", "mr_is_resolve_attached"],
//            ["u8", "mr_is_interface_rx"],
//            ["u8", "mr_is_rpf_id"],
//            ["u8", "mr_next_hop_proto_is_ip4"],
//            ["u8", "mr_next_hop_weight"],
//            ["u8", "mr_next_hop", 16],
//            ["u8", "mr_next_hop_n_out_labels"],
//            ["u32", "mr_next_hop_sw_if_index"],
//            ["u32", "mr_next_hop_table_id"],
//            ["u32", "mr_next_hop_via_label"],
//            ["u32", "mr_next_hop_out_label_stack", 0, "mr_next_hop_n_out_labels"],
//            {"crc" : "0x70edc775"}
//        ],
//
type MplsRouteAddDel struct {
	MrLabel                uint32
	MrEos                  uint8
	MrTableID              uint32
	MrClassifyTableIndex   uint32
	MrCreateTableIfNeeded  uint8
	MrIsAdd                uint8
	MrIsClassify           uint8
	MrIsMulticast          uint8
	MrIsMultipath          uint8
	MrIsResolveHost        uint8
	MrIsResolveAttached    uint8
	MrIsInterfaceRx        uint8
	MrIsRpfID              uint8
	MrNextHopProtoIsIP4    uint8
	MrNextHopWeight        uint8
	MrNextHop              []byte `struc:"[16]byte"`
	MrNextHopNOutLabels    uint8  `struc:"sizeof=MrNextHopOutLabelStack"`
	MrNextHopSwIfIndex     uint32
	MrNextHopTableID       uint32
	MrNextHopViaLabel      uint32
	MrNextHopOutLabelStack []uint32
}

func (*MplsRouteAddDel) GetMessageName() string {
	return "mpls_route_add_del"
}
func (*MplsRouteAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MplsRouteAddDel) GetCrcString() string {
	return "70edc775"
}
func NewMplsRouteAddDel() api.Message {
	return &MplsRouteAddDel{}
}

// MplsRouteAddDelReply represents the VPP binary API message 'mpls_route_add_del_reply'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 106:
//
//        ["mpls_route_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x21a12fe9"}
//        ],
//
type MplsRouteAddDelReply struct {
	Retval int32
}

func (*MplsRouteAddDelReply) GetMessageName() string {
	return "mpls_route_add_del_reply"
}
func (*MplsRouteAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MplsRouteAddDelReply) GetCrcString() string {
	return "21a12fe9"
}
func NewMplsRouteAddDelReply() api.Message {
	return &MplsRouteAddDelReply{}
}

// MplsFibDump represents the VPP binary API message 'mpls_fib_dump'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 112:
//
//        ["mpls_fib_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x7e82659e"}
//        ],
//
type MplsFibDump struct {
}

func (*MplsFibDump) GetMessageName() string {
	return "mpls_fib_dump"
}
func (*MplsFibDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MplsFibDump) GetCrcString() string {
	return "7e82659e"
}
func NewMplsFibDump() api.Message {
	return &MplsFibDump{}
}

// MplsFibDetails represents the VPP binary API message 'mpls_fib_details'.
// Generated from '/usr/share/vpp/api//mpls.api.json', line 118:
//
//        ["mpls_fib_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "table_id"],
//            ["u8", "eos_bit"],
//            ["u32", "label"],
//            ["u32", "count"],
//            ["vl_api_fib_path2_t", "path", 0, "count"],
//            {"crc" : "0x2804be60"}
//        ]
//
type MplsFibDetails struct {
	TableID uint32
	EosBit  uint8
	Label   uint32
	Count   uint32 `struc:"sizeof=Path"`
	Path    []FibPath2
}

func (*MplsFibDetails) GetMessageName() string {
	return "mpls_fib_details"
}
func (*MplsFibDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MplsFibDetails) GetCrcString() string {
	return "2804be60"
}
func NewMplsFibDetails() api.Message {
	return &MplsFibDetails{}
}
