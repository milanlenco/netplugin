// Package ioam_vxlan_gpe represents the VPP binary API of the 'ioam_vxlan_gpe' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package ioam_vxlan_gpe

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xfe93c8d5

// VxlanGpeIoamEnable represents the VPP binary API message 'vxlan_gpe_ioam_enable'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 6:
//
//        ["vxlan_gpe_ioam_enable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u16", "id"],
//            ["u8", "trace_ppc"],
//            ["u8", "pow_enable"],
//            ["u8", "trace_enable"],
//            {"crc" : "0x6bf84bd6"}
//        ],
//
type VxlanGpeIoamEnable struct {
	ID          uint16
	TracePpc    uint8
	PowEnable   uint8
	TraceEnable uint8
}

func (*VxlanGpeIoamEnable) GetMessageName() string {
	return "vxlan_gpe_ioam_enable"
}
func (*VxlanGpeIoamEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeIoamEnable) GetCrcString() string {
	return "6bf84bd6"
}
func NewVxlanGpeIoamEnable() api.Message {
	return &VxlanGpeIoamEnable{}
}

// VxlanGpeIoamEnableReply represents the VPP binary API message 'vxlan_gpe_ioam_enable_reply'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 16:
//
//        ["vxlan_gpe_ioam_enable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xca79fc00"}
//        ],
//
type VxlanGpeIoamEnableReply struct {
	Retval int32
}

func (*VxlanGpeIoamEnableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_enable_reply"
}
func (*VxlanGpeIoamEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeIoamEnableReply) GetCrcString() string {
	return "ca79fc00"
}
func NewVxlanGpeIoamEnableReply() api.Message {
	return &VxlanGpeIoamEnableReply{}
}

// VxlanGpeIoamDisable represents the VPP binary API message 'vxlan_gpe_ioam_disable'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 22:
//
//        ["vxlan_gpe_ioam_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u16", "id"],
//            {"crc" : "0x1a373a3b"}
//        ],
//
type VxlanGpeIoamDisable struct {
	ID uint16
}

func (*VxlanGpeIoamDisable) GetMessageName() string {
	return "vxlan_gpe_ioam_disable"
}
func (*VxlanGpeIoamDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeIoamDisable) GetCrcString() string {
	return "1a373a3b"
}
func NewVxlanGpeIoamDisable() api.Message {
	return &VxlanGpeIoamDisable{}
}

// VxlanGpeIoamDisableReply represents the VPP binary API message 'vxlan_gpe_ioam_disable_reply'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 29:
//
//        ["vxlan_gpe_ioam_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x711375e4"}
//        ],
//
type VxlanGpeIoamDisableReply struct {
	Retval int32
}

func (*VxlanGpeIoamDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_disable_reply"
}
func (*VxlanGpeIoamDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeIoamDisableReply) GetCrcString() string {
	return "711375e4"
}
func NewVxlanGpeIoamDisableReply() api.Message {
	return &VxlanGpeIoamDisableReply{}
}

// VxlanGpeIoamVniEnable represents the VPP binary API message 'vxlan_gpe_ioam_vni_enable'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 35:
//
//        ["vxlan_gpe_ioam_vni_enable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            ["u8", "local", 16],
//            ["u8", "remote", 16],
//            ["u8", "is_ipv6"],
//            {"crc" : "0x489195ec"}
//        ],
//
type VxlanGpeIoamVniEnable struct {
	Vni    uint32
	Local  []byte `struc:"[16]byte"`
	Remote []byte `struc:"[16]byte"`
	IsIpv6 uint8
}

func (*VxlanGpeIoamVniEnable) GetMessageName() string {
	return "vxlan_gpe_ioam_vni_enable"
}
func (*VxlanGpeIoamVniEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeIoamVniEnable) GetCrcString() string {
	return "489195ec"
}
func NewVxlanGpeIoamVniEnable() api.Message {
	return &VxlanGpeIoamVniEnable{}
}

// VxlanGpeIoamVniEnableReply represents the VPP binary API message 'vxlan_gpe_ioam_vni_enable_reply'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 45:
//
//        ["vxlan_gpe_ioam_vni_enable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb17ff9c6"}
//        ],
//
type VxlanGpeIoamVniEnableReply struct {
	Retval int32
}

func (*VxlanGpeIoamVniEnableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_vni_enable_reply"
}
func (*VxlanGpeIoamVniEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeIoamVniEnableReply) GetCrcString() string {
	return "b17ff9c6"
}
func NewVxlanGpeIoamVniEnableReply() api.Message {
	return &VxlanGpeIoamVniEnableReply{}
}

