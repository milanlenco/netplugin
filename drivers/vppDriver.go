package drivers

import (
	"fmt"
	"github.com/contiv/netplugin/core"
	govpp "github.com/fdio-stack/go-vpp/srv"
)

// // VppNetEpDriverConfig represents the configuration of the vppdriver
// type VppNetEpDriverConfig struct {
// 	VppClientName string
// }

// VppDriverOperState carries operational state of the VppDriver.
type VppDriverOperState struct {
	core.CommonState
}

// VppDriver implements the Network and Endpoint Driver interfaces
// specific to VPP
type VppDriver struct {
	oper     VppDriverOperState // Oper state of the driver
	clientID string
}

// // VppDriverConfig defines the configuration required to initialize the
// // VppDriver.
// type VppDriverConfig struct {
// 	Vpp struct {
// 		bollocks string
// 	}
// }

// // VppNetEpDriver implements core.NetworkDriver interface
// // for use with unit-tests
// type VppNetEpDriver struct {
// }

// Init is not implemented.
func Init(info *core.InstanceInfo) error {
	fmt.Println("Called Vpp: Init")
	govpp.VppConnect("vpp_contiv_client")
	return nil
}

// // Deinit is not implemented.
// func (d *VppNetEpDriver) Deinit() {
// }

// // CreateNetwork will leverage an external library .. bo boom cha.
// func (d *VppNetEpDriver) CreateNetwork(id string) error {
// 	vppc.CreateNetwork(id)
// 	return core.Errorf("VppNet: CreateNetwork Not implemented")
// }

// // DeleteNetwork is not implemented.
// func (d *VppNetEpDriver) DeleteNetwork(id, nwType, encap string, pktTag, extPktTag int, gateway string, tenant string) error {
// 	return core.Errorf("VppNet: DeleteNetwork Not implemented")
// }

// // CreateEndpoint is not implemented.
// func (d *VppNetEpDriver) CreateEndpoint(id string) error {
// 	return core.Errorf("VppNet: CreateEndpoint Not implemented")
// }

// // DeleteEndpoint is not implemented.
// func (d *VppNetEpDriver) DeleteEndpoint(id string) (err error) {
// 	return core.Errorf("VppNet DeleteEndpoint Not implemented")
// }
