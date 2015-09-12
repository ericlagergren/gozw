// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package time

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(Report{})
}

// <no value>
type Report struct {
	HourLocalTime struct {
		HourLocalTime byte

		RtcFailure bool
	}

	MinuteLocalTime byte

	SecondLocalTime byte
}

func (cmd Report) CommandClassID() byte {
	return 0x8A
}

func (cmd Report) CommandID() byte {
	return byte(CommandReport)
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.HourLocalTime.HourLocalTime = (payload[i] & 0x1F)

	cmd.HourLocalTime.RtcFailure = payload[i]&0x80 == 0x80

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

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	{
		var val byte

		val |= (cmd.HourLocalTime.HourLocalTime) & byte(0x1F)

		if cmd.HourLocalTime.RtcFailure {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.MinuteLocalTime)

	payload = append(payload, cmd.SecondLocalTime)

	return
}