// VxlanGpeIoamVniDisable represents the VPP binary API message 'vxlan_gpe_ioam_vni_disable'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 51:
//
//        ["vxlan_gpe_ioam_vni_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vni"],
//            ["u8", "local", 16],
//            ["u8", "remote", 16],
//            ["u8", "is_ipv6"],
//            {"crc" : "0x27392af3"}
//        ],
//
type VxlanGpeIoamVniDisable struct {
	Vni    uint32
	Local  []byte `struc:"[16]byte"`
	Remote []byte `struc:"[16]byte"`
	IsIpv6 uint8
}

func (*VxlanGpeIoamVniDisable) GetMessageName() string {
	return "vxlan_gpe_ioam_vni_disable"
}
func (*VxlanGpeIoamVniDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeIoamVniDisable) GetCrcString() string {
	return "27392af3"
}
func NewVxlanGpeIoamVniDisable() api.Message {
	return &VxlanGpeIoamVniDisable{}
}

// VxlanGpeIoamVniDisableReply represents the VPP binary API message 'vxlan_gpe_ioam_vni_disable_reply'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 61:
//
//        ["vxlan_gpe_ioam_vni_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x9444b2cd"}
//        ],
//
type VxlanGpeIoamVniDisableReply struct {
	Retval int32
}

func (*VxlanGpeIoamVniDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_vni_disable_reply"
}
func (*VxlanGpeIoamVniDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeIoamVniDisableReply) GetCrcString() string {
	return "9444b2cd"
}
func NewVxlanGpeIoamVniDisableReply() api.Message {
	return &VxlanGpeIoamVniDisableReply{}
}

// VxlanGpeIoamTransitEnable represents the VPP binary API message 'vxlan_gpe_ioam_transit_enable'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 67:
//
//        ["vxlan_gpe_ioam_transit_enable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "outer_fib_index"],
//            ["u8", "dst_addr", 16],
//            ["u8", "is_ipv6"],
//            {"crc" : "0x2c399a17"}
//        ],
//
type VxlanGpeIoamTransitEnable struct {
	OuterFibIndex uint32
	DstAddr       []byte `struc:"[16]byte"`
	IsIpv6        uint8
}

func (*VxlanGpeIoamTransitEnable) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_enable"
}
func (*VxlanGpeIoamTransitEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeIoamTransitEnable) GetCrcString() string {
	return "2c399a17"
}
func NewVxlanGpeIoamTransitEnable() api.Message {
	return &VxlanGpeIoamTransitEnable{}
}

// VxlanGpeIoamTransitEnableReply represents the VPP binary API message 'vxlan_gpe_ioam_transit_enable_reply'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 76:
//
//        ["vxlan_gpe_ioam_transit_enable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x4d80f122"}
//        ],
//
type VxlanGpeIoamTransitEnableReply struct {
	Retval int32
}

func (*VxlanGpeIoamTransitEnableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_enable_reply"
}
func (*VxlanGpeIoamTransitEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeIoamTransitEnableReply) GetCrcString() string {
	return "4d80f122"
}
func NewVxlanGpeIoamTransitEnableReply() api.Message {
	return &VxlanGpeIoamTransitEnableReply{}
}

// VxlanGpeIoamTransitDisable represents the VPP binary API message 'vxlan_gpe_ioam_transit_disable'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 82:
//
//        ["vxlan_gpe_ioam_transit_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "outer_fib_index"],
//            ["u8", "dst_addr", 16],
//            ["u8", "is_ipv6"],
//            {"crc" : "0xee3cf5f9"}
//        ],
//
type VxlanGpeIoamTransitDisable struct {
	OuterFibIndex uint32
	DstAddr       []byte `struc:"[16]byte"`
	IsIpv6        uint8
}

func (*VxlanGpeIoamTransitDisable) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_disable"
}
func (*VxlanGpeIoamTransitDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanGpeIoamTransitDisable) GetCrcString() string {
	return "ee3cf5f9"
}
func NewVxlanGpeIoamTransitDisable() api.Message {
	return &VxlanGpeIoamTransitDisable{}
}

// VxlanGpeIoamTransitDisableReply represents the VPP binary API message 'vxlan_gpe_ioam_transit_disable_reply'.
// Generated from '/usr/share/vpp/api//ioam_vxlan_gpe.api.json', line 91:
//
//        ["vxlan_gpe_ioam_transit_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb01272c0"}
//        ]
//
type VxlanGpeIoamTransitDisableReply struct {
	Retval int32
}

func (*VxlanGpeIoamTransitDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_disable_reply"
}
func (*VxlanGpeIoamTransitDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanGpeIoamTransitDisableReply) GetCrcString() string {
	return "b01272c0"
}
func NewVxlanGpeIoamTransitDisableReply() api.Message {
	return &VxlanGpeIoamTransitDisableReply{}
}
