// Package mock is an alternative VPP adapter aimed for unit/integration testing where the actual communication
// with VPP is not demanded.
package mock

import (
	"bytes"

	"github.com/lunixbochs/struc"

	"sync"

	"reflect"

	"github.com/fdio-stack/govpp/adapter"
	"github.com/fdio-stack/govpp/api"
	"github.com/fdio-stack/govpp/binapi_generator/reflect"
)

// VppAdapter represents a mock VPP adapter that can be used for unit/integration testing instead of the vppapiclient adapter.
type VppAdapter struct {
	callback func(context uint32, msgId uint16, data []byte)

	msgNameToIds *map[string]uint16
	msgIdsToName *map[uint16]string
	msgIdSeq     uint16
	binApiTypes  map[string]reflect.Type
	//TODO lock
}

// replyHeader represents a common header of each VPP request message.
type requestHeader struct {
	VlMsgID     uint16
	ClientIndex uint32
	Context     uint32
}

// replyHeader represents a common header of each VPP reply message.
type replyHeader struct {
	VlMsgID uint16
	Context uint32
}

// replyHeader represents a common header of each VPP reply message.
type vppOtherHeader struct {
	VlMsgID uint16
}

// defaultReply is a default reply message that mock adapter returns for a request.
type defaultReply struct {
	Retval int32
}

// MessageDTO is a structure used for propageating informations to ReplyHandlers
type MessageDTO struct {
	MsgID    uint16
	MsgName  string
	ClientID uint32
	Data     []byte
}

// ReplyHandler is a type that allows to extend the behaviour of VPP mock.
// Return value prepared is used to signalize that mock reply is calculated.
type ReplyHandler func(dto MessageDTO) (reply MessageDTO, prepared bool)

const (
	//defaultMsgID      = 1 // default message ID to be returned from GetMsgId
	defaultReplyMsgID = 2 // default message ID for the reply to be sent back via callback
)

var replies []api.Message        // FIFO queue of messages
var replyHandlers []ReplyHandler // callbacks that are able to calculate mock responses
var repliesLock sync.Mutex       // mutex for the queue
var mode = 0

const useRepliesQueue = 1  // use replies in the queue instead of the default one
const useReplyHandlers = 2 //use ReplyHandler

// NewVppAdapter returns a new mock adapter.
func NewVppAdapter() adapter.VppAdapter {
	return &VppAdapter{}
}

// Connect emulates connecting the process to VPP.
func (a *VppAdapter) Connect() error {
	return nil
}

// Disconnect emulates disconnecting the process from VPP.
func (a *VppAdapter) Disconnect() {
	// no op
}

func (a *VppAdapter) GetMsgNameByID(msgId uint16, msgCrc string) (string, bool) {
	a.initMaps()

	msgName, found := (*a.msgIdsToName)[msgId]

	return msgName, found
}

func (a *VppAdapter) RegisterBinApiTypes(binApiTypes map[string]reflect.Type) {
	a.initMaps()
	for k, v := range binApiTypes {
		a.binApiTypes[k] = v
	}
}

func (a *VppAdapter) ReplyFor(request reflect.Type) (reflect.Type, bool) {
	replyName, foundName := bin_api_reflect.ReplyNameFor(request.Name())
	if foundName {
		if reply, found := a.binApiTypes[replyName]; found {
			return reply, found
		}
	}

	return nil, false
}

// GetMsgID returns mocked message ID for the given message name and CRC.
func (a *VppAdapter) GetMsgID(msgName string, msgCrc string) (uint16, error) {
	switch msgName {
	case "control_ping":
		return 100, nil
	case "control_ping_reply":
		return 101, nil
	case "sw_interface_dump":
		return 200, nil
	case "sw_interface_details":
		return 201, nil
	}

	a.initMaps()

	if msgId, found := (*a.msgNameToIds)[msgName]; found {
		return msgId, nil
	} else {
		a.msgIdSeq++
		msgId = a.msgIdSeq
		(*a.msgNameToIds)[msgName] = msgId
		(*a.msgIdsToName)[msgId] = msgName
		return msgId, nil
	}
}

func (a *VppAdapter) initMaps() {
	if a.msgIdsToName == nil {
		a.msgIdsToName = &map[uint16]string{}
		a.msgNameToIds = &map[string]uint16{}
		a.msgIdSeq = 1000
	}

	if a.binApiTypes == nil {
		a.binApiTypes = map[string]reflect.Type{}
	}
}

// SendMsg emulates sending a binary-encoded message to VPP.
func (a *VppAdapter) SendMsg(clientID uint32, data []byte) error {
	msgID := uint16(defaultReplyMsgID)
	switch mode {
	case useRepliesQueue:
		repliesLock.Lock()
		defer repliesLock.Unlock()

		// pop all replies from queue
		for i, reply := range replies {
			if i > 0 && reply.GetMessageName() == "control_ping_reply" {
				// hack - do not send control_ping_reply immediately, leave it for the the next callback
				replies = nil
				replies = append(replies, reply)
				return nil
			}
			msgID, _ = a.GetMsgID(reply.GetMessageName(), reply.GetCrcString())
			buf := new(bytes.Buffer)
			if reply.GetMessageType() == api.ReplyMessage {
				struc.Pack(buf, &replyHeader{VlMsgID: msgID, Context: clientID})
			} else {
				struc.Pack(buf, &requestHeader{VlMsgID: msgID, Context: clientID})
			}
			struc.Pack(buf, reply)
			a.callback(clientID, msgID, buf.Bytes())
		}
		replies = nil
	case useReplyHandlers:
		for i := len(replyHandlers) - 1; i >= 0; i-- {
			callback := replyHandlers[i]

			buf := bytes.NewReader(data)
			reqHeader := requestHeader{}
			struc.Unpack(buf, &reqHeader)

			a.initMaps()
			reqMsgName, _ := (*a.msgIdsToName)[reqHeader.VlMsgID]

			reply, finished := callback(MessageDTO{reqHeader.VlMsgID, reqMsgName,
				clientID, data})
			if finished {
				buf := new(bytes.Buffer)
				struc.Pack(buf, &replyHeader{VlMsgID: msgID, Context: clientID})
				buf.Write(reply.Data)

				a.callback(clientID, msgID, buf.Bytes())
				break
			}
		}
		fallthrough
	default:
		// return default reply
		buf := new(bytes.Buffer)
		struc.Pack(buf, &replyHeader{VlMsgID: msgID, Context: clientID})
		struc.Pack(buf, &defaultReply{})
		a.callback(clientID, msgID, buf.Bytes())
	}
	return nil
}

// SetMsgCallback sets a callback function that will be called by the adapter whenever a message comes from the mock.
func (a *VppAdapter) SetMsgCallback(cb func(context uint32, msgID uint16, data []byte)) {
	a.callback = cb
}

// MockReply stores a message to be returned when the next request comes. It is a FIFO queue - multiple replies
// can be pushed into it, the first one will be popped when some request comes.
//
// It is able to also receive callback that calculates the reply
func (a *VppAdapter) MockReply(reply interface{}) {
	repliesLock.Lock()
	defer repliesLock.Unlock()

	if msg, ok := reply.(api.Message); ok {
		replies = append(replies, msg)
		mode = useRepliesQueue
	} else if replyHandler, ok := reply.(ReplyHandler); ok {
		replyHandlers = append(replyHandlers, replyHandler)
		mode = useReplyHandlers
	} else {
		//unknown
	}
}
