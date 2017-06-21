// Package ipsec represents the VPP binary API of the 'ipsec' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//ipsec.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package ipsec

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xccd0c837

// IpsecSpdAddDel represents the VPP binary API message 'ipsec_spd_add_del'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 6:
//
//        ["ipsec_spd_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "spd_id"],
//            {"crc" : "0x9b42314b"}
//        ],
//
type IpsecSpdAddDel struct {
	IsAdd uint8
	SpdID uint32
}

func (*IpsecSpdAddDel) GetMessageName() string {
	return "ipsec_spd_add_del"
}
func (*IpsecSpdAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecSpdAddDel) GetCrcString() string {
	return "9b42314b"
}
func NewIpsecSpdAddDel() api.Message {
	return &IpsecSpdAddDel{}
}

// IpsecSpdAddDelReply represents the VPP binary API message 'ipsec_spd_add_del_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 14:
//
//        ["ipsec_spd_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xc7439119"}
//        ],
//
type IpsecSpdAddDelReply struct {
	Retval int32
}

func (*IpsecSpdAddDelReply) GetMessageName() string {
	return "ipsec_spd_add_del_reply"
}
func (*IpsecSpdAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecSpdAddDelReply) GetCrcString() string {
	return "c7439119"
}
func NewIpsecSpdAddDelReply() api.Message {
	return &IpsecSpdAddDelReply{}
}

// IpsecInterfaceAddDelSpd represents the VPP binary API message 'ipsec_interface_add_del_spd'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 20:
//
//        ["ipsec_interface_add_del_spd",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "sw_if_index"],
//            ["u32", "spd_id"],
//            {"crc" : "0x52de89dc"}
//        ],
//
type IpsecInterfaceAddDelSpd struct {
	IsAdd     uint8
	SwIfIndex uint32
	SpdID     uint32
}

func (*IpsecInterfaceAddDelSpd) GetMessageName() string {
	return "ipsec_interface_add_del_spd"
}
func (*IpsecInterfaceAddDelSpd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecInterfaceAddDelSpd) GetCrcString() string {
	return "52de89dc"
}
func NewIpsecInterfaceAddDelSpd() api.Message {
	return &IpsecInterfaceAddDelSpd{}
}

// IpsecInterfaceAddDelSpdReply represents the VPP binary API message 'ipsec_interface_add_del_spd_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 29:
//
//        ["ipsec_interface_add_del_spd_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x977b7be9"}
//        ],
//
type IpsecInterfaceAddDelSpdReply struct {
	Retval int32
}

func (*IpsecInterfaceAddDelSpdReply) GetMessageName() string {
	return "ipsec_interface_add_del_spd_reply"
}
func (*IpsecInterfaceAddDelSpdReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecInterfaceAddDelSpdReply) GetCrcString() string {
	return "977b7be9"
}
func NewIpsecInterfaceAddDelSpdReply() api.Message {
	return &IpsecInterfaceAddDelSpdReply{}
}

// IpsecSpdAddDelEntry represents the VPP binary API message 'ipsec_spd_add_del_entry'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 35:
//
//        ["ipsec_spd_add_del_entry",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "spd_id"],
//            ["i32", "priority"],
//            ["u8", "is_outbound"],
//            ["u8", "is_ipv6"],
//            ["u8", "is_ip_any"],
//            ["u8", "remote_address_start", 16],
//            ["u8", "remote_address_stop", 16],
//            ["u8", "local_address_start", 16],
//            ["u8", "local_address_stop", 16],
//            ["u8", "protocol"],
//            ["u16", "remote_port_start"],
//            ["u16", "remote_port_stop"],
//            ["u16", "local_port_start"],
//            ["u16", "local_port_stop"],
//            ["u8", "policy"],
//            ["u32", "sa_id"],
//            {"crc" : "0xd85e0ed5"}
//        ],
//
type IpsecSpdAddDelEntry struct {
	IsAdd              uint8
	SpdID              uint32
	Priority           int32
	IsOutbound         uint8
	IsIpv6             uint8
	IsIPAny            uint8
	RemoteAddressStart []byte `struc:"[16]byte"`
	RemoteAddressStop  []byte `struc:"[16]byte"`
	LocalAddressStart  []byte `struc:"[16]byte"`
	LocalAddressStop   []byte `struc:"[16]byte"`
	Protocol           uint8
	RemotePortStart    uint16
	RemotePortStop     uint16
	LocalPortStart     uint16
	LocalPortStop      uint16
	Policy             uint8
	SaID               uint32
}

