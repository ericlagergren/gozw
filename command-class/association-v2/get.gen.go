// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandGet commandclass.CommandID = 0x02

func init() {
	gob.Register(Get{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x85),
		Command:      commandclass.CommandID(0x02),
		Version:      2,
	}, NewGet)
}

func NewGet() commandclass.Command {
	return &Get{}
}

// <no value>
type Get struct {
	GroupingIdentifier byte
}

func (cmd Get) CommandClassID() commandclass.CommandClassID {
	return 0x85
}

func (cmd Get) CommandID() commandclass.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "ASSOCIATION_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
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

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.GroupingIdentifier)

	return
}
