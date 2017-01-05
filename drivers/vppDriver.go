package drivers

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/contiv/netplugin/core"
	"github.com/contiv/netplugin/netmaster/mastercfg"
	netutils "github.com/contiv/netplugin/utils/netutils"
	govpp "github.com/fdio-stack/go-vpp/srv"
	netlink "github.com/vishvananda/netlink"
)

type vppOper int

// VppDriverOperState carries operational state of the VppDriver.
type VppDriverOperState struct {
	core.CommonState
}

// VppDriver implements the Network and Endpoint Driver interfaces
// specific to VPP
type VppDriver struct {
	vppOper VppDriverOperState // Oper state of the driver
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
	log.Infof("Initializing vpp driver")
	govpp.VppConnect("vpp_contiv_client")
	return nil
}

// Deinit is not implemented.
func (d *VppDriver) Deinit() {
	govpp.VppDisconnect()
}

// CreateNetwork creates a network for a given ID.
func (d *VppDriver) CreateNetwork(id string) error {
	cfgNw := mastercfg.CfgNetworkState{}
	cfgNw.StateDriver = d.vppOper.StateDriver
	err := cfgNw.Read(id)
	if err != nil {
		log.Errorf("Failed to read net %s \n", cfgNw.ID)
		return err
	}

	log.Infof("Create net %+v \n", cfgNw)
	bdID := govpp.VppAddBridgeDomain(id)
	if bdID == 0 {
		log.Infof("Could not create bridge domain")
	} else {
		log.Infof("Bridge domain successfully created with id: %d", bdID)
	}
	return nil
}

// DeleteNetwork deletes a network provided by the ID.
func (d *VppDriver) DeleteNetwork(id, nwType, encap string, pktTag, extPktTag int, Gw string, tenant string) error {
	return core.Errorf("Not implemented")
}

// FetchNetwork retrieves a network's state given an ID.
func (d *VppDriver) FetchNetwork(id string) (core.State, error) {
	return nil, core.Errorf("Not implemented")
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

// addVppIntf creates a veth pair give a name.
func (d *VppDriver) addVppIntf(id string, intfName string) error {

	// Get VPP name
	vppIntfName, err := d.getVppIntName(intfName)
	if err != nil {
		log.Infof("Error generating Vpp veth pair name. Err: %v", err)
		return err
	}

	// Create a Veth pair
	err = netutils.CreateVethPairVpp(intfName, vppIntfName)
	if err != nil {
		log.Errorf("Error creating veth pairs. Err: %v", err)
		return err
	}
	log.Infof("Veth Pair Ready - int1:%s, int2:%s", intfName, vppIntfName)

	// Set host-side of the veth link pair state up
	vppLinkIntfName, err := netlink.LinkByName(vppIntfName)
	if err != nil {
		return err
	}
	netlink.LinkSetUp(vppLinkIntfName)
	log.Infof("Creating interface in VPP with name host-%s", vppIntfName)

	govpp.VppAddInterface(vppIntfName)
	// brecode - add return on error
	govpp.VppInterfaceAdminUp(vppIntfName)
	// brecode - add return on error
	govpp.VppInterfaceL2Bridge(id, vppIntfName)

	return nil
}

//UpdateEndpointGroup updates the endpoint with the new endpointgroup specification for the given ID.
func (d *VppDriver) UpdateEndpointGroup(id string) error {
	return core.Errorf("Not implemented")
}

// DeleteEndpoint destroys an endpoint for an ID.
func (d *VppDriver) DeleteEndpoint(id string) error {
	return core.Errorf("Not implemented")
}

// CreateHostAccPort creates a host access port
func (d *VppDriver) CreateHostAccPort(portName, globalIP, localIP string) error {
	return core.Errorf("Not implemented")
}

// DeleteHostAccPort creates a host access port
func (d *VppDriver) DeleteHostAccPort(portName string) error {
	return core.Errorf("Not implemented")
}

// FetchEndpoint retrieves an endpoint's state for a given ID
func (d *VppDriver) FetchEndpoint(id string) (core.State, error) {
	return nil, core.Errorf("Not implemented")
}

// AddPeerHost adds an peer host.
func (d *VppDriver) AddPeerHost(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

// DeletePeerHost removes a peer host.
func (d *VppDriver) DeletePeerHost(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

// AddMaster adds a master node.
func (d *VppDriver) AddMaster(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

// DeleteMaster removes a master node
func (d *VppDriver) DeleteMaster(node core.ServiceInfo) error {
	return core.Errorf("Not implemented")
}

//AddBgp adds bgp configs
func (d *VppDriver) AddBgp(id string) error {
	return core.Errorf("Not implemented")
}

//DeleteBgp deletes bgp configs
func (d *VppDriver) DeleteBgp(id string) error {
	return core.Errorf("Not implemented")
}

// AddSvcSpec invokes switch api
func (d *VppDriver) AddSvcSpec(svcName string, spec *core.ServiceSpec) error {
	return core.Errorf("Not implemented")
}

// DelSvcSpec invokes switch api
func (d *VppDriver) DelSvcSpec(svcName string, spec *core.ServiceSpec) error {
	return core.Errorf("Not implemented")
}

//DeleteServiceLB deletes service
func (d *VppDriver) DeleteServiceLB(servicename string, spec *core.ServiceSpec) error {
	return core.Errorf("Not implemented")
}

//SvcProviderUpdate hhhh
func (d *VppDriver) SvcProviderUpdate(servicename string, providers []string) {
	fmt.Println("Not implemented")
}

// GetEndpointStats returns all endpoint stats
func (d *VppDriver) GetEndpointStats() ([]byte, error) {
	return nil, core.Errorf("Not implemented")
}

// InspectState returns current state of the plugin
func (d *VppDriver) InspectState() ([]byte, error) {
	return nil, core.Errorf("Not implemented")
}

// InspectBgp returns current state of the plugin
func (d *VppDriver) InspectBgp() ([]byte, error) {
	return nil, core.Errorf("Not implemented")
}

func (d *VppDriver) getIntfName(cfgEp *mastercfg.CfgEndpointState) (string, error) {
	//Create a random interface name using Endpoint ID
	vethPrefix := "veth"
	vethID := cfgEp.EndpointID[:9]
	intfName := fmt.Sprint(vethPrefix + vethID)
	return intfName, nil

}

func (d *VppDriver) getVppIntName(intfName string) (string, error) {
	// Same interface format for vpp veth pair without the prefix
	vppIntfName := intfName[4:]
	return vppIntfName, nil
}
