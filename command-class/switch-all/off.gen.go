// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchall

import (
	"encoding/gob"

	"github.com/helioslabs/gozw/command-class"
)

const CommandOff commandclass.CommandID = 0x05

func init() {
	gob.Register(Off{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x27),
		Command:      commandclass.CommandID(0x05),
		Version:      1,
	}, NewOff)
}

func NewOff() commandclass.Command {
	return &Off{}
}

// <no value>
type Off struct {
}

func (cmd Off) CommandClassID() commandclass.CommandClassID {
	return 0x27
}

func (cmd Off) CommandID() commandclass.CommandID {
	return CommandOff
}

func (cmd Off) CommandIDString() string {
	return "SWITCH_ALL_OFF"
}

func (cmd *Off) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Off) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
