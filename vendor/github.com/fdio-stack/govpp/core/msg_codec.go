package core

import (
	"bytes"
	"reflect"

	logger "github.com/Sirupsen/logrus"
	"github.com/lunixbochs/struc"

	"errors"

	"fmt"

	"github.com/fdio-stack/govpp/api"
)

// MsgCodec provides encoding and decoding functionality of `api.Message` structs into/from
// binary format as accepted by VPP.
type MsgCodec struct{}

// vppRequestHeader struct contains header fields implemented by all VPP requests.
type vppRequestHeader struct {
	VlMsgID     uint16
	ClientIndex uint32
	Context     uint32
}

// vppReplyHeader struct contains header fields implemented by all VPP replies.
type vppReplyHeader struct {
	VlMsgID uint16
	Context uint32
}

// vppOtherHeader struct contains header fields implemented by other VPP messages (not requests nor replies).
type vppOtherHeader struct {
	VlMsgID uint16
}

const (
	vppRequestHeaderSize = 10 // size of a VPP request header
	vppReplyHeaderSize   = 6  // size of a VPP reply header
	vppOtherHeaderSize   = 2  // size of the header of other VPP messages
)

// EncodeMsg encodes provided `Message` structure into its binary-encoded data representation.
func (*MsgCodec) EncodeMsg(msg api.Message, msgID uint16) ([]byte, error) {
	if msg == nil {
		return nil, errors.New("nil message passed in")
	}

	buf := new(bytes.Buffer)

	// encode message header
	var header interface{}
	if msg.GetMessageType() == api.RequestMessage {
		header = &vppRequestHeader{VlMsgID: msgID}
	} else if msg.GetMessageType() == api.ReplyMessage {
		header = &vppReplyHeader{VlMsgID: msgID}
	} else {
		header = &vppOtherHeader{VlMsgID: msgID}
	}
	err := struc.Pack(buf, header)
	if err != nil {
		log.WithFields(logger.Fields{
			"error":  err,
			"header": header,
		}).Error("Unable to encode the message header: ", err)
		return nil, fmt.Errorf("unable to encode the message header: %v", err)
	}

	// encode message content
	if reflect.Indirect(reflect.ValueOf(msg)).NumField() > 0 {
		err := struc.Pack(buf, msg)
		if err != nil {
			log.WithFields(logger.Fields{
				"error":   err,
				"message": msg,
			}).Error("Unable to encode the message: ", err)
			return nil, fmt.Errorf("unable to encode the message: %v", err)
		}
	}

	return buf.Bytes(), nil
}

// DecodeMsg decodes binary-encoded data of a message into provided `Message` structure.
func (*MsgCodec) DecodeMsg(data []byte, msg api.Message) error {
	if msg == nil {
		return errors.New("nil message passed in")
	}

	buf := bytes.NewReader(data)

	// check which header is expected
	var header interface{}
	if msg.GetMessageType() == api.RequestMessage {
		header = &vppRequestHeader{}
	} else if msg.GetMessageType() == api.ReplyMessage {
		header = &vppReplyHeader{}
	} else {
		header = &vppOtherHeader{}
	}

	// decode message header
	err := struc.Unpack(buf, header)
	if err != nil {
		log.WithFields(logger.Fields{
			"error": err,
			"data":  data,
		}).Error("Unable to decode header of the message.")
		return fmt.Errorf("unable to decode the message header: %v", err)
	}

	// get rid of the message header
	if msg.GetMessageType() == api.RequestMessage {
		buf = bytes.NewReader(data[vppRequestHeaderSize:])
	} else if msg.GetMessageType() == api.ReplyMessage {
		buf = bytes.NewReader(data[vppReplyHeaderSize:])
	} else {
		buf = bytes.NewReader(data[vppOtherHeaderSize:])
	}

	// decode message content
	err = struc.Unpack(buf, msg)
	if err != nil {
		log.WithFields(logger.Fields{
			"error": err,
			"data":  buf,
		}).Error("Unable to decode the message.")
		return fmt.Errorf("unable to decode the message: %v", err)
	}

	return nil
}
