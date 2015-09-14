// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmode

import (
	"encoding/gob"

	"github.com/helioslabs/gozw/command-class"
)

const CommandSupportedGet commandclass.CommandID = 0x04

func init() {
	gob.Register(SupportedGet{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x40),
		Command:      commandclass.CommandID(0x04),
		Version:      1,
	}, NewSupportedGet)
}

func NewSupportedGet() commandclass.Command {
	return &SupportedGet{}
}

// <no value>
type SupportedGet struct {
}

func (cmd SupportedGet) CommandClassID() commandclass.CommandClassID {
	return 0x40
}

func (cmd SupportedGet) CommandID() commandclass.CommandID {
	return CommandSupportedGet
}

func (cmd SupportedGet) CommandIDString() string {
	return "THERMOSTAT_MODE_SUPPORTED_GET"
}

func (cmd *SupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
