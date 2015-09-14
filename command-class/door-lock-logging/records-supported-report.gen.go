// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package doorlocklogging

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandRecordsSupportedReport commandclass.CommandID = 0x02

func init() {
	gob.Register(RecordsSupportedReport{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x4C),
		Command:      commandclass.CommandID(0x02),
		Version:      1,
	}, NewRecordsSupportedReport)
}

func NewRecordsSupportedReport() commandclass.Command {
	return &RecordsSupportedReport{}
}

// <no value>
type RecordsSupportedReport struct {
	MaxRecordsStored byte
}

func (cmd RecordsSupportedReport) CommandClassID() commandclass.CommandClassID {
	return 0x4C
}

func (cmd RecordsSupportedReport) CommandID() commandclass.CommandID {
	return CommandRecordsSupportedReport
}

func (cmd RecordsSupportedReport) CommandIDString() string {
	return "DOOR_LOCK_LOGGING_RECORDS_SUPPORTED_REPORT"
}

func (cmd *RecordsSupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.MaxRecordsStored = payload[i]
	i++

	return nil
}

func (cmd *RecordsSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.MaxRecordsStored)

	return
}
