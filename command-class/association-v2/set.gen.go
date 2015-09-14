// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandSet commandclass.CommandID = 0x01

func init() {
	gob.Register(Set{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x85),
		Command:      commandclass.CommandID(0x01),
		Version:      2,
	}, NewSet)
}

func NewSet() commandclass.Command {
	return &Set{}
}

// <no value>
type Set struct {
	GroupingIdentifier byte

	NodeId []byte
}

func (cmd Set) CommandClassID() commandclass.CommandClassID {
	return 0x85
}

func (cmd Set) CommandID() commandclass.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "ASSOCIATION_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
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

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.NodeId...)

	return
}
