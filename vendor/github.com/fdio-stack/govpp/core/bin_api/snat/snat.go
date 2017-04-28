// Package snat represents the VPP binary API of the 'snat' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//snat.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package snat

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xab353caa

// SnatAddAddressRange represents the VPP binary API message 'snat_add_address_range'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 6:
//
//        ["snat_add_address_range",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "first_ip_address", 16],
//            ["u8", "last_ip_address", 16],
//            ["u32", "vrf_id"],
//            ["u8", "is_add"],
//            {"crc" : "0xd45f0a85"}
//        ],
//
type SnatAddAddressRange struct {
	IsIP4          uint8
	FirstIPAddress []byte `struc:"[16]byte"`
	LastIPAddress  []byte `struc:"[16]byte"`
	VrfID          uint32
	IsAdd          uint8
}

func (*SnatAddAddressRange) GetMessageName() string {
	return "snat_add_address_range"
}
func (*SnatAddAddressRange) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatAddAddressRange) GetCrcString() string {
	return "d45f0a85"
}
func NewSnatAddAddressRange() api.Message {
	return &SnatAddAddressRange{}
}

// SnatAddAddressRangeReply represents the VPP binary API message 'snat_add_address_range_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 17:
//
//        ["snat_add_address_range_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x7d360766"}
//        ],
//
type SnatAddAddressRangeReply struct {
	Retval int32
}

func (*SnatAddAddressRangeReply) GetMessageName() string {
	return "snat_add_address_range_reply"
}
func (*SnatAddAddressRangeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatAddAddressRangeReply) GetCrcString() string {
	return "7d360766"
}
func NewSnatAddAddressRangeReply() api.Message {
	return &SnatAddAddressRangeReply{}
}

// SnatAddressDump represents the VPP binary API message 'snat_address_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 23:
//
//        ["snat_address_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x27c92402"}
//        ],
//
type SnatAddressDump struct {
}

func (*SnatAddressDump) GetMessageName() string {
	return "snat_address_dump"
}
func (*SnatAddressDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatAddressDump) GetCrcString() string {
	return "27c92402"
}
func NewSnatAddressDump() api.Message {
	return &SnatAddressDump{}
}

// SnatAddressDetails represents the VPP binary API message 'snat_address_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 29:
//
//        ["snat_address_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "ip_address", 16],
//            ["u32", "vrf_id"],
//            {"crc" : "0xf4e07104"}
//        ],
//
type SnatAddressDetails struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	VrfID     uint32
}

func (*SnatAddressDetails) GetMessageName() string {
	return "snat_address_details"
}
func (*SnatAddressDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatAddressDetails) GetCrcString() string {
	return "f4e07104"
}
func NewSnatAddressDetails() api.Message {
	return &SnatAddressDetails{}
}

// SnatInterfaceAddDelFeature represents the VPP binary API message 'snat_interface_add_del_feature'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 37:
//
//        ["snat_interface_add_del_feature",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_inside"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xdd74ea24"}
//        ],
//
type SnatInterfaceAddDelFeature struct {
	IsAdd     uint8
	IsInside  uint8
	SwIfIndex uint32
}

func (*SnatInterfaceAddDelFeature) GetMessageName() string {
	return "snat_interface_add_del_feature"
}
func (*SnatInterfaceAddDelFeature) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatInterfaceAddDelFeature) GetCrcString() string {
	return "dd74ea24"
}
func NewSnatInterfaceAddDelFeature() api.Message {
	return &SnatInterfaceAddDelFeature{}
}

// SnatInterfaceAddDelFeatureReply represents the VPP binary API message 'snat_interface_add_del_feature_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 46:
//
//        ["snat_interface_add_del_feature_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xabd1e17f"}
//        ],
//
type SnatInterfaceAddDelFeatureReply struct {
	Retval int32
}

func (*SnatInterfaceAddDelFeatureReply) GetMessageName() string {
	return "snat_interface_add_del_feature_reply"
}
func (*SnatInterfaceAddDelFeatureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatInterfaceAddDelFeatureReply) GetCrcString() string {
	return "abd1e17f"
}
func NewSnatInterfaceAddDelFeatureReply() api.Message {
	return &SnatInterfaceAddDelFeatureReply{}
}

// SnatInterfaceDump represents the VPP binary API message 'snat_interface_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 52:
//
//        ["snat_interface_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x01fe9590"}
//        ],
//
type SnatInterfaceDump struct {
}

