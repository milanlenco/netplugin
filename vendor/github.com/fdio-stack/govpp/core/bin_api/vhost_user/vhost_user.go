// Package vhost_user represents the VPP binary API of the 'vhost_user' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//vhost_user.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package vhost_user

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xb524ab11

// CreateVhostUserIf represents the VPP binary API message 'create_vhost_user_if'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 6:
//
//        ["create_vhost_user_if",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_server"],
//            ["u8", "sock_filename", 256],
//            ["u8", "renumber"],
//            ["u32", "custom_dev_instance"],
//            ["u8", "use_custom_mac"],
//            ["u8", "mac_address", 6],
//            ["u8", "tag", 64],
//            ["u8", "operation_mode"],
//            {"crc" : "0x2ef4dbed"}
//        ],
//
type CreateVhostUserIf struct {
	IsServer          uint8
	SockFilename      []byte `struc:"[256]byte"`
	Renumber          uint8
	CustomDevInstance uint32
	UseCustomMac      uint8
	MacAddress        []byte `struc:"[6]byte"`
	Tag               []byte `struc:"[64]byte"`
	OperationMode     uint8
}

func (*CreateVhostUserIf) GetMessageName() string {
	return "create_vhost_user_if"
}
func (*CreateVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CreateVhostUserIf) GetCrcString() string {
	return "2ef4dbed"
}
func NewCreateVhostUserIf() api.Message {
	return &CreateVhostUserIf{}
}

// CreateVhostUserIfReply represents the VPP binary API message 'create_vhost_user_if_reply'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 20:
//
//        ["create_vhost_user_if_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x20c3fea9"}
//        ],
//
type CreateVhostUserIfReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*CreateVhostUserIfReply) GetMessageName() string {
	return "create_vhost_user_if_reply"
}
func (*CreateVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CreateVhostUserIfReply) GetCrcString() string {
	return "20c3fea9"
}
func NewCreateVhostUserIfReply() api.Message {
	return &CreateVhostUserIfReply{}
}

// ModifyVhostUserIf represents the VPP binary API message 'modify_vhost_user_if'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 27:
//
//        ["modify_vhost_user_if",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "is_server"],
//            ["u8", "sock_filename", 256],
//            ["u8", "renumber"],
//            ["u32", "custom_dev_instance"],
//            ["u8", "operation_mode"],
//            {"crc" : "0x352e01c6"}
//        ],
//
type ModifyVhostUserIf struct {
	SwIfIndex         uint32
	IsServer          uint8
	SockFilename      []byte `struc:"[256]byte"`
	Renumber          uint8
	CustomDevInstance uint32
	OperationMode     uint8
}

func (*ModifyVhostUserIf) GetMessageName() string {
	return "modify_vhost_user_if"
}
func (*ModifyVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ModifyVhostUserIf) GetCrcString() string {
	return "352e01c6"
}
func NewModifyVhostUserIf() api.Message {
	return &ModifyVhostUserIf{}
}

// ModifyVhostUserIfReply represents the VPP binary API message 'modify_vhost_user_if_reply'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 39:
//
//        ["modify_vhost_user_if_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xdc039556"}
//        ],
//
type ModifyVhostUserIfReply struct {
	Retval int32
}

func (*ModifyVhostUserIfReply) GetMessageName() string {
	return "modify_vhost_user_if_reply"
}
func (*ModifyVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ModifyVhostUserIfReply) GetCrcString() string {
	return "dc039556"
}
func NewModifyVhostUserIfReply() api.Message {
	return &ModifyVhostUserIfReply{}
}

// DeleteVhostUserIf represents the VPP binary API message 'delete_vhost_user_if'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 45:
//
//        ["delete_vhost_user_if",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xed230259"}
//        ],
//
type DeleteVhostUserIf struct {
	SwIfIndex uint32
}

func (*DeleteVhostUserIf) GetMessageName() string {
	return "delete_vhost_user_if"
}
func (*DeleteVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DeleteVhostUserIf) GetCrcString() string {
	return "ed230259"
}
func NewDeleteVhostUserIf() api.Message {
	return &DeleteVhostUserIf{}
}

// DeleteVhostUserIfReply represents the VPP binary API message 'delete_vhost_user_if_reply'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 52:
//
//        ["delete_vhost_user_if_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x2d7b6fe6"}
//        ],
//
type DeleteVhostUserIfReply struct {
	Retval int32
}

func (*DeleteVhostUserIfReply) GetMessageName() string {
	return "delete_vhost_user_if_reply"
}
func (*DeleteVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DeleteVhostUserIfReply) GetCrcString() string {
	return "2d7b6fe6"
}
func NewDeleteVhostUserIfReply() api.Message {
	return &DeleteVhostUserIfReply{}
}

// SwInterfaceVhostUserDetails represents the VPP binary API message 'sw_interface_vhost_user_details'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 58:
//
//        ["sw_interface_vhost_user_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "interface_name", 64],
//            ["u32", "virtio_net_hdr_sz"],
//            ["u64", "features"],
//            ["u8", "is_server"],
//            ["u8", "sock_filename", 256],
//            ["u32", "num_regions"],
//            ["i32", "sock_errno"],
//            ["u8", "operation_mode"],
//            {"crc" : "0x00d6399c"}
//        ],
//
type SwInterfaceVhostUserDetails struct {
	SwIfIndex      uint32
	InterfaceName  []byte `struc:"[64]byte"`
	VirtioNetHdrSz uint32
	Features       uint64
	IsServer       uint8
	SockFilename   []byte `struc:"[256]byte"`
	NumRegions     uint32
	SockErrno      int32
	OperationMode  uint8
}

func (*SwInterfaceVhostUserDetails) GetMessageName() string {
	return "sw_interface_vhost_user_details"
}
func (*SwInterfaceVhostUserDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceVhostUserDetails) GetCrcString() string {
	return "00d6399c"
}
func NewSwInterfaceVhostUserDetails() api.Message {
	return &SwInterfaceVhostUserDetails{}
}

// SwInterfaceVhostUserDump represents the VPP binary API message 'sw_interface_vhost_user_dump'.
// Generated from '/usr/share/vpp/api//vhost_user.api.json', line 72:
//
//        ["sw_interface_vhost_user_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xda1923e2"}
//        ]
//
type SwInterfaceVhostUserDump struct {
}

func (*SwInterfaceVhostUserDump) GetMessageName() string {
	return "sw_interface_vhost_user_dump"
}
func (*SwInterfaceVhostUserDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceVhostUserDump) GetCrcString() string {
	return "da1923e2"
}
func NewSwInterfaceVhostUserDump() api.Message {
	return &SwInterfaceVhostUserDump{}
}
