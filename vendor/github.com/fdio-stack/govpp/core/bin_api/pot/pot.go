// Package pot represents the VPP binary API of the 'pot' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//pot.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package pot

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xd9cfc29f

// PotProfileAdd represents the VPP binary API message 'pot_profile_add'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 6:
//
//        ["pot_profile_add",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "id"],
//            ["u8", "validator"],
//            ["u64", "secret_key"],
//            ["u64", "secret_share"],
//            ["u64", "prime"],
//            ["u8", "max_bits"],
//            ["u64", "lpc"],
//            ["u64", "polynomial_public"],
//            ["u8", "list_name_len"],
//            ["u8", "list_name", 0],
//            {"crc" : "0x0e44e40d"}
//        ],
//
type PotProfileAdd struct {
	ID               uint8
	Validator        uint8
	SecretKey        uint64
	SecretShare      uint64
	Prime            uint64
	MaxBits          uint8
	Lpc              uint64
	PolynomialPublic uint64
	ListNameLen      uint8
	ListName         []byte
}

func (*PotProfileAdd) GetMessageName() string {
	return "pot_profile_add"
}
func (*PotProfileAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PotProfileAdd) GetCrcString() string {
	return "0e44e40d"
}
func NewPotProfileAdd() api.Message {
	return &PotProfileAdd{}
}

// PotProfileAddReply represents the VPP binary API message 'pot_profile_add_reply'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 22:
//
//        ["pot_profile_add_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa1af9d17"}
//        ],
//
type PotProfileAddReply struct {
	Retval int32
}

func (*PotProfileAddReply) GetMessageName() string {
	return "pot_profile_add_reply"
}
func (*PotProfileAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PotProfileAddReply) GetCrcString() string {
	return "a1af9d17"
}
func NewPotProfileAddReply() api.Message {
	return &PotProfileAddReply{}
}

// PotProfileActivate represents the VPP binary API message 'pot_profile_activate'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 28:
//
//        ["pot_profile_activate",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "id"],
//            ["u8", "list_name_len"],
//            ["u8", "list_name", 0],
//            {"crc" : "0x8bb6738e"}
//        ],
//
type PotProfileActivate struct {
	ID          uint8
	ListNameLen uint8
	ListName    []byte
}

func (*PotProfileActivate) GetMessageName() string {
	return "pot_profile_activate"
}
func (*PotProfileActivate) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PotProfileActivate) GetCrcString() string {
	return "8bb6738e"
}
func NewPotProfileActivate() api.Message {
	return &PotProfileActivate{}
}

// PotProfileActivateReply represents the VPP binary API message 'pot_profile_activate_reply'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 37:
//
//        ["pot_profile_activate_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x3c0680ae"}
//        ],
//
type PotProfileActivateReply struct {
	Retval int32
}

func (*PotProfileActivateReply) GetMessageName() string {
	return "pot_profile_activate_reply"
}
func (*PotProfileActivateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PotProfileActivateReply) GetCrcString() string {
	return "3c0680ae"
}
func NewPotProfileActivateReply() api.Message {
	return &PotProfileActivateReply{}
}

// PotProfileDel represents the VPP binary API message 'pot_profile_del'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 43:
//
//        ["pot_profile_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "list_name_len"],
//            ["u8", "list_name", 0],
//            {"crc" : "0xd796848c"}
//        ],
//
type PotProfileDel struct {
	ListNameLen uint8
	ListName    []byte
}

func (*PotProfileDel) GetMessageName() string {
	return "pot_profile_del"
}
func (*PotProfileDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PotProfileDel) GetCrcString() string {
	return "d796848c"
}
func NewPotProfileDel() api.Message {
	return &PotProfileDel{}
}

// PotProfileDelReply represents the VPP binary API message 'pot_profile_del_reply'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 51:
//
//        ["pot_profile_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x8d1e11f5"}
//        ],
//
type PotProfileDelReply struct {
	Retval int32
}

func (*PotProfileDelReply) GetMessageName() string {
	return "pot_profile_del_reply"
}
func (*PotProfileDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PotProfileDelReply) GetCrcString() string {
	return "8d1e11f5"
}
func NewPotProfileDelReply() api.Message {
	return &PotProfileDelReply{}
}

// PotProfileShowConfigDump represents the VPP binary API message 'pot_profile_show_config_dump'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 57:
//
//        ["pot_profile_show_config_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "id"],
//            {"crc" : "0x8ba1cb4d"}
//        ],
//
type PotProfileShowConfigDump struct {
	ID uint8
}

func (*PotProfileShowConfigDump) GetMessageName() string {
	return "pot_profile_show_config_dump"
}
func (*PotProfileShowConfigDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PotProfileShowConfigDump) GetCrcString() string {
	return "8ba1cb4d"
}
func NewPotProfileShowConfigDump() api.Message {
	return &PotProfileShowConfigDump{}
}

// PotProfileShowConfigDetails represents the VPP binary API message 'pot_profile_show_config_details'.
// Generated from '/usr/share/vpp/api//pot.api.json', line 64:
//
//        ["pot_profile_show_config_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "id"],
//            ["u8", "validator"],
//            ["u64", "secret_key"],
//            ["u64", "secret_share"],
//            ["u64", "prime"],
//            ["u64", "bit_mask"],
//            ["u64", "lpc"],
//            ["u64", "polynomial_public"],
//            {"crc" : "0xa9f1800f"}
//        ]
//
type PotProfileShowConfigDetails struct {
	Retval           int32
	ID               uint8
	Validator        uint8
	SecretKey        uint64
	SecretShare      uint64
	Prime            uint64
	BitMask          uint64
	Lpc              uint64
	PolynomialPublic uint64
}

func (*PotProfileShowConfigDetails) GetMessageName() string {
	return "pot_profile_show_config_details"
}
func (*PotProfileShowConfigDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PotProfileShowConfigDetails) GetCrcString() string {
	return "a9f1800f"
}
func NewPotProfileShowConfigDetails() api.Message {
	return &PotProfileShowConfigDetails{}
}
