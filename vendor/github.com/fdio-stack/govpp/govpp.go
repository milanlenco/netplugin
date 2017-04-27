package govpp

import (
	"github.com/fdio-stack/govpp/adapter"
	"github.com/fdio-stack/govpp/adapter/vppapiclient"
	"github.com/fdio-stack/govpp/core"
)

var vppAdapter adapter.VppAdapter // VPP Adapter that will be used in the subsequent Connect calls

// Connect connects the govpp core to VPP either using the default VPP Adapter, or using the adapter previously
// set by SetAdapter (useful mostly just for unit/integration tests with mocked VPP adapter).
func Connect() (*core.Connection, error) {
	if vppAdapter == nil {
		vppAdapter = vppapiclient.NewVppAdapter()
	}
	return core.Connect(vppAdapter)
}

// SetAdapter sets the adapter that will be used for connections to VPP in the subsequent `Connect` calls.
func SetAdapter(ad adapter.VppAdapter) {
	vppAdapter = ad
}
