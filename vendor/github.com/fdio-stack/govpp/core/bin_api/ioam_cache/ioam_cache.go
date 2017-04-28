// Package ioam_cache represents the VPP binary API of the 'ioam_cache' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//ioam_cache.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package ioam_cache

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x6e24611e

// IoamCacheIP6EnableDisable represents the VPP binary API message 'ioam_cache_ip6_enable_disable'.
// Generated from '/usr/share/vpp/api//ioam_cache.api.json', line 6:
//
//        ["ioam_cache_ip6_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_disable"],
//            {"crc" : "0xde631cb6"}
//        ],
//
type IoamCacheIP6EnableDisable struct {
	IsDisable uint8
}

func (*IoamCacheIP6EnableDisable) GetMessageName() string {
	return "ioam_cache_ip6_enable_disable"
}
func (*IoamCacheIP6EnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IoamCacheIP6EnableDisable) GetCrcString() string {
	return "de631cb6"
}
func NewIoamCacheIP6EnableDisable() api.Message {
	return &IoamCacheIP6EnableDisable{}
}

// IoamCacheIP6EnableDisableReply represents the VPP binary API message 'ioam_cache_ip6_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//ioam_cache.api.json', line 13:
//
//        ["ioam_cache_ip6_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x67f2a36b"}
//        ]
//
type IoamCacheIP6EnableDisableReply struct {
	Retval int32
}

func (*IoamCacheIP6EnableDisableReply) GetMessageName() string {
	return "ioam_cache_ip6_enable_disable_reply"
}
func (*IoamCacheIP6EnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IoamCacheIP6EnableDisableReply) GetCrcString() string {
	return "67f2a36b"
}
func NewIoamCacheIP6EnableDisableReply() api.Message {
	return &IoamCacheIP6EnableDisableReply{}
}