func (*SnatInterfaceDump) GetMessageName() string {
	return "snat_interface_dump"
}
func (*SnatInterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatInterfaceDump) GetCrcString() string {
	return "01fe9590"
}
func NewSnatInterfaceDump() api.Message {
	return &SnatInterfaceDump{}
}

// SnatInterfaceDetails represents the VPP binary API message 'snat_interface_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 58:
//
//        ["snat_interface_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_inside"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x8e43ab7a"}
//        ],
//
type SnatInterfaceDetails struct {
	IsInside  uint8
	SwIfIndex uint32
}

func (*SnatInterfaceDetails) GetMessageName() string {
	return "snat_interface_details"
}
func (*SnatInterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatInterfaceDetails) GetCrcString() string {
	return "8e43ab7a"
}
func NewSnatInterfaceDetails() api.Message {
	return &SnatInterfaceDetails{}
}

// SnatAddStaticMapping represents the VPP binary API message 'snat_add_static_mapping'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 65:
//
//        ["snat_add_static_mapping",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_ip4"],
//            ["u8", "addr_only"],
//            ["u8", "local_ip_address", 16],
//            ["u8", "external_ip_address", 16],
//            ["u8", "protocol"],
//            ["u16", "local_port"],
//            ["u16", "external_port"],
//            ["u32", "external_sw_if_index"],
//            ["u32", "vrf_id"],
//            {"crc" : "0xd478ee49"}
//        ],
//
type SnatAddStaticMapping struct {
	IsAdd             uint8
	IsIP4             uint8
	AddrOnly          uint8
	LocalIPAddress    []byte `struc:"[16]byte"`
	ExternalIPAddress []byte `struc:"[16]byte"`
	Protocol          uint8
	LocalPort         uint16
	ExternalPort      uint16
	ExternalSwIfIndex uint32
	VrfID             uint32
}

func (*SnatAddStaticMapping) GetMessageName() string {
	return "snat_add_static_mapping"
}
func (*SnatAddStaticMapping) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatAddStaticMapping) GetCrcString() string {
	return "d478ee49"
}
func NewSnatAddStaticMapping() api.Message {
	return &SnatAddStaticMapping{}
}

// SnatAddStaticMappingReply represents the VPP binary API message 'snat_add_static_mapping_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 81:
//
//        ["snat_add_static_mapping_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xc9db5568"}
//        ],
//
type SnatAddStaticMappingReply struct {
	Retval int32
}

func (*SnatAddStaticMappingReply) GetMessageName() string {
	return "snat_add_static_mapping_reply"
}
func (*SnatAddStaticMappingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatAddStaticMappingReply) GetCrcString() string {
	return "c9db5568"
}
func NewSnatAddStaticMappingReply() api.Message {
	return &SnatAddStaticMappingReply{}
}

// SnatStaticMappingDump represents the VPP binary API message 'snat_static_mapping_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 87:
//
//        ["snat_static_mapping_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x84172d2b"}
//        ],
//
type SnatStaticMappingDump struct {
}

func (*SnatStaticMappingDump) GetMessageName() string {
	return "snat_static_mapping_dump"
}
func (*SnatStaticMappingDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatStaticMappingDump) GetCrcString() string {
	return "84172d2b"
}
func NewSnatStaticMappingDump() api.Message {
	return &SnatStaticMappingDump{}
}

// SnatStaticMappingDetails represents the VPP binary API message 'snat_static_mapping_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 93:
//
//        ["snat_static_mapping_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "addr_only"],
//            ["u8", "local_ip_address", 16],
//            ["u8", "external_ip_address", 16],
//            ["u8", "protocol"],
//            ["u16", "local_port"],
//            ["u16", "external_port"],
//            ["u32", "external_sw_if_index"],
//            ["u32", "vrf_id"],
//            {"crc" : "0x5e4d17bf"}
//        ],
//
type SnatStaticMappingDetails struct {
	IsIP4             uint8
	AddrOnly          uint8
	LocalIPAddress    []byte `struc:"[16]byte"`
	ExternalIPAddress []byte `struc:"[16]byte"`
	Protocol          uint8
	LocalPort         uint16
	ExternalPort      uint16
	ExternalSwIfIndex uint32
	VrfID             uint32
}

func (*SnatStaticMappingDetails) GetMessageName() string {
	return "snat_static_mapping_details"
}
func (*SnatStaticMappingDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatStaticMappingDetails) GetCrcString() string {
	return "5e4d17bf"
}
func NewSnatStaticMappingDetails() api.Message {
	return &SnatStaticMappingDetails{}
}

