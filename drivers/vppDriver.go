package drivers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/contiv/netplugin/core"
	"github.com/contiv/netplugin/netmaster/mastercfg"
	"github.com/contiv/netplugin/netplugin/nameserver"
	"github.com/contiv/netplugin/utils/netutils"
	"github.com/contiv/objdb"
	govpp "github.com/fdio-stack/govpp/srv"
	netlink "github.com/vishvananda/netlink"
)

type vppOper int

// VppDriverOperState carries operational state of the VppDriver.
type VppDriverOperState struct {
	core.CommonState
}

var defaultDbURL = "etcd://127.0.0.1:2379"

// VppDriver implements the Network and Endpoint Driver interfaces
// specific to VPP
type VppDriver struct {
	vppOper     VppDriverOperState              // Oper state of the driver
	localIP     string                          // Local IP address
	lock        sync.Mutex                      // lock for modifying shared state
	nameServer  *nameserver.NetpluginNameServer // nameServer
	objdbClient objdb.API                       // Objdb client
}

// Write the state
func (s *VppDriverOperState) Write() error {
	key := fmt.Sprintf(vppOperPath, s.ID)
	return s.StateDriver.WriteState(key, s, json.Marshal)
}

// Read the state given an ID.
func (s *VppDriverOperState) Read(id string) error {
	key := fmt.Sprintf(vppOperPath, id)
	return s.StateDriver.ReadState(key, s, json.Unmarshal)
}

// ReadAll reads all the state
func (s *VppDriverOperState) ReadAll() ([]core.State, error) {
	return s.StateDriver.ReadAllState(vppOperPathPrefix, s, json.Unmarshal)
}

// Clear removes the state.
func (s *VppDriverOperState) Clear() error {
	key := fmt.Sprintf(vppOperPath, s.ID)
	return s.StateDriver.ClearState(key)
}

// Init initializes the VPP driver
func (d *VppDriver) Init(info *core.InstanceInfo) error {
	if info == nil || info.StateDriver == nil {
		return core.Errorf("Invalid arguments. instance-info: %+v", info)
	}
	d.vppOper.StateDriver = info.StateDriver
	d.localIP = info.VtepIP
	err := d.vppOper.Read(info.HostLabel)
	if core.ErrIfKeyExists(err) != nil {
		log.Errorf("Failed to read driver oper state for key %q. Error: %s",
			info.HostLabel, err)
		return err
	} else if err != nil {
		// create the oper state as it is first time start up
		d.vppOper.ID = info.HostLabel

		// write the oper
		err = d.vppOper.Write()
		if err != nil {
			return err
		}
	}

	d.objdbClient, err = objdb.NewClient(defaultDbURL)
	if err != nil {
		log.Fatalf("Error connecting to state store: %s. Err: %v", defaultDbURL, err)
	}

	log.Infof("Initializing vpp driver")
	govpp.VppConnect()
	return nil
}

// Deinit is not implemented.
func (d *VppDriver) Deinit() {
}

// CreateNetwork creates a bridge domain network for a given ID in VPP
func (d *VppDriver) CreateNetwork(id string) error {
	cfgNw := mastercfg.CfgNetworkState{}
	cfgNw.StateDriver = d.vppOper.StateDriver
	err := cfgNw.Read(id)
	if err != nil {
		log.Errorf("Failed to read net %s \n", cfgNw.ID)
		return err
	}
	isAdd := 1
	log.Infof("Create net %+v \n", cfgNw)
	err = govpp.VppAddDelBridgeDomain(id, uint32(cfgNw.PktTag), uint8(isAdd))
	if err != nil {
		return err
	}
	// localIP := d.localIP
	// clusterIP, err := d.objdbClient.GetService("netplugin.vtep")
	// if err != nil {
	// 	log.Fatalf("Error connecting to state store: %s. Err: %v", defaultDbURL, err)
	// }

	// for _, nodeIP := range clusterIP {
	// 	if nodeIP.HostAddr != localIP {
	// 		tunnelIfIndex, err := govpp.VppVxlanAddDelTunnel(uint8(isAdd), 0, net.ParseIP(localIP).To4(), net.ParseIP(nodeIP.HostAddr).To4(), uint32(cfgNw.ExtPktTag))
	// 		if err != nil {
	// 			log.Infof("Could not create vxlan tunnel")
	// 			return err
	// 		}
	// 		err = govpp.VppSetInterfaceL2Bridge(id, 1, "", tunnelIfIndex, uint8(1))
	// 		if err != nil {
	// 			log.Errorf("Error adding interface to bridge domain, Err: %v", err)
	// 			return err
	// 		}
	// 	}
	// }

	var localIP string
	if d.localIP == "192.168.2.10" {
		localIP = "192.168.10.10"
	} else {
		localIP = "192.168.10.11"
	}
	clusterIP := []string{"192.168.10.10", "192.168.10.11"}

	for _, nodeIP := range clusterIP {
		if nodeIP != localIP {
			tunnelIfIndex, err := govpp.VppVxlanAddDelTunnel(uint8(isAdd), 0, net.ParseIP(localIP).To4(), net.ParseIP(nodeIP).To4(), uint32(cfgNw.ExtPktTag))
			if err != nil {
				log.Infof("Could not create vxlan tunnel")
				return err
			}
			err = govpp.VppSetInterfaceL2Bridge(id, 1, "", tunnelIfIndex, uint8(1))
			if err != nil {
				log.Errorf("Error adding interface to bridge domain, Err: %v", err)
				return err
			}
		}
	}

	return nil
}

