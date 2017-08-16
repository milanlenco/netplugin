package vppd

import (
	"github.com/contiv/netplugin/netmaster/mastercfg"
	agent_core "github.com/ligato/cn-infra/core"
)

const (
	// StateOperPath is the path to the operations stored in state.
	vppOperPathPrefix = mastercfg.StateOperPath + "vpp-driver/"
	vppOperPath       = vppOperPathPrefix + "%s"
	vppDriverID agent_core.PluginName = "contiv-vpp-driver"
)