func (*IpsecSpdAddDelEntry) GetMessageName() string {
	return "ipsec_spd_add_del_entry"
}
func (*IpsecSpdAddDelEntry) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecSpdAddDelEntry) GetCrcString() string {
	return "d85e0ed5"
}
func NewIpsecSpdAddDelEntry() api.Message {
	return &IpsecSpdAddDelEntry{}
}

// IpsecSpdAddDelEntryReply represents the VPP binary API message 'ipsec_spd_add_del_entry_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 58:
//
//        ["ipsec_spd_add_del_entry_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xf96c8b2d"}
//        ],
//
type IpsecSpdAddDelEntryReply struct {
	Retval int32
}

func (*IpsecSpdAddDelEntryReply) GetMessageName() string {
	return "ipsec_spd_add_del_entry_reply"
}
func (*IpsecSpdAddDelEntryReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecSpdAddDelEntryReply) GetCrcString() string {
	return "f96c8b2d"
}
func NewIpsecSpdAddDelEntryReply() api.Message {
	return &IpsecSpdAddDelEntryReply{}
}

// IpsecSadAddDelEntry represents the VPP binary API message 'ipsec_sad_add_del_entry'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 64:
//
//        ["ipsec_sad_add_del_entry",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "sad_id"],
//            ["u32", "spi"],
//            ["u8", "protocol"],
//            ["u8", "crypto_algorithm"],
//            ["u8", "crypto_key_length"],
//            ["u8", "crypto_key", 128],
//            ["u8", "integrity_algorithm"],
//            ["u8", "integrity_key_length"],
//            ["u8", "integrity_key", 128],
//            ["u8", "use_extended_sequence_number"],
//            ["u8", "is_tunnel"],
//            ["u8", "is_tunnel_ipv6"],
//            ["u8", "tunnel_src_address", 16],
//            ["u8", "tunnel_dst_address", 16],
//            {"crc" : "0x7d6709e1"}
//        ],
//
type IpsecSadAddDelEntry struct {
	IsAdd                     uint8
	SadID                     uint32
	Spi                       uint32
	Protocol                  uint8
	CryptoAlgorithm           uint8
	CryptoKeyLength           uint8
	CryptoKey                 []byte `struc:"[128]byte"`
	IntegrityAlgorithm        uint8
	IntegrityKeyLength        uint8
	IntegrityKey              []byte `struc:"[128]byte"`
	UseExtendedSequenceNumber uint8
	IsTunnel                  uint8
	IsTunnelIpv6              uint8
	TunnelSrcAddress          []byte `struc:"[16]byte"`
	TunnelDstAddress          []byte `struc:"[16]byte"`
}

func (*IpsecSadAddDelEntry) GetMessageName() string {
	return "ipsec_sad_add_del_entry"
}
func (*IpsecSadAddDelEntry) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecSadAddDelEntry) GetCrcString() string {
	return "7d6709e1"
}
func NewIpsecSadAddDelEntry() api.Message {
	return &IpsecSadAddDelEntry{}
}

// IpsecSadAddDelEntryReply represents the VPP binary API message 'ipsec_sad_add_del_entry_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 85:
//
//        ["ipsec_sad_add_del_entry_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5cf382d8"}
//        ],
//
type IpsecSadAddDelEntryReply struct {
	Retval int32
}

