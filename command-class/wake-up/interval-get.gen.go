// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeup

import (
	"encoding/gob"

	"github.com/helioslabs/gozw/command-class"
)

const CommandIntervalGet commandclass.CommandID = 0x05

func init() {
	gob.Register(IntervalGet{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x84),
		Command:      commandclass.CommandID(0x05),
		Version:      1,
	}, NewIntervalGet)
}

func NewIntervalGet() commandclass.Command {
	return &IntervalGet{}
}

// <no value>
type IntervalGet struct {
}

func (cmd IntervalGet) CommandClassID() commandclass.CommandClassID {
	return 0x84
}

func (cmd IntervalGet) CommandID() commandclass.CommandID {
	return CommandIntervalGet
}

func (cmd IntervalGet) CommandIDString() string {
	return "WAKE_UP_INTERVAL_GET"
}

func (cmd *IntervalGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *IntervalGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
