// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import (
	"encoding/gob"

	"github.com/helioslabs/gozw/command-class"
)

const CommandSpecificGroupGet commandclass.CommandID = 0x0B

func init() {
	gob.Register(SpecificGroupGet{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x85),
		Command:      commandclass.CommandID(0x0B),
		Version:      2,
	}, NewSpecificGroupGet)
}

func NewSpecificGroupGet() commandclass.Command {
	return &SpecificGroupGet{}
}

// <no value>
type SpecificGroupGet struct {
}

func (cmd SpecificGroupGet) CommandClassID() commandclass.CommandClassID {
	return 0x85
}

func (cmd SpecificGroupGet) CommandID() commandclass.CommandID {
	return CommandSpecificGroupGet
}

func (cmd SpecificGroupGet) CommandIDString() string {
	return "ASSOCIATION_SPECIFIC_GROUP_GET"
}

func (cmd *SpecificGroupGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SpecificGroupGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