// SnatControlPing represents the VPP binary API message 'snat_control_ping'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 107:
//
//        ["snat_control_ping",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x7a3d60f6"}
//        ],
//
type SnatControlPing struct {
}

func (*SnatControlPing) GetMessageName() string {
	return "snat_control_ping"
}
func (*SnatControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatControlPing) GetCrcString() string {
	return "7a3d60f6"
}
func NewSnatControlPing() api.Message {
	return &SnatControlPing{}
}

// SnatControlPingReply represents the VPP binary API message 'snat_control_ping_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 113:
//
//        ["snat_control_ping_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "client_index"],
//            ["u32", "vpe_pid"],
//            {"crc" : "0x979a5618"}
//        ],
//
type SnatControlPingReply struct {
	Retval      int32
	ClientIndex uint32
	VpePid      uint32
}

func (*SnatControlPingReply) GetMessageName() string {
	return "snat_control_ping_reply"
}
func (*SnatControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatControlPingReply) GetCrcString() string {
	return "979a5618"
}
func NewSnatControlPingReply() api.Message {
	return &SnatControlPingReply{}
}

// SnatShowConfig represents the VPP binary API message 'snat_show_config'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 121:
//
//        ["snat_show_config",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x2c6077c4"}
//        ],
//
type SnatShowConfig struct {
}

func (*SnatShowConfig) GetMessageName() string {
	return "snat_show_config"
}
func (*SnatShowConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatShowConfig) GetCrcString() string {
	return "2c6077c4"
}
func NewSnatShowConfig() api.Message {
	return &SnatShowConfig{}
}

// SnatShowConfigReply represents the VPP binary API message 'snat_show_config_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 127:
//
//        ["snat_show_config_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "static_mapping_only"],
//            ["u8", "static_mapping_connection_tracking"],
//            ["u8", "deterministic"],
//            ["u32", "translation_buckets"],
//            ["u32", "translation_memory_size"],
//            ["u32", "user_buckets"],
//            ["u32", "user_memory_size"],
//            ["u32", "max_translations_per_user"],
//            ["u32", "outside_vrf_id"],
//            ["u32", "inside_vrf_id"],
//            {"crc" : "0xbf9c4060"}
//        ],
//
type SnatShowConfigReply struct {
	Retval                          int32
	StaticMappingOnly               uint8
	StaticMappingConnectionTracking uint8
	Deterministic                   uint8
	TranslationBuckets              uint32
	TranslationMemorySize           uint32
	UserBuckets                     uint32
	UserMemorySize                  uint32
	MaxTranslationsPerUser          uint32
	OutsideVrfID                    uint32
	InsideVrfID                     uint32
}

func (*SnatShowConfigReply) GetMessageName() string {
	return "snat_show_config_reply"
}
func (*SnatShowConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatShowConfigReply) GetCrcString() string {
	return "bf9c4060"
}
func NewSnatShowConfigReply() api.Message {
	return &SnatShowConfigReply{}
}

// SnatSetWorkers represents the VPP binary API message 'snat_set_workers'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 143:
//
//        ["snat_set_workers",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u64", "worker_mask"],
//            {"crc" : "0xe884f45e"}
//        ],
//
type SnatSetWorkers struct {
	WorkerMask uint64
}

func (*SnatSetWorkers) GetMessageName() string {
	return "snat_set_workers"
}
func (*SnatSetWorkers) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatSetWorkers) GetCrcString() string {
	return "e884f45e"
}
func NewSnatSetWorkers() api.Message {
	return &SnatSetWorkers{}
}

// SnatSetWorkersReply represents the VPP binary API message 'snat_set_workers_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 150:
//
//        ["snat_set_workers_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x47fb5f11"}
//        ],
//
type SnatSetWorkersReply struct {
	Retval int32
}

func (*SnatSetWorkersReply) GetMessageName() string {
	return "snat_set_workers_reply"
}
func (*SnatSetWorkersReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatSetWorkersReply) GetCrcString() string {
	return "47fb5f11"
}
func NewSnatSetWorkersReply() api.Message {
	return &SnatSetWorkersReply{}
}

// SnatWorkerDump represents the VPP binary API message 'snat_worker_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 156:
//
//        ["snat_worker_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xb7593228"}
//        ],
//
type SnatWorkerDump struct {
}

