package adapter

// VppAdapter provides connection to VPP. It is responsible for sending and receiving of binary-encoded messages to/from VPP.
type VppAdapter interface {
	// Connect connects the process to VPP.
	Connect() error

	// Disconnect disconnects the process from VPP.
	Disconnect()

	// GetMsgID returns a runtime message ID for the given message name and CRC.
	GetMsgID(msgName string, msgCrc string) (uint16, error)

	// SendMsg sends a binary-encoded message to VPP.
	SendMsg(clientID uint32, data []byte) error

	// SetMsgCallback sets a callback function that will be called by the adapter whenever a message comes from VPP.
	SetMsgCallback(func(context uint32, msgId uint16, data []byte))
}
