// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandSpecificGroupReport commandclass.CommandID = 0x0C

func init() {
	gob.Register(SpecificGroupReport{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x85),
		Command:      commandclass.CommandID(0x0C),
		Version:      2,
	}, NewSpecificGroupReport)
}

func NewSpecificGroupReport() commandclass.Command {
	return &SpecificGroupReport{}
}

// <no value>
type SpecificGroupReport struct {
	Group byte
}

func (cmd SpecificGroupReport) CommandClassID() commandclass.CommandClassID {
	return 0x85
}

func (cmd SpecificGroupReport) CommandID() commandclass.CommandID {
	return CommandSpecificGroupReport
}

func (cmd SpecificGroupReport) CommandIDString() string {
	return "ASSOCIATION_SPECIFIC_GROUP_REPORT"
}

func (cmd *SpecificGroupReport) UnmarshalBinary(data []byte) error {
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

	cmd.Group = payload[i]
	i++

	return nil
}

func (cmd *SpecificGroupReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Group)

	return
}