func (*SnatWorkerDump) GetMessageName() string {
	return "snat_worker_dump"
}
func (*SnatWorkerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatWorkerDump) GetCrcString() string {
	return "b7593228"
}
func NewSnatWorkerDump() api.Message {
	return &SnatWorkerDump{}
}

// SnatWorkerDetails represents the VPP binary API message 'snat_worker_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 162:
//
//        ["snat_worker_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "worker_index"],
//            ["u32", "lcore_id"],
//            ["u8", "name", 64],
//            {"crc" : "0x6a1d1344"}
//        ],
//
type SnatWorkerDetails struct {
	WorkerIndex uint32
	LcoreID     uint32
	Name        []byte `struc:"[64]byte"`
}

func (*SnatWorkerDetails) GetMessageName() string {
	return "snat_worker_details"
}
func (*SnatWorkerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatWorkerDetails) GetCrcString() string {
	return "6a1d1344"
}
func NewSnatWorkerDetails() api.Message {
	return &SnatWorkerDetails{}
}

// SnatAddDelInterfaceAddr represents the VPP binary API message 'snat_add_del_interface_addr'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 170:
//
//        ["snat_add_del_interface_addr",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_inside"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x42b97914"}
//        ],
//
type SnatAddDelInterfaceAddr struct {
	IsAdd     uint8
	IsInside  uint8
	SwIfIndex uint32
}

func (*SnatAddDelInterfaceAddr) GetMessageName() string {
	return "snat_add_del_interface_addr"
}
func (*SnatAddDelInterfaceAddr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatAddDelInterfaceAddr) GetCrcString() string {
	return "42b97914"
}
func NewSnatAddDelInterfaceAddr() api.Message {
	return &SnatAddDelInterfaceAddr{}
}

// SnatAddDelInterfaceAddrReply represents the VPP binary API message 'snat_add_del_interface_addr_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 179:
//
//        ["snat_add_del_interface_addr_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x8d0be39e"}
//        ],
//
type SnatAddDelInterfaceAddrReply struct {
	Retval int32
}

func (*SnatAddDelInterfaceAddrReply) GetMessageName() string {
	return "snat_add_del_interface_addr_reply"
}
func (*SnatAddDelInterfaceAddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatAddDelInterfaceAddrReply) GetCrcString() string {
	return "8d0be39e"
}
func NewSnatAddDelInterfaceAddrReply() api.Message {
	return &SnatAddDelInterfaceAddrReply{}
}

// SnatInterfaceAddrDump represents the VPP binary API message 'snat_interface_addr_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 185:
//
//        ["snat_interface_addr_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x3e801286"}
//        ],
//
type SnatInterfaceAddrDump struct {
}

func (*SnatInterfaceAddrDump) GetMessageName() string {
	return "snat_interface_addr_dump"
}
func (*SnatInterfaceAddrDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatInterfaceAddrDump) GetCrcString() string {
	return "3e801286"
}
func NewSnatInterfaceAddrDump() api.Message {
	return &SnatInterfaceAddrDump{}
}

// SnatInterfaceAddrDetails represents the VPP binary API message 'snat_interface_addr_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 191:
//
//        ["snat_interface_addr_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xd65481ed"}
//        ],
//
type SnatInterfaceAddrDetails struct {
	SwIfIndex uint32
}

func (*SnatInterfaceAddrDetails) GetMessageName() string {
	return "snat_interface_addr_details"
}
func (*SnatInterfaceAddrDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatInterfaceAddrDetails) GetCrcString() string {
	return "d65481ed"
}
func NewSnatInterfaceAddrDetails() api.Message {
	return &SnatInterfaceAddrDetails{}
}

// SnatIpfixEnableDisable represents the VPP binary API message 'snat_ipfix_enable_disable'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 197:
//
//        ["snat_ipfix_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "domain_id"],
//            ["u16", "src_port"],
//            ["u8", "enable"],
//            {"crc" : "0x88f422c2"}
//        ],
//
type SnatIpfixEnableDisable struct {
	DomainID uint32
	SrcPort  uint16
	Enable   uint8
}

func (*SnatIpfixEnableDisable) GetMessageName() string {
	return "snat_ipfix_enable_disable"
}
func (*SnatIpfixEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatIpfixEnableDisable) GetCrcString() string {
	return "88f422c2"
}
func NewSnatIpfixEnableDisable() api.Message {
	return &SnatIpfixEnableDisable{}
}

