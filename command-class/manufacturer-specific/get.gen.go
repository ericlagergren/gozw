// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package manufacturerspecific

import (
	"encoding/gob"

	"github.com/helioslabs/gozw/command-class"
)

const CommandGet commandclass.CommandID = 0x04

func init() {
	gob.Register(Get{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x72),
		Command:      commandclass.CommandID(0x04),
		Version:      1,
	}, NewGet)
}

func NewGet() commandclass.Command {
	return &Get{}
}

// <no value>
type Get struct {
}

func (cmd Get) CommandClassID() commandclass.CommandClassID {
	return 0x72
}

func (cmd Get) CommandID() commandclass.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "MANUFACTURER_SPECIFIC_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
