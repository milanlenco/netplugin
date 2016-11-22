package drivers

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/contiv/netplugin/core"
	"github.com/contiv/netplugin/netmaster/mastercfg"
	netutils "github.com/contiv/netplugin/utils/netutils"
	govpp "github.com/fdio-stack/go-vpp/srv"
	//"github.com/go-shortid"
	// "github.com/vishvananda/netlink"
)

var uid = 1

// // VppDriverConfig represents the configuration of the vppdriver
// type VppDriverConfig struct {
// 	VppClientName string
// }

// VppDriverOperState carries operational state of the VppDriver.
type VppDriverOperState struct {
	core.CommonState
}

// VppDriver implements the Network and Endpoint Driver interfaces
// specific to VPP
type VppDriver struct {
	oper VppDriverOperState // Oper state of the driver
}

// Init is not
func (d *VppDriver) Init(info *core.InstanceInfo) error {
	fmt.Println("Called Vpp: Init")
	govpp.VppConnect("vpp_contiv_client")
	return nil
}

// Deinit is not implemented.
func (d *VppDriver) Deinit() {
	govpp.VppDisconnect()
}

func (d *VppDriver) getIntfName() (string, error) {
	// // Create a random interface name
	// if sid, err := shortid.New(1, shortid.DefaultABC, 2342); err != nil {
	// 	log.Infof("Could not generate interface name")
	// }
	uid++

	intfName := fmt.Sprintf("veth_%d", uid)
	return intfName, nil
}

func (d *VppDriver) getVppIntName(intfName string) (string, error) {
	// Same interface format for vpp veth pair
	vppIntfName := intfName[5:]
	return vppIntfName, nil
}

// CreateNetwork creates a network for a given ID.
func (d *VppDriver) CreateNetwork(id string) error {
	fmt.Println(id)
	cfgNw := mastercfg.CfgNetworkState{}
	cfgNw.StateDriver = d.oper.StateDriver
	log.Infof("create net %+v \n", cfgNw)
	bdID := govpp.VppBridgeDomain(id)
	if bdID == 0 {
		log.Infof("Could not create bridge domain")
	} else {
		fmt.Println("Bridge domain successfully created with id:")
		fmt.Println(bdID)
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
	cfgEp.StateDriver = d.oper.StateDriver
	// brecode - read endpoint
	cfgNw := mastercfg.CfgNetworkState{}
	cfgNw.StateDriver = d.oper.StateDriver
	//brecode - read mastercgf

	cfgEpGroup := &mastercfg.EndpointGroupState{}
	cfgEpGroup.StateDriver = d.oper.StateDriver

	// Get the interface name to use
	intfName, err = d.getIntfName()
	if err != nil {
		return err
	}

	// Ask VPP to create the Port. Part is to create a veth pair.
	err = d.CreateVppIntf(intfName)
	if err != nil {
		log.Errorf("Error creating port %s. Err: %v", intfName, err)
		return err
	}

	// Handle the Operation State of Endpoints
	// operEp := &VppOperEndpointState{}
	// operEp.StateDriver = d.oper.StateDriver
	// // Save the oper state
	// operEp = &VppOperEndpointState{
	// 	NetID:       cfgEp.NetID,
	// 	AttachUUID:  cfgEp.AttachUUID,
	// 	ContName:    cfgEp.ContName,
	// 	ServiceName: cfgEp.ServiceName,
	// 	IPAddress:   cfgEp.IPAddress,
	// 	MacAddress:  cfgEp.MacAddress,
	// 	IntfName:    cfgEp.IntfName,
	// 	PortName:    intfName,
	// 	HomingHost:  cfgEp.HomingHost,
	// 	VtepIP:      cfgEp.VtepIP}
	// operEp.StateDriver = d.oper.StateDriver
	// operEp.ID = id
	// err = operEp.Write()
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	if err != nil {
	// 		operEp.Clear()
	// 	}
	// }()

	return nil

}

// CreatePort creates an interface in VPP for a given veth pair name.
func (d *VppDriver) CreateVppIntf(intfName string) error {

	// Get VPP port name
	vppIntfName, err := d.getVppIntName(intfName)
	if err != nil {
		log.Errorf("Error generating Vpp veth pair name. Err: %v", err)
		return err
	}

	// Create a Veth pair
	err := netutils.CreateVethPairVpp(intfName, vppIntfName)
	if err != nil {
		log.Errorf("Error creating veth pairs. Err: %v", err)
		return err
	}

	// Add VPP interface
	govpp.VppAddInterface(vppIntfName)

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
