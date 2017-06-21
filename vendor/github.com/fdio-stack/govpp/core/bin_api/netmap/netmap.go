// Package netmap represents the VPP binary API of the 'netmap' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//netmap.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package netmap

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x54ce130c

// NetmapCreate represents the VPP binary API message 'netmap_create'.
// Generated from '/usr/share/vpp/api//netmap.api.json', line 6:
//
//        ["netmap_create",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "netmap_if_name", 64],
//            ["u8", "hw_addr", 6],
//            ["u8", "use_random_hw_addr"],
//            ["u8", "is_pipe"],
//            ["u8", "is_master"],
//            {"crc" : "0x0f13a603"}
//        ],
//
type NetmapCreate struct {
	NetmapIfName    []byte `struc:"[64]byte"`
	HwAddr          []byte `struc:"[6]byte"`
	UseRandomHwAddr uint8
	IsPipe          uint8
	IsMaster        uint8
}

func (*NetmapCreate) GetMessageName() string {
	return "netmap_create"
}
func (*NetmapCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NetmapCreate) GetCrcString() string {
	return "0f13a603"
}
func NewNetmapCreate() api.Message {
	return &NetmapCreate{}
}

// NetmapCreateReply represents the VPP binary API message 'netmap_create_reply'.
// Generated from '/usr/share/vpp/api//netmap.api.json', line 17:
//
//        ["netmap_create_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x70d29cfe"}
//        ],
//
type NetmapCreateReply struct {
	Retval int32
}

func (*NetmapCreateReply) GetMessageName() string {
	return "netmap_create_reply"
}
func (*NetmapCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NetmapCreateReply) GetCrcString() string {
	return "70d29cfe"
}
func NewNetmapCreateReply() api.Message {
	return &NetmapCreateReply{}
}

// NetmapDelete represents the VPP binary API message 'netmap_delete'.
// Generated from '/usr/share/vpp/api//netmap.api.json', line 23:
//
//        ["netmap_delete",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "netmap_if_name", 64],
//            {"crc" : "0x43e5c963"}
//        ],
//
type NetmapDelete struct {
	NetmapIfName []byte `struc:"[64]byte"`
}

func (*NetmapDelete) GetMessageName() string {
	return "netmap_delete"
}
func (*NetmapDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NetmapDelete) GetCrcString() string {
	return "43e5c963"
}
func NewNetmapDelete() api.Message {
	return &NetmapDelete{}
}

// NetmapDeleteReply represents the VPP binary API message 'netmap_delete_reply'.
// Generated from '/usr/share/vpp/api//netmap.api.json', line 30:
//
//        ["netmap_delete_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x81ecfa0d"}
//        ]
//
type NetmapDeleteReply struct {
	Retval int32
}

func (*NetmapDeleteReply) GetMessageName() string {
	return "netmap_delete_reply"
}
func (*NetmapDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NetmapDeleteReply) GetCrcString() string {
	return "81ecfa0d"
}
func NewNetmapDeleteReply() api.Message {
	return &NetmapDeleteReply{}
}
