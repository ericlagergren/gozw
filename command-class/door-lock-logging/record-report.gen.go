// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package doorlocklogging

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandRecordReport commandclass.CommandID = 0x04

func init() {
	gob.Register(RecordReport{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x4C),
		Command:      commandclass.CommandID(0x04),
		Version:      1,
	}, NewRecordReport)
}

func NewRecordReport() commandclass.Command {
	return &RecordReport{}
}

// <no value>
type RecordReport struct {
	RecordNumber byte

	Year uint16

	Month byte

	Day byte

	Properties1 struct {
		HourLocalTime byte

		RecordStatus byte
	}

	MinuteLocalTime byte

	SecondLocalTime byte

	EventType byte

	UserIdentifier byte

	UserCodeLength byte

	UserCode []byte
}

func (cmd RecordReport) CommandClassID() commandclass.CommandClassID {
	return 0x4C
}

func (cmd RecordReport) CommandID() commandclass.CommandID {
	return CommandRecordReport
}

func (cmd RecordReport) CommandIDString() string {
	return "RECORD_REPORT"
}

func (cmd *RecordReport) UnmarshalBinary(data []byte) error {
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

	cmd.RecordNumber = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Month = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Day = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.HourLocalTime = (payload[i] & 0x1F)

	cmd.Properties1.RecordStatus = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SecondLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.EventType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserCodeLength = payload[i]
	i++

	if len(payload) <= i {
		return nil
	}

	cmd.UserCode = payload[i:]

	return nil
}

func (cmd *RecordReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.RecordNumber)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Year)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.Month)

	payload = append(payload, cmd.Day)

	{
		var val byte

		val |= (cmd.Properties1.HourLocalTime) & byte(0x1F)

		val |= (cmd.Properties1.RecordStatus << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.MinuteLocalTime)

	payload = append(payload, cmd.SecondLocalTime)

	payload = append(payload, cmd.EventType)

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.UserCodeLength)

	payload = append(payload, cmd.UserCode...)

	return
}