func (*IpsecSadAddDelEntryReply) GetMessageName() string {
	return "ipsec_sad_add_del_entry_reply"
}
func (*IpsecSadAddDelEntryReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecSadAddDelEntryReply) GetCrcString() string {
	return "5cf382d8"
}
func NewIpsecSadAddDelEntryReply() api.Message {
	return &IpsecSadAddDelEntryReply{}
}

// IpsecSaSetKey represents the VPP binary API message 'ipsec_sa_set_key'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 91:
//
//        ["ipsec_sa_set_key",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sa_id"],
//            ["u8", "crypto_key_length"],
//            ["u8", "crypto_key", 128],
//            ["u8", "integrity_key_length"],
//            ["u8", "integrity_key", 128],
//            {"crc" : "0x99a67c60"}
//        ],
//
type IpsecSaSetKey struct {
	SaID               uint32
	CryptoKeyLength    uint8
	CryptoKey          []byte `struc:"[128]byte"`
	IntegrityKeyLength uint8
	IntegrityKey       []byte `struc:"[128]byte"`
}

func (*IpsecSaSetKey) GetMessageName() string {
	return "ipsec_sa_set_key"
}
func (*IpsecSaSetKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecSaSetKey) GetCrcString() string {
	return "99a67c60"
}
func NewIpsecSaSetKey() api.Message {
	return &IpsecSaSetKey{}
}

// IpsecSaSetKeyReply represents the VPP binary API message 'ipsec_sa_set_key_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 102:
//
//        ["ipsec_sa_set_key_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5c5b7b46"}
//        ],
//
type IpsecSaSetKeyReply struct {
	Retval int32
}

func (*IpsecSaSetKeyReply) GetMessageName() string {
	return "ipsec_sa_set_key_reply"
}
func (*IpsecSaSetKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecSaSetKeyReply) GetCrcString() string {
	return "5c5b7b46"
}
func NewIpsecSaSetKeyReply() api.Message {
	return &IpsecSaSetKeyReply{}
}

// Ikev2ProfileAddDel represents the VPP binary API message 'ikev2_profile_add_del'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 108:
//
//        ["ikev2_profile_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u8", "is_add"],
//            {"crc" : "0x37b6925c"}
//        ],
//
type Ikev2ProfileAddDel struct {
	Name  []byte `struc:"[64]byte"`
	IsAdd uint8
}

func (*Ikev2ProfileAddDel) GetMessageName() string {
	return "ikev2_profile_add_del"
}
func (*Ikev2ProfileAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2ProfileAddDel) GetCrcString() string {
	return "37b6925c"
}
func NewIkev2ProfileAddDel() api.Message {
	return &Ikev2ProfileAddDel{}
}

// Ikev2ProfileAddDelReply represents the VPP binary API message 'ikev2_profile_add_del_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 116:
//
//        ["ikev2_profile_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x7621f627"}
//        ],
//
type Ikev2ProfileAddDelReply struct {
	Retval int32
}

func (*Ikev2ProfileAddDelReply) GetMessageName() string {
	return "ikev2_profile_add_del_reply"
}
func (*Ikev2ProfileAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2ProfileAddDelReply) GetCrcString() string {
	return "7621f627"
}
func NewIkev2ProfileAddDelReply() api.Message {
	return &Ikev2ProfileAddDelReply{}
}

// Ikev2ProfileSetAuth represents the VPP binary API message 'ikev2_profile_set_auth'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 122:
//
//        ["ikev2_profile_set_auth",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u8", "auth_method"],
//            ["u8", "is_hex"],
//            ["u32", "data_len"],
//            ["u8", "data", 0],
//            {"crc" : "0xa0747739"}
//        ],
//
type Ikev2ProfileSetAuth struct {
	Name       []byte `struc:"[64]byte"`
	AuthMethod uint8
	IsHex      uint8
	DataLen    uint32
	Data       []byte
}

func (*Ikev2ProfileSetAuth) GetMessageName() string {
	return "ikev2_profile_set_auth"
}
func (*Ikev2ProfileSetAuth) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2ProfileSetAuth) GetCrcString() string {
	return "a0747739"
}
func NewIkev2ProfileSetAuth() api.Message {
	return &Ikev2ProfileSetAuth{}
}

