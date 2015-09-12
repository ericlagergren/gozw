// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrol

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(StopStateChange{})
}

// <no value>
type StopStateChange struct {
	CapabilityId byte
}

func (cmd StopStateChange) CommandClassID() byte {
	return 0x33
}

func (cmd StopStateChange) CommandID() byte {
	return byte(CommandStopStateChange)
}

func (cmd *StopStateChange) UnmarshalBinary(data []byte) error {
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

	cmd.CapabilityId = payload[i]
	i++

	return nil
}

func (cmd *StopStateChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.CapabilityId)

	return
}