// SnatIpfixEnableDisableReply represents the VPP binary API message 'snat_ipfix_enable_disable_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 206:
//
//        ["snat_ipfix_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa7f6c343"}
//        ],
//
type SnatIpfixEnableDisableReply struct {
	Retval int32
}

func (*SnatIpfixEnableDisableReply) GetMessageName() string {
	return "snat_ipfix_enable_disable_reply"
}
func (*SnatIpfixEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatIpfixEnableDisableReply) GetCrcString() string {
	return "a7f6c343"
}
func NewSnatIpfixEnableDisableReply() api.Message {
	return &SnatIpfixEnableDisableReply{}
}

// SnatUserDump represents the VPP binary API message 'snat_user_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 212:
//
//        ["snat_user_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x062dc900"}
//        ],
//
type SnatUserDump struct {
}

func (*SnatUserDump) GetMessageName() string {
	return "snat_user_dump"
}
func (*SnatUserDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatUserDump) GetCrcString() string {
	return "062dc900"
}
func NewSnatUserDump() api.Message {
	return &SnatUserDump{}
}

// SnatUserDetails represents the VPP binary API message 'snat_user_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 218:
//
//        ["snat_user_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "vrf_id"],
//            ["u8", "is_ip4"],
//            ["u8", "ip_address", 16],
//            ["u32", "nsessions"],
//            ["u32", "nstaticsessions"],
//            {"crc" : "0xfe7dad8e"}
//        ],
//
type SnatUserDetails struct {
	VrfID           uint32
	IsIP4           uint8
	IPAddress       []byte `struc:"[16]byte"`
	Nsessions       uint32
	Nstaticsessions uint32
}

func (*SnatUserDetails) GetMessageName() string {
	return "snat_user_details"
}
func (*SnatUserDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatUserDetails) GetCrcString() string {
	return "fe7dad8e"
}
func NewSnatUserDetails() api.Message {
	return &SnatUserDetails{}
}

// SnatUserSessionDump represents the VPP binary API message 'snat_user_session_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 228:
//
//        ["snat_user_session_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "ip_address", 16],
//            ["u32", "vrf_id"],
//            {"crc" : "0x98bd1a88"}
//        ],
//
type SnatUserSessionDump struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	VrfID     uint32
}

func (*SnatUserSessionDump) GetMessageName() string {
	return "snat_user_session_dump"
}
func (*SnatUserSessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatUserSessionDump) GetCrcString() string {
	return "98bd1a88"
}
func NewSnatUserSessionDump() api.Message {
	return &SnatUserSessionDump{}
}

// SnatUserSessionDetails represents the VPP binary API message 'snat_user_session_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 237:
//
//        ["snat_user_session_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "outside_ip_address", 16],
//            ["u16", "outside_port"],
//            ["u8", "inside_ip_address", 16],
//            ["u16", "inside_port"],
//            ["u16", "protocol"],
//            ["u8", "is_static"],
//            ["u64", "last_heard"],
//            ["u64", "total_bytes"],
//            ["u32", "total_pkts"],
//            {"crc" : "0x1a632582"}
//        ],
//
type SnatUserSessionDetails struct {
	IsIP4            uint8
	OutsideIPAddress []byte `struc:"[16]byte"`
	OutsidePort      uint16
	InsideIPAddress  []byte `struc:"[16]byte"`
	InsidePort       uint16
	Protocol         uint16
	IsStatic         uint8
	LastHeard        uint64
	TotalBytes       uint64
	TotalPkts        uint32
}

func (*SnatUserSessionDetails) GetMessageName() string {
	return "snat_user_session_details"
}
func (*SnatUserSessionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatUserSessionDetails) GetCrcString() string {
	return "1a632582"
}
func NewSnatUserSessionDetails() api.Message {
	return &SnatUserSessionDetails{}
}

// SnatAddDetMap represents the VPP binary API message 'snat_add_det_map'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 252:
//
//        ["snat_add_det_map",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "is_ip4"],
//            ["u8", "addr_only"],
//            ["u8", "in_addr", 16],
//            ["u8", "in_plen"],
//            ["u8", "out_addr", 16],
//            ["u8", "out_plen"],
//            {"crc" : "0xb62c5cfd"}
//        ],
//
type SnatAddDetMap struct {
	IsAdd    uint8
	IsIP4    uint8
	AddrOnly uint8
	InAddr   []byte `struc:"[16]byte"`
	InPlen   uint8
	OutAddr  []byte `struc:"[16]byte"`
	OutPlen  uint8
}

