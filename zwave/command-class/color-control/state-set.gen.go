// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrol

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(StateSet{})
}

// <no value>
type StateSet struct {
	Properties1 struct {
		StateDataLength byte
	}
}

func (cmd StateSet) CommandClassID() byte {
	return 0x33
}

func (cmd StateSet) CommandID() byte {
	return byte(CommandStateSet)
}

func (cmd *StateSet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.StateDataLength = (payload[i] & 0x1F)

	i += 1

	return nil
}

func (cmd *StateSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	{
		var val byte

		val |= (cmd.Properties1.StateDataLength) & byte(0x1F)

		payload = append(payload, val)
	}

	return
}
