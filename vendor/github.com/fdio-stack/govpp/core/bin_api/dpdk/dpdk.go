// Package dpdk represents the VPP binary API of the 'dpdk' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//dpdk.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package dpdk

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xecd34367

// SwInterfaceSetDpdkHqosPipe represents the VPP binary API message 'sw_interface_set_dpdk_hqos_pipe'.
// Generated from '/usr/share/vpp/api//dpdk.api.json', line 6:
//
//        ["sw_interface_set_dpdk_hqos_pipe",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "subport"],
//            ["u32", "pipe"],
//            ["u32", "profile"],
//            {"crc" : "0xbe9b2181"}
//        ],
//
type SwInterfaceSetDpdkHqosPipe struct {
	SwIfIndex uint32
	Subport   uint32
	Pipe      uint32
	Profile   uint32
}

func (*SwInterfaceSetDpdkHqosPipe) GetMessageName() string {
	return "sw_interface_set_dpdk_hqos_pipe"
}
func (*SwInterfaceSetDpdkHqosPipe) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetDpdkHqosPipe) GetCrcString() string {
	return "be9b2181"
}
func NewSwInterfaceSetDpdkHqosPipe() api.Message {
	return &SwInterfaceSetDpdkHqosPipe{}
}

// SwInterfaceSetDpdkHqosPipeReply represents the VPP binary API message 'sw_interface_set_dpdk_hqos_pipe_reply'.
// Generated from '/usr/share/vpp/api//dpdk.api.json', line 16:
//
//        ["sw_interface_set_dpdk_hqos_pipe_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x1222fa49"}
//        ],
//
type SwInterfaceSetDpdkHqosPipeReply struct {
	Retval int32
}

func (*SwInterfaceSetDpdkHqosPipeReply) GetMessageName() string {
	return "sw_interface_set_dpdk_hqos_pipe_reply"
}
func (*SwInterfaceSetDpdkHqosPipeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetDpdkHqosPipeReply) GetCrcString() string {
	return "1222fa49"
}
func NewSwInterfaceSetDpdkHqosPipeReply() api.Message {
	return &SwInterfaceSetDpdkHqosPipeReply{}
}

// SwInterfaceSetDpdkHqosSubport represents the VPP binary API message 'sw_interface_set_dpdk_hqos_subport'.
// Generated from '/usr/share/vpp/api//dpdk.api.json', line 22:
//
//        ["sw_interface_set_dpdk_hqos_subport",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "subport"],
//            ["u32", "tb_rate"],
//            ["u32", "tb_size"],
//            ["u32", "tc_rate", 4],
//            ["u32", "tc_period"],
//            {"crc" : "0x05a9fa2c"}
//        ],
//
type SwInterfaceSetDpdkHqosSubport struct {
	SwIfIndex uint32
	Subport   uint32
	TbRate    uint32
	TbSize    uint32
	TcRate    []uint32 `struc:"[4]uint32"`
	TcPeriod  uint32
}

func (*SwInterfaceSetDpdkHqosSubport) GetMessageName() string {
	return "sw_interface_set_dpdk_hqos_subport"
}
func (*SwInterfaceSetDpdkHqosSubport) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetDpdkHqosSubport) GetCrcString() string {
	return "05a9fa2c"
}
func NewSwInterfaceSetDpdkHqosSubport() api.Message {
	return &SwInterfaceSetDpdkHqosSubport{}
}

// SwInterfaceSetDpdkHqosSubportReply represents the VPP binary API message 'sw_interface_set_dpdk_hqos_subport_reply'.
// Generated from '/usr/share/vpp/api//dpdk.api.json', line 34:
//
//        ["sw_interface_set_dpdk_hqos_subport_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x147f258a"}
//        ],
//
type SwInterfaceSetDpdkHqosSubportReply struct {
	Retval int32
}

func (*SwInterfaceSetDpdkHqosSubportReply) GetMessageName() string {
	return "sw_interface_set_dpdk_hqos_subport_reply"
}
func (*SwInterfaceSetDpdkHqosSubportReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetDpdkHqosSubportReply) GetCrcString() string {
	return "147f258a"
}
func NewSwInterfaceSetDpdkHqosSubportReply() api.Message {
	return &SwInterfaceSetDpdkHqosSubportReply{}
}

// SwInterfaceSetDpdkHqosTctbl represents the VPP binary API message 'sw_interface_set_dpdk_hqos_tctbl'.
// Generated from '/usr/share/vpp/api//dpdk.api.json', line 40:
//
//        ["sw_interface_set_dpdk_hqos_tctbl",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "entry"],
//            ["u32", "tc"],
//            ["u32", "queue"],
//            {"crc" : "0xf9d98f13"}
//        ],
//
type SwInterfaceSetDpdkHqosTctbl struct {
	SwIfIndex uint32
	Entry     uint32
	Tc        uint32
	Queue     uint32
}

func (*SwInterfaceSetDpdkHqosTctbl) GetMessageName() string {
	return "sw_interface_set_dpdk_hqos_tctbl"
}
func (*SwInterfaceSetDpdkHqosTctbl) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetDpdkHqosTctbl) GetCrcString() string {
	return "f9d98f13"
}
func NewSwInterfaceSetDpdkHqosTctbl() api.Message {
	return &SwInterfaceSetDpdkHqosTctbl{}
}

// SwInterfaceSetDpdkHqosTctblReply represents the VPP binary API message 'sw_interface_set_dpdk_hqos_tctbl_reply'.
// Generated from '/usr/share/vpp/api//dpdk.api.json', line 50:
//
//        ["sw_interface_set_dpdk_hqos_tctbl_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x4e2f524d"}
//        ]
//
type SwInterfaceSetDpdkHqosTctblReply struct {
	Retval int32
}

func (*SwInterfaceSetDpdkHqosTctblReply) GetMessageName() string {
	return "sw_interface_set_dpdk_hqos_tctbl_reply"
}
func (*SwInterfaceSetDpdkHqosTctblReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetDpdkHqosTctblReply) GetCrcString() string {
	return "4e2f524d"
}
func NewSwInterfaceSetDpdkHqosTctblReply() api.Message {
	return &SwInterfaceSetDpdkHqosTctblReply{}
}