// Ikev2ProfileSetAuthReply represents the VPP binary API message 'ikev2_profile_set_auth_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 133:
//
//        ["ikev2_profile_set_auth_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x46083d00"}
//        ],
//
type Ikev2ProfileSetAuthReply struct {
	Retval int32
}

func (*Ikev2ProfileSetAuthReply) GetMessageName() string {
	return "ikev2_profile_set_auth_reply"
}
func (*Ikev2ProfileSetAuthReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2ProfileSetAuthReply) GetCrcString() string {
	return "46083d00"
}
func NewIkev2ProfileSetAuthReply() api.Message {
	return &Ikev2ProfileSetAuthReply{}
}

// Ikev2ProfileSetID represents the VPP binary API message 'ikev2_profile_set_id'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 139:
//
//        ["ikev2_profile_set_id",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u8", "is_local"],
//            ["u8", "id_type"],
//            ["u32", "data_len"],
//            ["u8", "data", 0],
//            {"crc" : "0x0c2331dc"}
//        ],
//
type Ikev2ProfileSetID struct {
	Name    []byte `struc:"[64]byte"`
	IsLocal uint8
	IDType  uint8
	DataLen uint32
	Data    []byte
}

func (*Ikev2ProfileSetID) GetMessageName() string {
	return "ikev2_profile_set_id"
}
func (*Ikev2ProfileSetID) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2ProfileSetID) GetCrcString() string {
	return "0c2331dc"
}
func NewIkev2ProfileSetID() api.Message {
	return &Ikev2ProfileSetID{}
}

// Ikev2ProfileSetIDReply represents the VPP binary API message 'ikev2_profile_set_id_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 150:
//
//        ["ikev2_profile_set_id_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x66803be5"}
//        ],
//
type Ikev2ProfileSetIDReply struct {
	Retval int32
}

func (*Ikev2ProfileSetIDReply) GetMessageName() string {
	return "ikev2_profile_set_id_reply"
}
func (*Ikev2ProfileSetIDReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2ProfileSetIDReply) GetCrcString() string {
	return "66803be5"
}
func NewIkev2ProfileSetIDReply() api.Message {
	return &Ikev2ProfileSetIDReply{}
}

// Ikev2ProfileSetTs represents the VPP binary API message 'ikev2_profile_set_ts'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 156:
//
//        ["ikev2_profile_set_ts",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u8", "is_local"],
//            ["u8", "proto"],
//            ["u16", "start_port"],
//            ["u16", "end_port"],
//            ["u32", "start_addr"],
//            ["u32", "end_addr"],
//            {"crc" : "0x69587e0e"}
//        ],
//
type Ikev2ProfileSetTs struct {
	Name      []byte `struc:"[64]byte"`
	IsLocal   uint8
	Proto     uint8
	StartPort uint16
	EndPort   uint16
	StartAddr uint32
	EndAddr   uint32
}

func (*Ikev2ProfileSetTs) GetMessageName() string {
	return "ikev2_profile_set_ts"
}
func (*Ikev2ProfileSetTs) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2ProfileSetTs) GetCrcString() string {
	return "69587e0e"
}
func NewIkev2ProfileSetTs() api.Message {
	return &Ikev2ProfileSetTs{}
}

// Ikev2ProfileSetTsReply represents the VPP binary API message 'ikev2_profile_set_ts_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 169:
//
//        ["ikev2_profile_set_ts_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xe1c33583"}
//        ],
//
type Ikev2ProfileSetTsReply struct {
	Retval int32
}

func (*Ikev2ProfileSetTsReply) GetMessageName() string {
	return "ikev2_profile_set_ts_reply"
}
func (*Ikev2ProfileSetTsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2ProfileSetTsReply) GetCrcString() string {
	return "e1c33583"
}
func NewIkev2ProfileSetTsReply() api.Message {
	return &Ikev2ProfileSetTsReply{}
}

