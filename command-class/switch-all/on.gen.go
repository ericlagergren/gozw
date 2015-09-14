// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchall

import (
	"encoding/gob"

	"github.com/helioslabs/gozw/command-class"
)

const CommandOn commandclass.CommandID = 0x04

func init() {
	gob.Register(On{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x27),
		Command:      commandclass.CommandID(0x04),
		Version:      1,
	}, NewOn)
}

func NewOn() commandclass.Command {
	return &On{}
}

// <no value>
type On struct {
}

func (cmd On) CommandClassID() commandclass.CommandClassID {
	return 0x27
}

func (cmd On) CommandID() commandclass.CommandID {
	return CommandOn
}

func (cmd On) CommandIDString() string {
	return "SWITCH_ALL_ON"
}

func (cmd *On) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *On) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
