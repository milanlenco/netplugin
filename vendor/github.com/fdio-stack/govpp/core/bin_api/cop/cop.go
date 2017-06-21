// Package cop represents the VPP binary API of the 'cop' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//cop.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package cop

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x67d58ee7

// CopInterfaceEnableDisable represents the VPP binary API message 'cop_interface_enable_disable'.
// Generated from '/usr/share/vpp/api//cop.api.json', line 6:
//
//        ["cop_interface_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "enable_disable"],
//            {"crc" : "0x1c65bd42"}
//        ],
//
type CopInterfaceEnableDisable struct {
	SwIfIndex     uint32
	EnableDisable uint8
}

func (*CopInterfaceEnableDisable) GetMessageName() string {
	return "cop_interface_enable_disable"
}
func (*CopInterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CopInterfaceEnableDisable) GetCrcString() string {
	return "1c65bd42"
}
func NewCopInterfaceEnableDisable() api.Message {
	return &CopInterfaceEnableDisable{}
}

// CopInterfaceEnableDisableReply represents the VPP binary API message 'cop_interface_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//cop.api.json', line 14:
//
//        ["cop_interface_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x123dd020"}
//        ],
//
type CopInterfaceEnableDisableReply struct {
	Retval int32
}

func (*CopInterfaceEnableDisableReply) GetMessageName() string {
	return "cop_interface_enable_disable_reply"
}
func (*CopInterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CopInterfaceEnableDisableReply) GetCrcString() string {
	return "123dd020"
}
func NewCopInterfaceEnableDisableReply() api.Message {
	return &CopInterfaceEnableDisableReply{}
}

// CopWhitelistEnableDisable represents the VPP binary API message 'cop_whitelist_enable_disable'.
// Generated from '/usr/share/vpp/api//cop.api.json', line 20:
//
//        ["cop_whitelist_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "fib_id"],
//            ["u8", "ip4"],
//            ["u8", "ip6"],
//            ["u8", "default_cop"],
//            {"crc" : "0x9a0ec2ec"}
//        ],
//
type CopWhitelistEnableDisable struct {
	SwIfIndex  uint32
	FibID      uint32
	IP4        uint8
	IP6        uint8
	DefaultCop uint8
}

func (*CopWhitelistEnableDisable) GetMessageName() string {
	return "cop_whitelist_enable_disable"
}
func (*CopWhitelistEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CopWhitelistEnableDisable) GetCrcString() string {
	return "9a0ec2ec"
}
func NewCopWhitelistEnableDisable() api.Message {
	return &CopWhitelistEnableDisable{}
}

// CopWhitelistEnableDisableReply represents the VPP binary API message 'cop_whitelist_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//cop.api.json', line 31:
//
//        ["cop_whitelist_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x3f660dee"}
//        ]
//
type CopWhitelistEnableDisableReply struct {
	Retval int32
}

func (*CopWhitelistEnableDisableReply) GetMessageName() string {
	return "cop_whitelist_enable_disable_reply"
}
func (*CopWhitelistEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CopWhitelistEnableDisableReply) GetCrcString() string {
	return "3f660dee"
}
func NewCopWhitelistEnableDisableReply() api.Message {
	return &CopWhitelistEnableDisableReply{}
}