// Ikev2SetLocalKey represents the VPP binary API message 'ikev2_set_local_key'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 175:
//
//        ["ikev2_set_local_key",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "key_file", 256],
//            {"crc" : "0xa99b238a"}
//        ],
//
type Ikev2SetLocalKey struct {
	KeyFile []byte `struc:"[256]byte"`
}

func (*Ikev2SetLocalKey) GetMessageName() string {
	return "ikev2_set_local_key"
}
func (*Ikev2SetLocalKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2SetLocalKey) GetCrcString() string {
	return "a99b238a"
}
func NewIkev2SetLocalKey() api.Message {
	return &Ikev2SetLocalKey{}
}

// Ikev2SetLocalKeyReply represents the VPP binary API message 'ikev2_set_local_key_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 182:
//
//        ["ikev2_set_local_key_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x8f7a80e0"}
//        ],
//
type Ikev2SetLocalKeyReply struct {
	Retval int32
}

func (*Ikev2SetLocalKeyReply) GetMessageName() string {
	return "ikev2_set_local_key_reply"
}
func (*Ikev2SetLocalKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2SetLocalKeyReply) GetCrcString() string {
	return "8f7a80e0"
}
func NewIkev2SetLocalKeyReply() api.Message {
	return &Ikev2SetLocalKeyReply{}
}

// Ikev2SetResponder represents the VPP binary API message 'ikev2_set_responder'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 188:
//
//        ["ikev2_set_responder",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u32", "sw_if_index"],
//            ["u8", "address", 4],
//            {"crc" : "0x519141b4"}
//        ],
//
type Ikev2SetResponder struct {
	Name      []byte `struc:"[64]byte"`
	SwIfIndex uint32
	Address   []byte `struc:"[4]byte"`
}

func (*Ikev2SetResponder) GetMessageName() string {
	return "ikev2_set_responder"
}
func (*Ikev2SetResponder) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2SetResponder) GetCrcString() string {
	return "519141b4"
}
func NewIkev2SetResponder() api.Message {
	return &Ikev2SetResponder{}
}

// Ikev2SetResponderReply represents the VPP binary API message 'ikev2_set_responder_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 197:
//
//        ["ikev2_set_responder_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x2f1a94d7"}
//        ],
//
type Ikev2SetResponderReply struct {
	Retval int32
}

func (*Ikev2SetResponderReply) GetMessageName() string {
	return "ikev2_set_responder_reply"
}
func (*Ikev2SetResponderReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2SetResponderReply) GetCrcString() string {
	return "2f1a94d7"
}
func NewIkev2SetResponderReply() api.Message {
	return &Ikev2SetResponderReply{}
}

// Ikev2SetIkeTransforms represents the VPP binary API message 'ikev2_set_ike_transforms'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 203:
//
//        ["ikev2_set_ike_transforms",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u32", "crypto_alg"],
//            ["u32", "crypto_key_size"],
//            ["u32", "integ_alg"],
//            ["u32", "dh_group"],
//            {"crc" : "0xf0bf018e"}
//        ],
//
type Ikev2SetIkeTransforms struct {
	Name          []byte `struc:"[64]byte"`
	CryptoAlg     uint32
	CryptoKeySize uint32
	IntegAlg      uint32
	DhGroup       uint32
}

func (*Ikev2SetIkeTransforms) GetMessageName() string {
	return "ikev2_set_ike_transforms"
}
func (*Ikev2SetIkeTransforms) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2SetIkeTransforms) GetCrcString() string {
	return "f0bf018e"
}
func NewIkev2SetIkeTransforms() api.Message {
	return &Ikev2SetIkeTransforms{}
}

// Ikev2SetIkeTransformsReply represents the VPP binary API message 'ikev2_set_ike_transforms_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 214:
//
//        ["ikev2_set_ike_transforms_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x18883302"}
//        ],
//
type Ikev2SetIkeTransformsReply struct {
	Retval int32
}