func (*SnatAddDetMap) GetMessageName() string {
	return "snat_add_det_map"
}
func (*SnatAddDetMap) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatAddDetMap) GetCrcString() string {
	return "b62c5cfd"
}
func NewSnatAddDetMap() api.Message {
	return &SnatAddDetMap{}
}

// SnatAddDetMapReply represents the VPP binary API message 'snat_add_det_map_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 265:
//
//        ["snat_add_det_map_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd84be59c"}
//        ],
//
type SnatAddDetMapReply struct {
	Retval int32
}

func (*SnatAddDetMapReply) GetMessageName() string {
	return "snat_add_det_map_reply"
}
func (*SnatAddDetMapReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatAddDetMapReply) GetCrcString() string {
	return "d84be59c"
}
func NewSnatAddDetMapReply() api.Message {
	return &SnatAddDetMapReply{}
}

// SnatDetForward represents the VPP binary API message 'snat_det_forward'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 271:
//
//        ["snat_det_forward",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "in_addr", 16],
//            {"crc" : "0x8584635a"}
//        ],
//
type SnatDetForward struct {
	IsIP4  uint8
	InAddr []byte `struc:"[16]byte"`
}

func (*SnatDetForward) GetMessageName() string {
	return "snat_det_forward"
}
func (*SnatDetForward) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetForward) GetCrcString() string {
	return "8584635a"
}
func NewSnatDetForward() api.Message {
	return &SnatDetForward{}
}

// SnatDetForwardReply represents the VPP binary API message 'snat_det_forward_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 279:
//
//        ["snat_det_forward_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u16", "out_port_lo"],
//            ["u16", "out_port_hi"],
//            ["u8", "is_ip4"],
//            ["u8", "out_addr", 16],
//            {"crc" : "0xc04ce963"}
//        ],
//
type SnatDetForwardReply struct {
	Retval    int32
	OutPortLo uint16
	OutPortHi uint16
	IsIP4     uint8
	OutAddr   []byte `struc:"[16]byte"`
}

func (*SnatDetForwardReply) GetMessageName() string {
	return "snat_det_forward_reply"
}
func (*SnatDetForwardReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatDetForwardReply) GetCrcString() string {
	return "c04ce963"
}
func NewSnatDetForwardReply() api.Message {
	return &SnatDetForwardReply{}
}

// SnatDetReverse represents the VPP binary API message 'snat_det_reverse'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 289:
//
//        ["snat_det_reverse",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u16", "out_port"],
//            ["u8", "is_ip4"],
//            ["u8", "out_addr", 16],
//            {"crc" : "0xf6c78383"}
//        ],
//
type SnatDetReverse struct {
	OutPort uint16
	IsIP4   uint8
	OutAddr []byte `struc:"[16]byte"`
}

func (*SnatDetReverse) GetMessageName() string {
	return "snat_det_reverse"
}
func (*SnatDetReverse) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetReverse) GetCrcString() string {
	return "f6c78383"
}
func NewSnatDetReverse() api.Message {
	return &SnatDetReverse{}
}

// SnatDetReverseReply represents the VPP binary API message 'snat_det_reverse_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 298:
//
//        ["snat_det_reverse_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "is_ip4"],
//            ["u8", "in_addr", 16],
//            {"crc" : "0x0b0c08d5"}
//        ],
//
type SnatDetReverseReply struct {
	Retval int32
	IsIP4  uint8
	InAddr []byte `struc:"[16]byte"`
}

func (*SnatDetReverseReply) GetMessageName() string {
	return "snat_det_reverse_reply"
}
func (*SnatDetReverseReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatDetReverseReply) GetCrcString() string {
	return "0b0c08d5"
}
func NewSnatDetReverseReply() api.Message {
	return &SnatDetReverseReply{}
}

// SnatDetMapDump represents the VPP binary API message 'snat_det_map_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 306:
//
//        ["snat_det_map_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x9da7fd51"}
//        ],
//
type SnatDetMapDump struct {
}

func (*SnatDetMapDump) GetMessageName() string {
	return "snat_det_map_dump"
}
func (*SnatDetMapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetMapDump) GetCrcString() string {
	return "9da7fd51"
}
func NewSnatDetMapDump() api.Message {
	return &SnatDetMapDump{}
}

