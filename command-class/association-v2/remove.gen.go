// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandRemove commandclass.CommandID = 0x04

func init() {
	gob.Register(Remove{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x85),
		Command:      commandclass.CommandID(0x04),
		Version:      2,
	}, NewRemove)
}

func NewRemove() commandclass.Command {
	return &Remove{}
}

// <no value>
type Remove struct {
	GroupingIdentifier byte

	NodeId []byte
}

func (cmd Remove) CommandClassID() commandclass.CommandClassID {
	return 0x85
}

func (cmd Remove) CommandID() commandclass.CommandID {
	return CommandRemove
}

func (cmd Remove) CommandIDString() string {
	return "ASSOCIATION_REMOVE"
}

func (cmd *Remove) UnmarshalBinary(data []byte) error {
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

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return nil
	}

	cmd.NodeId = payload[i:]

	return nil
}

func (cmd *Remove) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.NodeId...)

	return
}