func (*Ikev2SetIkeTransformsReply) GetMessageName() string {
	return "ikev2_set_ike_transforms_reply"
}
func (*Ikev2SetIkeTransformsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2SetIkeTransformsReply) GetCrcString() string {
	return "18883302"
}
func NewIkev2SetIkeTransformsReply() api.Message {
	return &Ikev2SetIkeTransformsReply{}
}

// Ikev2SetEspTransforms represents the VPP binary API message 'ikev2_set_esp_transforms'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 220:
//
//        ["ikev2_set_esp_transforms",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u32", "crypto_alg"],
//            ["u32", "crypto_key_size"],
//            ["u32", "integ_alg"],
//            ["u32", "dh_group"],
//            {"crc" : "0xdadd693a"}
//        ],
//
type Ikev2SetEspTransforms struct {
	Name          []byte `struc:"[64]byte"`
	CryptoAlg     uint32
	CryptoKeySize uint32
	IntegAlg      uint32
	DhGroup       uint32
}

func (*Ikev2SetEspTransforms) GetMessageName() string {
	return "ikev2_set_esp_transforms"
}
func (*Ikev2SetEspTransforms) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2SetEspTransforms) GetCrcString() string {
	return "dadd693a"
}
func NewIkev2SetEspTransforms() api.Message {
	return &Ikev2SetEspTransforms{}
}

// Ikev2SetEspTransformsReply represents the VPP binary API message 'ikev2_set_esp_transforms_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 231:
//
//        ["ikev2_set_esp_transforms_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x9f3a1a67"}
//        ],
//
type Ikev2SetEspTransformsReply struct {
	Retval int32
}

func (*Ikev2SetEspTransformsReply) GetMessageName() string {
	return "ikev2_set_esp_transforms_reply"
}
func (*Ikev2SetEspTransformsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2SetEspTransformsReply) GetCrcString() string {
	return "9f3a1a67"
}
func NewIkev2SetEspTransformsReply() api.Message {
	return &Ikev2SetEspTransformsReply{}
}

// Ikev2SetSaLifetime represents the VPP binary API message 'ikev2_set_sa_lifetime'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 237:
//
//        ["ikev2_set_sa_lifetime",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            ["u64", "lifetime"],
//            ["u32", "lifetime_jitter"],
//            ["u32", "handover"],
//            ["u64", "lifetime_maxdata"],
//            {"crc" : "0x61b715bf"}
//        ],
//
type Ikev2SetSaLifetime struct {
	Name            []byte `struc:"[64]byte"`
	Lifetime        uint64
	LifetimeJitter  uint32
	Handover        uint32
	LifetimeMaxdata uint64
}

func (*Ikev2SetSaLifetime) GetMessageName() string {
	return "ikev2_set_sa_lifetime"
}
func (*Ikev2SetSaLifetime) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2SetSaLifetime) GetCrcString() string {
	return "61b715bf"
}
func NewIkev2SetSaLifetime() api.Message {
	return &Ikev2SetSaLifetime{}
}

// Ikev2SetSaLifetimeReply represents the VPP binary API message 'ikev2_set_sa_lifetime_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 248:
//
//        ["ikev2_set_sa_lifetime_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x7f73b9e3"}
//        ],
//
type Ikev2SetSaLifetimeReply struct {
	Retval int32
}

func (*Ikev2SetSaLifetimeReply) GetMessageName() string {
	return "ikev2_set_sa_lifetime_reply"
}
func (*Ikev2SetSaLifetimeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2SetSaLifetimeReply) GetCrcString() string {
	return "7f73b9e3"
}
func NewIkev2SetSaLifetimeReply() api.Message {
	return &Ikev2SetSaLifetimeReply{}
}

// Ikev2InitiateSaInit represents the VPP binary API message 'ikev2_initiate_sa_init'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 254:
//
//        ["ikev2_initiate_sa_init",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "name", 64],
//            {"crc" : "0x991b95f7"}
//        ],
//
type Ikev2InitiateSaInit struct {
	Name []byte `struc:"[64]byte"`
}

