// +build windows darwin

/*
	This is just an empty adapter that does nothing. It builds only on Windows and OSX, where the real
	VPP API client adapter does not build. Its sole purpose is to make the compiler happy on Windows and OSX.
*/

package vppapiclient

import (
	"github.com/fdio-stack/govpp/adapter"
)

type vppAPIClientAdapter struct{}

func NewVppAdapter() adapter.VppAdapter {
	return &vppAPIClientAdapter{}
}

func (a *vppAPIClientAdapter) Connect() error {
	return nil
}

func (a *vppAPIClientAdapter) Disconnect() {
	// no op
}

func (a *vppAPIClientAdapter) GetMsgID(msgName string, msgCrc string) (uint16, error) {
	return 0, nil
}

func (a *vppAPIClientAdapter) SendMsg(clientID uint32, data []byte) error {
	return nil
}

func (a *vppAPIClientAdapter) SetMsgCallback(cb func(context uint32, msgID uint16, data []byte)) {
	// no op
}