// DeleteNetwork deletes a network for a given ID from VPP
func (d *VppDriver) DeleteNetwork(id string, nwType, encap string, pktTag, extPktTag int, gateway string, tenant string) error {
	isAdd := 0
	err := govpp.VppAddDelBridgeDomain(id, uint32(pktTag), uint8(isAdd))
	if err != nil {
		return err
	}
	return nil
}

// CreateEndpoint creates an endpoint for a given ID.
func (d *VppDriver) CreateEndpoint(id string) error {
	log.Infof("Create endpoint with id: %s", id)
	var (
		err      error
		intfName string
	)

	cfgEp := &mastercfg.CfgEndpointState{}
	cfgEp.StateDriver = d.vppOper.StateDriver
	err = cfgEp.Read(id)
	if err != nil {
		log.Errorf("Unable to get EpState %s. Err: %v", cfgEp.NetID, err)
		return err
	}

	cfgNw := mastercfg.CfgNetworkState{}
	cfgNw.StateDriver = d.vppOper.StateDriver
	err = cfgNw.Read(cfgEp.NetID)
	if err != nil {
		log.Errorf("Unable to get network %s. Err: %v", cfgNw.NetworkName, err)
		return err
	}

	cfgEpGroup := &mastercfg.EndpointGroupState{}
	cfgEpGroup.StateDriver = d.vppOper.StateDriver

	operEp := &VppOperEndpointState{}
	operEp.StateDriver = d.vppOper.StateDriver

	intfName, err = d.getIntfName(cfgEp)
	if err != nil {
		log.Errorf("Error generating intfName %s. Err: %v", intfName, err)
		return err
	}

	// Ask VPP to create the interface. Part is to create a veth pair.
	networkID := cfgNw.CommonState.ID
	err = d.addVppIntf(networkID, intfName)
	if err != nil {
		log.Errorf("Error creating vpp interface %s. Err: %v", intfName, err)
		return err
	}

	// Save the oper state
	operEp = &VppOperEndpointState{
		NetID:       cfgEp.NetID,
		EndpointID:  cfgEp.EndpointID,
		ServiceName: cfgEp.ServiceName,
		IPAddress:   cfgEp.IPAddress,
		MacAddress:  cfgEp.MacAddress,
		IntfName:    cfgEp.IntfName,
		HomingHost:  cfgEp.HomingHost,
		PortName:    intfName,
	}

	operEp.StateDriver = d.vppOper.StateDriver
	operEp.ID = id
	err = operEp.Write()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			operEp.Clear()
		}
	}()

	return nil
}

//UpdateEndpointGroup is not implemented.
func (d *VppDriver) UpdateEndpointGroup(id string) error {
	return core.Errorf("Not implemented")
}

// DeleteEndpoint is not implemented.
func (d *VppDriver) DeleteEndpoint(id string) (err error) {
	return core.Errorf("Not implemented")
}

// CreateHostAccPort is not implemented.
func (d *VppDriver) CreateHostAccPort(id, a string, nw int) (string, error) {
	return "", core.Errorf("Not implemented")
}

// DeleteHostAccPort is not implemented.
func (d *VppDriver) DeleteHostAccPort(id string) (err error) {
	return core.Errorf("Not implemented")
}