func (*Ikev2InitiateSaInit) GetMessageName() string {
	return "ikev2_initiate_sa_init"
}
func (*Ikev2InitiateSaInit) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2InitiateSaInit) GetCrcString() string {
	return "991b95f7"
}
func NewIkev2InitiateSaInit() api.Message {
	return &Ikev2InitiateSaInit{}
}

// Ikev2InitiateSaInitReply represents the VPP binary API message 'ikev2_initiate_sa_init_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 261:
//
//        ["ikev2_initiate_sa_init_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x66860cf0"}
//        ],
//
type Ikev2InitiateSaInitReply struct {
	Retval int32
}

func (*Ikev2InitiateSaInitReply) GetMessageName() string {
	return "ikev2_initiate_sa_init_reply"
}
func (*Ikev2InitiateSaInitReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2InitiateSaInitReply) GetCrcString() string {
	return "66860cf0"
}
func NewIkev2InitiateSaInitReply() api.Message {
	return &Ikev2InitiateSaInitReply{}
}

// Ikev2InitiateDelIkeSa represents the VPP binary API message 'ikev2_initiate_del_ike_sa'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 267:
//
//        ["ikev2_initiate_del_ike_sa",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u64", "ispi"],
//            {"crc" : "0x2816f367"}
//        ],
//
type Ikev2InitiateDelIkeSa struct {
	Ispi uint64
}

func (*Ikev2InitiateDelIkeSa) GetMessageName() string {
	return "ikev2_initiate_del_ike_sa"
}
func (*Ikev2InitiateDelIkeSa) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2InitiateDelIkeSa) GetCrcString() string {
	return "2816f367"
}
func NewIkev2InitiateDelIkeSa() api.Message {
	return &Ikev2InitiateDelIkeSa{}
}

// Ikev2InitiateDelIkeSaReply represents the VPP binary API message 'ikev2_initiate_del_ike_sa_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 274:
//
//        ["ikev2_initiate_del_ike_sa_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa5df9a20"}
//        ],
//
type Ikev2InitiateDelIkeSaReply struct {
	Retval int32
}

func (*Ikev2InitiateDelIkeSaReply) GetMessageName() string {
	return "ikev2_initiate_del_ike_sa_reply"
}
func (*Ikev2InitiateDelIkeSaReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2InitiateDelIkeSaReply) GetCrcString() string {
	return "a5df9a20"
}
func NewIkev2InitiateDelIkeSaReply() api.Message {
	return &Ikev2InitiateDelIkeSaReply{}
}

// Ikev2InitiateDelChildSa represents the VPP binary API message 'ikev2_initiate_del_child_sa'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 280:
//
//        ["ikev2_initiate_del_child_sa",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "ispi"],
//            {"crc" : "0x26eaba1d"}
//        ],
//
type Ikev2InitiateDelChildSa struct {
	Ispi uint32
}

func (*Ikev2InitiateDelChildSa) GetMessageName() string {
	return "ikev2_initiate_del_child_sa"
}
func (*Ikev2InitiateDelChildSa) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2InitiateDelChildSa) GetCrcString() string {
	return "26eaba1d"
}
func NewIkev2InitiateDelChildSa() api.Message {
	return &Ikev2InitiateDelChildSa{}
}

// Ikev2InitiateDelChildSaReply represents the VPP binary API message 'ikev2_initiate_del_child_sa_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 287:
//
//        ["ikev2_initiate_del_child_sa_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb31728ac"}
//        ],
//
type Ikev2InitiateDelChildSaReply struct {
	Retval int32
}

func (*Ikev2InitiateDelChildSaReply) GetMessageName() string {
	return "ikev2_initiate_del_child_sa_reply"
}
func (*Ikev2InitiateDelChildSaReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2InitiateDelChildSaReply) GetCrcString() string {
	return "b31728ac"
}
func NewIkev2InitiateDelChildSaReply() api.Message {
	return &Ikev2InitiateDelChildSaReply{}
}

