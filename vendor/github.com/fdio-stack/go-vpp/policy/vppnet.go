/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ofnet

// This package implements openflow network manager

import (
	"net"
	"time"
)

// VppnetEndpoint has info about an endpoint
type VppnetEndpoint struct {
	EndpointID        string    // Unique identifier for the endpoint
	EndpointType      string    // Type of the endpoint "internal", "external" or "externalRoute"
	EndpointGroup     int       // Endpoint group identifier for policies.
	IpAddr            net.IP    // IP address of the end point
	IpMask            net.IP    // IP mask for the end point
	Ipv6Addr          net.IP    // IPv6 address of the end point
	Ipv6Mask          net.IP    // IPv6 mask for the end point
	Vrf               string    // IP address namespace
	MacAddrStr        string    // Mac address of the end point(in string format)
	Vlan              uint16    // Vlan Id for the endpoint
	Vni               uint32    // Vxlan VNI
	EndpointGroupVlan uint16    // EnpointGroup Vlan, needed in non-Standalone mode of netplugin
	OriginatorIp      net.IP    // Originating switch
	PortNo            uint32    `json:"-"` // Port number on originating switch
	Dscp              int       `json:"-"` // DSCP value for the endpoint
	Timestamp         time.Time // Timestamp of the last event
}

// VppnetPolicyRule has security rule to be installed
type VppnetPolicyRule struct {
	RuleId           string // Unique identifier for the rule
	Priority         int    // Priority for the rule (1..100. 100 is highest)
	SrcEndpointGroup int    // Source endpoint group
	DstEndpointGroup int    // Destination endpoint group
	SrcIpAddr        string // source IP addrss and mask
	DstIpAddr        string // Destination IP address and mask
	IpProtocol       uint8  // IP protocol number
	SrcPort          uint16 // Source port
	DstPort          uint16 // destination port
	TcpFlags         string // TCP flags to match: syn || syn,ack || ack || syn,!ack || !syn,ack;
	Action           string // rule action: 'accept' or 'deny'
}
