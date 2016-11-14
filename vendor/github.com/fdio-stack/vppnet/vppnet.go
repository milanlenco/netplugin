package vppnet

// This file contains the ofnet master implementation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/contiv/ofnet/rpcHub"

	log "github.com/Sirupsen/logrus"
)

// Ofnet master state
type VppNetMaster struct {
	myPort      uint16       // port where we are listening
	rpcServer   *rpc.Server  // json-rpc server
	rpcListener net.Listener // Listener
}

func NewVppMaster(portNo uint16) *VppNetMaster {
	// Create the master
	master := new(VppNetMaster)

	// Init params
	master.myPort = portNo

	// Create a new RPC server
	master.rpcServer, master.rpcListener = rpcHub.NewRpcServer(portNo)

	// Register RPC handler
	err := master.rpcServer.Register(master)
	if err != nil {
		log.Fatalf("Error Registering RPC callbacks. Err: %v", err)
		return nil
	}

	return master
}

// Delete closes rpc listener
func (self *VppNetMaster) Delete() error {
	self.rpcListener.Close()
	time.Sleep(100 * time.Millisecond)

	return nil
}
