// Package dhcp represents the VPP binary API of the 'dhcp' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api//dhcp.api.json' on Fri, 28 Apr 2017 17:43:59 UTC.
package dhcp

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xa0e9d3e4

// DhcpServer represents the VPP binary API data type 'dhcp_server'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 3:
//
//        ["dhcp_server",
//            ["u32", "server_vrf_id"],
//            ["u8", "dhcp_server", 16],
//            {"crc" : "0xe6266622"}
//        ]
//
type DhcpServer struct {
	ServerVrfID uint32
	DhcpServer  []byte `struc:"[16]byte"`
}

func (*DhcpServer) GetTypeName() string {
	return "dhcp_server"
}
func (*DhcpServer) GetCrcString() string {
	return "e6266622"
}

// DhcpProxyConfig represents the VPP binary API message 'dhcp_proxy_config'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 10:
//
//        ["dhcp_proxy_config",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "rx_vrf_id"],
//            ["u32", "server_vrf_id"],
//            ["u8", "is_ipv6"],
//            ["u8", "is_add"],
//            ["u8", "dhcp_server", 16],
//            ["u8", "dhcp_src_address", 16],
//            {"crc" : "0x3b4f2bc8"}
//        ],
//
type DhcpProxyConfig struct {
	RxVrfID        uint32
	ServerVrfID    uint32
	IsIpv6         uint8
	IsAdd          uint8
	DhcpServer     []byte `struc:"[16]byte"`
	DhcpSrcAddress []byte `struc:"[16]byte"`
}

func (*DhcpProxyConfig) GetMessageName() string {
	return "dhcp_proxy_config"
}
func (*DhcpProxyConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpProxyConfig) GetCrcString() string {
	return "3b4f2bc8"
}
func NewDhcpProxyConfig() api.Message {
	return &DhcpProxyConfig{}
}

// DhcpProxyConfigReply represents the VPP binary API message 'dhcp_proxy_config_reply'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 22:
//
//        ["dhcp_proxy_config_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xfe63196f"}
//        ],
//
type DhcpProxyConfigReply struct {
	Retval int32
}

func (*DhcpProxyConfigReply) GetMessageName() string {
	return "dhcp_proxy_config_reply"
}
func (*DhcpProxyConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpProxyConfigReply) GetCrcString() string {
	return "fe63196f"
}
func NewDhcpProxyConfigReply() api.Message {
	return &DhcpProxyConfigReply{}
}

// DhcpProxySetVss represents the VPP binary API message 'dhcp_proxy_set_vss'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 28:
//
//        ["dhcp_proxy_set_vss",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "tbl_id"],
//            ["u32", "oui"],
//            ["u32", "fib_id"],
//            ["u8", "is_ipv6"],
//            ["u8", "is_add"],
//            {"crc" : "0xbe54d194"}
//        ],
//
type DhcpProxySetVss struct {
	TblID  uint32
	Oui    uint32
	FibID  uint32
	IsIpv6 uint8
	IsAdd  uint8
}

func (*DhcpProxySetVss) GetMessageName() string {
	return "dhcp_proxy_set_vss"
}
func (*DhcpProxySetVss) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpProxySetVss) GetCrcString() string {
	return "be54d194"
}
func NewDhcpProxySetVss() api.Message {
	return &DhcpProxySetVss{}
}

// DhcpProxySetVssReply represents the VPP binary API message 'dhcp_proxy_set_vss_reply'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 39:
//
//        ["dhcp_proxy_set_vss_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5bb4e754"}
//        ],
//
type DhcpProxySetVssReply struct {
	Retval int32
}

func (*DhcpProxySetVssReply) GetMessageName() string {
	return "dhcp_proxy_set_vss_reply"
}
func (*DhcpProxySetVssReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpProxySetVssReply) GetCrcString() string {
	return "5bb4e754"
}
func NewDhcpProxySetVssReply() api.Message {
	return &DhcpProxySetVssReply{}
}