// AddPeerHost is not implemented.
func (d *VppDriver) AddPeerHost(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

// DeletePeerHost is not implemented.
func (d *VppDriver) DeletePeerHost(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

// AddMaster is not implemented
func (d *VppDriver) AddMaster(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

// DeleteMaster is not implemented
func (d *VppDriver) DeleteMaster(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

// AddBgp is not implemented.
func (d *VppDriver) AddBgp(id string) (err error) {
	return core.Errorf("Not implemented")
}

// DeleteBgp is not implemented.
func (d *VppDriver) DeleteBgp(id string) (err error) {
	return core.Errorf("Not implemented")
}

// AddSvcSpec is not implemented.
func (d *VppDriver) AddSvcSpec(svcName string, spec *core.ServiceSpec) error {
	return core.Errorf("Not implemented")
}

// DelSvcSpec is not implemented.
func (d *VppDriver) DelSvcSpec(svcName string, spec *core.ServiceSpec) error {
	return core.Errorf("Not implemented")
}

// SvcProviderUpdate is not implemented.
func (d *VppDriver) SvcProviderUpdate(svcName string, providers []string) {
}

// GetEndpointStats is not implemented
func (d *VppDriver) GetEndpointStats() ([]byte, error) {
	return []byte{}, core.Errorf("Not implemented")
}

// InspectState is not implemented
func (d *VppDriver) InspectState() ([]byte, error) {
	return []byte{}, core.Errorf("Not implemented")
}

// InspectBgp is not implemented
func (d *VppDriver) InspectBgp() ([]byte, error) {
	return []byte{}, core.Errorf("Not implemented")
}

// GlobalConfigUpdate is not implemented
func (d *VppDriver) GlobalConfigUpdate(inst core.InstanceInfo) error {
	return core.Errorf("Not implemented")
}

// InspectNameserver returns nameserver state as json string
func (d *VppDriver) InspectNameserver() ([]byte, error) {
	return []byte{}, core.Errorf("Not implemented")
}

// CreateRemoteEndpoint is not implemented.
func (d *VppDriver) CreateRemoteEndpoint(id string) error {
	// var err error
	// operEp := &VppOperEndpointState{}
	// operEp.StateDriver = d.vppOper.StateDriver
	// err = operEp.Read(id)
	// if err != nil {
	// 	return err
	// }

	// cfgNw := mastercfg.CfgNetworkState{}
	// cfgNw.StateDriver = d.vppOper.StateDriver
	// err = cfgNw.Read(operEp.NetID)
	// if err != nil {
	// 	log.Errorf("Unable to get network %s. Err: %v", operEp.NetID, err)
	// 	return err
	// }
	// isAdd := uint8(1)
	// isIPv6 := uint8(0)
	// srcAddr := []byte(operEp.VtepIP)
	// dstAddr := []byte(d.localIP)
	// vni := uint32(cfgNw.ExtPktTag)
	// shg := uint8(1)
	// tunnelIfIndex, err := govpp.VppVxlanAddDelTunnel(isAdd, isIPv6, srcAddr, dstAddr, vni)
	// if err != nil {
	// 	return err
	// }
	// err = govpp.VppSetInterfaceL2Bridge(id, string(tunnelIfIndex), shg)
	// if err != nil {
	// 	return err
	// }
	return core.Errorf("Not implemented")
}

// DeleteRemoteEndpoint is not implemented.
func (d *VppDriver) DeleteRemoteEndpoint(id string) (err error) {
	return core.Errorf("Not implemented")
}

// AddPolicyRule is not implemented
func (d *VppDriver) AddPolicyRule(id string) error {
	return core.Errorf("Not implemented")
}

// DelPolicyRule is not implemented
func (d *VppDriver) DelPolicyRule(id string) error {
	return core.Errorf("Not implemented")
}

// getVppIntfName returns VPP Interface name
func getVppIntfName(intfName string) (string, error) {
	// Same interface format for vpp veth pair without the prefix
	vppIntfName := intfName[5:]
	if vppIntfName == "" {
		err := errors.New("Could not generate name for VPP interface")
		return "", err
	}
	return vppIntfName, nil
}

// getIntfName generates an interface name from cfgEndpointState
func (d *VppDriver) getIntfName(cfgEp *mastercfg.CfgEndpointState) (string, error) {
	//Create a random interface name using Endpoint ID
	vethPrefix := "veth_"
	vethID := cfgEp.EndpointID[:9]
	if vethID == "" {
		err := errors.New("Error getting ID from cfgEp ID")
		return "", err
	}
	intfName := fmt.Sprint(vethPrefix + vethID)
	return intfName, nil
}

// addVppIntf creates a veth pair give a name and attaches one end to VPP.
func (d *VppDriver) addVppIntf(id string, intfName string) error {
	// Get VPP name
	vppIntfName, err := getVppIntfName(intfName)
	if err != nil {
		log.Errorf("Error generating vpp veth pair name. Err: %v", err)
		return err
	}
	// Create a Veth pair
	err = netutils.CreateVethPairVpp(intfName, vppIntfName)
	if err != nil {
		log.Errorf("Error creating the veth pair. Err: %v", err)
		return err
	}
	// Set host-side link for the veth pair
	vppLinkIntfName, err := netlink.LinkByName(vppIntfName)
	if err != nil {
		log.Errorf("Error setting host-side link for the veth pair, Err: %v", err)
		return err
	}
	err = netlink.LinkSetUp(vppLinkIntfName)
	if err != nil {
		log.Errorf("Error setting state up for veth pair, Err: %v", err)
		return err
	}

	err = govpp.VppAddInterface(vppIntfName)
	if err != nil {
		log.Errorf("Error creating the vpp-side interface, Err: %v", err)
		return err
	}
	err = govpp.VppInterfaceAdminUp(vppIntfName)
	if err != nil {
		log.Errorf("Error setting the vpp-side interface state to up, Err: %v", err)
		return err
	}
	err = govpp.VppSetInterfaceL2Bridge(id, 0, vppIntfName, uint32(0), uint8(0))
	if err != nil {
		log.Errorf("Error adding interface to bridge domain, Err: %v", err)
		return err
	}
	return nil
}
