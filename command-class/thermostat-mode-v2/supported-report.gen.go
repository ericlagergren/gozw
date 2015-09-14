// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev2

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandSupportedReport commandclass.CommandID = 0x05

func init() {
	gob.Register(SupportedReport{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x40),
		Command:      commandclass.CommandID(0x05),
		Version:      2,
	}, NewSupportedReport)
}

func NewSupportedReport() commandclass.Command {
	return &SupportedReport{}
}

// <no value>
type SupportedReport struct {
	BitMask []byte
}

func (cmd SupportedReport) CommandClassID() commandclass.CommandClassID {
	return 0x40
}

func (cmd SupportedReport) CommandID() commandclass.CommandID {
	return CommandSupportedReport
}

func (cmd SupportedReport) CommandIDString() string {
	return "THERMOSTAT_MODE_SUPPORTED_REPORT"
}

func (cmd *SupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.BitMask = payload[i:]

	return nil
}

func (cmd *SupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.BitMask...)

	return
}