// DhcpClientConfig represents the VPP binary API message 'dhcp_client_config'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 45:
//
//        ["dhcp_client_config",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "hostname", 64],
//            ["u8", "is_add"],
//            ["u8", "want_dhcp_event"],
//            ["u32", "pid"],
//            {"crc" : "0x87920bb6"}
//        ],
//
type DhcpClientConfig struct {
	SwIfIndex     uint32
	Hostname      []byte `struc:"[64]byte"`
	IsAdd         uint8
	WantDhcpEvent uint8
	Pid           uint32
}

func (*DhcpClientConfig) GetMessageName() string {
	return "dhcp_client_config"
}
func (*DhcpClientConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpClientConfig) GetCrcString() string {
	return "87920bb6"
}
func NewDhcpClientConfig() api.Message {
	return &DhcpClientConfig{}
}

// DhcpClientConfigReply represents the VPP binary API message 'dhcp_client_config_reply'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 56:
//
//        ["dhcp_client_config_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd947f4c8"}
//        ],
//
type DhcpClientConfigReply struct {
	Retval int32
}

func (*DhcpClientConfigReply) GetMessageName() string {
	return "dhcp_client_config_reply"
}
func (*DhcpClientConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpClientConfigReply) GetCrcString() string {
	return "d947f4c8"
}
func NewDhcpClientConfigReply() api.Message {
	return &DhcpClientConfigReply{}
}

// DhcpComplEvent represents the VPP binary API message 'dhcp_compl_event'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 62:
//
//        ["dhcp_compl_event",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "pid"],
//            ["u8", "hostname", 64],
//            ["u8", "is_ipv6"],
//            ["u8", "host_address", 16],
//            ["u8", "router_address", 16],
//            ["u8", "host_mac", 6],
//            {"crc" : "0x5218db55"}
//        ],
//
type DhcpComplEvent struct {
	Pid           uint32
	Hostname      []byte `struc:"[64]byte"`
	IsIpv6        uint8
	HostAddress   []byte `struc:"[16]byte"`
	RouterAddress []byte `struc:"[16]byte"`
	HostMac       []byte `struc:"[6]byte"`
}

func (*DhcpComplEvent) GetMessageName() string {
	return "dhcp_compl_event"
}
func (*DhcpComplEvent) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpComplEvent) GetCrcString() string {
	return "5218db55"
}
func NewDhcpComplEvent() api.Message {
	return &DhcpComplEvent{}
}

// DhcpProxyDump represents the VPP binary API message 'dhcp_proxy_dump'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 73:
//
//        ["dhcp_proxy_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ip6"],
//            {"crc" : "0x175bc073"}
//        ],
//
type DhcpProxyDump struct {
	IsIP6 uint8
}

func (*DhcpProxyDump) GetMessageName() string {
	return "dhcp_proxy_dump"
}
func (*DhcpProxyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpProxyDump) GetCrcString() string {
	return "175bc073"
}
func NewDhcpProxyDump() api.Message {
	return &DhcpProxyDump{}
}

// DhcpProxyDetails represents the VPP binary API message 'dhcp_proxy_details'.
// Generated from '/usr/share/vpp/api//dhcp.api.json', line 80:
//
//        ["dhcp_proxy_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "rx_vrf_id"],
//            ["u32", "vss_oui"],
//            ["u32", "vss_fib_id"],
//            ["u8", "is_ipv6"],
//            ["u8", "dhcp_src_address", 16],
//            ["u8", "count"],
//            ["vl_api_dhcp_server_t", "servers", 0, "count"],
//            {"crc" : "0x9727dbdd"}
//        ]
//
type DhcpProxyDetails struct {
	RxVrfID        uint32
	VssOui         uint32
	VssFibID       uint32
	IsIpv6         uint8
	DhcpSrcAddress []byte `struc:"[16]byte"`
	Count          uint8  `struc:"sizeof=Servers"`
	Servers        []DhcpServer
}

func (*DhcpProxyDetails) GetMessageName() string {
	return "dhcp_proxy_details"
}
func (*DhcpProxyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpProxyDetails) GetCrcString() string {
	return "9727dbdd"
}
func NewDhcpProxyDetails() api.Message {
	return &DhcpProxyDetails{}
}