// SnatDetMapDetails represents the VPP binary API message 'snat_det_map_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 312:
//
//        ["snat_det_map_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "in_addr", 16],
//            ["u8", "in_plen"],
//            ["u8", "out_addr", 16],
//            ["u8", "out_plen"],
//            ["u32", "sharing_ratio"],
//            ["u16", "ports_per_host"],
//            ["u32", "ses_num"],
//            {"crc" : "0x0b0ddd08"}
//        ],
//
type SnatDetMapDetails struct {
	IsIP4        uint8
	InAddr       []byte `struc:"[16]byte"`
	InPlen       uint8
	OutAddr      []byte `struc:"[16]byte"`
	OutPlen      uint8
	SharingRatio uint32
	PortsPerHost uint16
	SesNum       uint32
}

func (*SnatDetMapDetails) GetMessageName() string {
	return "snat_det_map_details"
}
func (*SnatDetMapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatDetMapDetails) GetCrcString() string {
	return "0b0ddd08"
}
func NewSnatDetMapDetails() api.Message {
	return &SnatDetMapDetails{}
}

// SnatDetSetTimeouts represents the VPP binary API message 'snat_det_set_timeouts'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 325:
//
//        ["snat_det_set_timeouts",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "udp"],
//            ["u32", "tcp_established"],
//            ["u32", "tcp_transitory"],
//            ["u32", "icmp"],
//            {"crc" : "0xfe713877"}
//        ],
//
type SnatDetSetTimeouts struct {
	UDP            uint32
	TCPEstablished uint32
	TCPTransitory  uint32
	ICMP           uint32
}

func (*SnatDetSetTimeouts) GetMessageName() string {
	return "snat_det_set_timeouts"
}
func (*SnatDetSetTimeouts) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetSetTimeouts) GetCrcString() string {
	return "fe713877"
}
func NewSnatDetSetTimeouts() api.Message {
	return &SnatDetSetTimeouts{}
}

// SnatDetSetTimeoutsReply represents the VPP binary API message 'snat_det_set_timeouts_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 335:
//
//        ["snat_det_set_timeouts_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x194bb7e2"}
//        ],
//
type SnatDetSetTimeoutsReply struct {
	Retval int32
}

func (*SnatDetSetTimeoutsReply) GetMessageName() string {
	return "snat_det_set_timeouts_reply"
}
func (*SnatDetSetTimeoutsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatDetSetTimeoutsReply) GetCrcString() string {
	return "194bb7e2"
}
func NewSnatDetSetTimeoutsReply() api.Message {
	return &SnatDetSetTimeoutsReply{}
}

// SnatDetGetTimeouts represents the VPP binary API message 'snat_det_get_timeouts'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 341:
//
//        ["snat_det_get_timeouts",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x834fa3bc"}
//        ],
//
type SnatDetGetTimeouts struct {
}

func (*SnatDetGetTimeouts) GetMessageName() string {
	return "snat_det_get_timeouts"
}
func (*SnatDetGetTimeouts) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetGetTimeouts) GetCrcString() string {
	return "834fa3bc"
}
func NewSnatDetGetTimeouts() api.Message {
	return &SnatDetGetTimeouts{}
}

// SnatDetGetTimeoutsReply represents the VPP binary API message 'snat_det_get_timeouts_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 347:
//
//        ["snat_det_get_timeouts_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "udp"],
//            ["u32", "tcp_established"],
//            ["u32", "tcp_transitory"],
//            ["u32", "icmp"],
//            {"crc" : "0x4d03d12e"}
//        ],
//
type SnatDetGetTimeoutsReply struct {
	Retval         int32
	UDP            uint32
	TCPEstablished uint32
	TCPTransitory  uint32
	ICMP           uint32
}

func (*SnatDetGetTimeoutsReply) GetMessageName() string {
	return "snat_det_get_timeouts_reply"
}
func (*SnatDetGetTimeoutsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatDetGetTimeoutsReply) GetCrcString() string {
	return "4d03d12e"
}
func NewSnatDetGetTimeoutsReply() api.Message {
	return &SnatDetGetTimeoutsReply{}
}

// SnatDetCloseSessionOut represents the VPP binary API message 'snat_det_close_session_out'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 357:
//
//        ["snat_det_close_session_out",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "out_addr", 16],
//            ["u16", "out_port"],
//            ["u8", "ext_addr", 16],
//            ["u16", "ext_port"],
//            {"crc" : "0xaacd5d95"}
//        ],
//
type SnatDetCloseSessionOut struct {
	IsIP4   uint8
	OutAddr []byte `struc:"[16]byte"`
	OutPort uint16
	ExtAddr []byte `struc:"[16]byte"`
	ExtPort uint16
}