// Ikev2InitiateRekeyChildSa represents the VPP binary API message 'ikev2_initiate_rekey_child_sa'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 293:
//
//        ["ikev2_initiate_rekey_child_sa",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "ispi"],
//            {"crc" : "0x48a0ad23"}
//        ],
//
type Ikev2InitiateRekeyChildSa struct {
	Ispi uint32
}

func (*Ikev2InitiateRekeyChildSa) GetMessageName() string {
	return "ikev2_initiate_rekey_child_sa"
}
func (*Ikev2InitiateRekeyChildSa) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Ikev2InitiateRekeyChildSa) GetCrcString() string {
	return "48a0ad23"
}
func NewIkev2InitiateRekeyChildSa() api.Message {
	return &Ikev2InitiateRekeyChildSa{}
}

// Ikev2InitiateRekeyChildSaReply represents the VPP binary API message 'ikev2_initiate_rekey_child_sa_reply'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 300:
//
//        ["ikev2_initiate_rekey_child_sa_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x7e5a8a19"}
//        ],
//
type Ikev2InitiateRekeyChildSaReply struct {
	Retval int32
}

func (*Ikev2InitiateRekeyChildSaReply) GetMessageName() string {
	return "ikev2_initiate_rekey_child_sa_reply"
}
func (*Ikev2InitiateRekeyChildSaReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Ikev2InitiateRekeyChildSaReply) GetCrcString() string {
	return "7e5a8a19"
}
func NewIkev2InitiateRekeyChildSaReply() api.Message {
	return &Ikev2InitiateRekeyChildSaReply{}
}

// IpsecSpdDump represents the VPP binary API message 'ipsec_spd_dump'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 306:
//
//        ["ipsec_spd_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "spd_id"],
//            ["u32", "sa_id"],
//            {"crc" : "0x5e9ae88e"}
//        ],
//
type IpsecSpdDump struct {
	SpdID uint32
	SaID  uint32
}

func (*IpsecSpdDump) GetMessageName() string {
	return "ipsec_spd_dump"
}
func (*IpsecSpdDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IpsecSpdDump) GetCrcString() string {
	return "5e9ae88e"
}
func NewIpsecSpdDump() api.Message {
	return &IpsecSpdDump{}
}

// IpsecSpdDetails represents the VPP binary API message 'ipsec_spd_details'.
// Generated from '/usr/share/vpp/api//ipsec.api.json', line 314:
//
//        ["ipsec_spd_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "spd_id"],
//            ["i32", "priority"],
//            ["u8", "is_outbound"],
//            ["u8", "is_ipv6"],
//            ["u8", "local_start_addr", 16],
//            ["u8", "local_stop_addr", 16],
//            ["u16", "local_start_port"],
//            ["u16", "local_stop_port"],
//            ["u8", "remote_start_addr", 16],
//            ["u8", "remote_stop_addr", 16],
//            ["u16", "remote_start_port"],
//            ["u16", "remote_stop_port"],
//            ["u8", "protocol"],
//            ["u8", "policy"],
//            ["u32", "sa_id"],
//            ["u64", "bytes"],
//            ["u64", "packets"],
//            {"crc" : "0x6f7821b0"}
//        ]
//
type IpsecSpdDetails struct {
	SpdID           uint32
	Priority        int32
	IsOutbound      uint8
	IsIpv6          uint8
	LocalStartAddr  []byte `struc:"[16]byte"`
	LocalStopAddr   []byte `struc:"[16]byte"`
	LocalStartPort  uint16
	LocalStopPort   uint16
	RemoteStartAddr []byte `struc:"[16]byte"`
	RemoteStopAddr  []byte `struc:"[16]byte"`
	RemoteStartPort uint16
	RemoteStopPort  uint16
	Protocol        uint8
	Policy          uint8
	SaID            uint32
	Bytes           uint64
	Packets         uint64
}

func (*IpsecSpdDetails) GetMessageName() string {
	return "ipsec_spd_details"
}
func (*IpsecSpdDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IpsecSpdDetails) GetCrcString() string {
	return "6f7821b0"
}
func NewIpsecSpdDetails() api.Message {
	return &IpsecSpdDetails{}
}
