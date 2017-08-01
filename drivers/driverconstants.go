package drivers

import "github.com/contiv/netplugin/netmaster/mastercfg"

const (
	// StateOperPath is the path to the operations stored in state.
	networkOperPathPrefix  = mastercfg.StateOperPath + "nets/"
	endpointOperPathPrefix = mastercfg.StateOperPath + "eps/"
	networkOperPath        = networkOperPathPrefix + "%s"
	endpointOperPath       = endpointOperPathPrefix + "%s"
)