func (*SnatDetCloseSessionOut) GetMessageName() string {
	return "snat_det_close_session_out"
}
func (*SnatDetCloseSessionOut) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetCloseSessionOut) GetCrcString() string {
	return "aacd5d95"
}
func NewSnatDetCloseSessionOut() api.Message {
	return &SnatDetCloseSessionOut{}
}

// SnatDetCloseSessionOutReply represents the VPP binary API message 'snat_det_close_session_out_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 368:
//
//        ["snat_det_close_session_out_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa9997082"}
//        ],
//
type SnatDetCloseSessionOutReply struct {
	Retval int32
}

func (*SnatDetCloseSessionOutReply) GetMessageName() string {
	return "snat_det_close_session_out_reply"
}
func (*SnatDetCloseSessionOutReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatDetCloseSessionOutReply) GetCrcString() string {
	return "a9997082"
}
func NewSnatDetCloseSessionOutReply() api.Message {
	return &SnatDetCloseSessionOutReply{}
}

// SnatDetCloseSessionIn represents the VPP binary API message 'snat_det_close_session_in'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 374:
//
//        ["snat_det_close_session_in",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "in_addr", 16],
//            ["u16", "in_port"],
//            ["u8", "ext_addr", 16],
//            ["u16", "ext_port"],
//            {"crc" : "0x170dc301"}
//        ],
//
type SnatDetCloseSessionIn struct {
	IsIP4   uint8
	InAddr  []byte `struc:"[16]byte"`
	InPort  uint16
	ExtAddr []byte `struc:"[16]byte"`
	ExtPort uint16
}

func (*SnatDetCloseSessionIn) GetMessageName() string {
	return "snat_det_close_session_in"
}
func (*SnatDetCloseSessionIn) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetCloseSessionIn) GetCrcString() string {
	return "170dc301"
}
func NewSnatDetCloseSessionIn() api.Message {
	return &SnatDetCloseSessionIn{}
}

// SnatDetCloseSessionInReply represents the VPP binary API message 'snat_det_close_session_in_reply'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 385:
//
//        ["snat_det_close_session_in_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x1150e58b"}
//        ],
//
type SnatDetCloseSessionInReply struct {
	Retval int32
}

func (*SnatDetCloseSessionInReply) GetMessageName() string {
	return "snat_det_close_session_in_reply"
}
func (*SnatDetCloseSessionInReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SnatDetCloseSessionInReply) GetCrcString() string {
	return "1150e58b"
}
func NewSnatDetCloseSessionInReply() api.Message {
	return &SnatDetCloseSessionInReply{}
}

// SnatDetSessionDump represents the VPP binary API message 'snat_det_session_dump'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 391:
//
//        ["snat_det_session_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u8", "user_addr", 16],
//            {"crc" : "0xe412e169"}
//        ],
//
type SnatDetSessionDump struct {
	IsIP4    uint8
	UserAddr []byte `struc:"[16]byte"`
}

func (*SnatDetSessionDump) GetMessageName() string {
	return "snat_det_session_dump"
}
func (*SnatDetSessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetSessionDump) GetCrcString() string {
	return "e412e169"
}
func NewSnatDetSessionDump() api.Message {
	return &SnatDetSessionDump{}
}

// SnatDetSessionDetails represents the VPP binary API message 'snat_det_session_details'.
// Generated from '/usr/share/vpp/api//snat.api.json', line 399:
//
//        ["snat_det_session_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip4"],
//            ["u16", "in_port"],
//            ["u8", "ext_addr", 16],
//            ["u16", "ext_port"],
//            ["u16", "out_port"],
//            ["u8", "state"],
//            ["u32", "expire"],
//            {"crc" : "0x6e9cddb1"}
//        ]
//
type SnatDetSessionDetails struct {
	IsIP4   uint8
	InPort  uint16
	ExtAddr []byte `struc:"[16]byte"`
	ExtPort uint16
	OutPort uint16
	State   uint8
	Expire  uint32
}

func (*SnatDetSessionDetails) GetMessageName() string {
	return "snat_det_session_details"
}
func (*SnatDetSessionDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SnatDetSessionDetails) GetCrcString() string {
	return "6e9cddb1"
}
func NewSnatDetSessionDetails() api.Message {
	return &SnatDetSessionDetails{}
}
