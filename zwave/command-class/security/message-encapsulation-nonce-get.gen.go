// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(MessageEncapsulationNonceGet{})
}

// <no value>
type MessageEncapsulationNonceGet struct {
	InitializationVectorByte []byte

	Properties1 struct {
		SequenceCounter byte

		Sequenced bool

		SecondFrame bool
	}

	CommandClassIdentifier byte

	CommandIdentifier byte

	CommandByte []byte

	ReceiversNonceIdentifier byte

	MessageAuthenticationCodeByte []byte
}

func (cmd MessageEncapsulationNonceGet) CommandClassID() byte {
	return 0x98
}

func (cmd MessageEncapsulationNonceGet) CommandID() byte {
	return byte(CommandMessageEncapsulationNonceGet)
}

func (cmd *MessageEncapsulationNonceGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.InitializationVectorByte = payload[i : i+8]

	i += 8

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.SequenceCounter = (payload[i] & 0x0F)

	cmd.Properties1.Sequenced = payload[i]&0x10 == 0x10

	cmd.Properties1.SecondFrame = payload[i]&0x20 == 0x20

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClassIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandByte = payload[i : len(payload)-9]
	i += len(cmd.CommandByte)

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReceiversNonceIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MessageAuthenticationCodeByte = payload[i : i+8]

	i += 8

	return nil
}

func (cmd *MessageEncapsulationNonceGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	if paramLen := len(cmd.InitializationVectorByte); paramLen > 8 {
		return nil, errors.New("Length overflow in array parameter InitializationVectorByte")
	}

	payload = append(payload, cmd.InitializationVectorByte...)

	{
		var val byte

		val |= (cmd.Properties1.SequenceCounter) & byte(0x0F)

		if cmd.Properties1.Sequenced {
			val |= byte(0x10) // flip bits on
		} else {
			val &= ^byte(0x10) // flip bits off
		}

		if cmd.Properties1.SecondFrame {
			val |= byte(0x20) // flip bits on
		} else {
			val &= ^byte(0x20) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.CommandClassIdentifier)

	payload = append(payload, cmd.CommandIdentifier)

	payload = append(payload, cmd.CommandByte...)

	payload = append(payload, cmd.ReceiversNonceIdentifier)

	if paramLen := len(cmd.MessageAuthenticationCodeByte); paramLen > 8 {
		return nil, errors.New("Length overflow in array parameter MessageAuthenticationCodeByte")
	}

	payload = append(payload, cmd.MessageAuthenticationCodeByte...)

	return
}
