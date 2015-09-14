// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrol

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandStateGet commandclass.CommandID = 0x03

func init() {
	gob.Register(StateGet{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x33),
		Command:      commandclass.CommandID(0x03),
		Version:      1,
	}, NewStateGet)
}

func NewStateGet() commandclass.Command {
	return &StateGet{}
}

// <no value>
type StateGet struct {
	CapabilityId byte
}

func (cmd StateGet) CommandClassID() commandclass.CommandClassID {
	return 0x33
}

func (cmd StateGet) CommandID() commandclass.CommandID {
	return CommandStateGet
}

func (cmd StateGet) CommandIDString() string {
	return "STATE_GET"
}

func (cmd *StateGet) UnmarshalBinary(data []byte) error {
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

func (cmd *StateGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.CapabilityId)

	return
}
