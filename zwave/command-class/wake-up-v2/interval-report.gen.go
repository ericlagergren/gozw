// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeupv2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(IntervalReport{})
}

// <no value>
type IntervalReport struct {
	Seconds uint32

	Nodeid byte
}

func (cmd IntervalReport) CommandClassID() byte {
	return 0x84
}

func (cmd IntervalReport) CommandID() byte {
	return byte(CommandIntervalReport)
}

func (cmd *IntervalReport) UnmarshalBinary(data []byte) error {
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

	cmd.Seconds = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Nodeid = payload[i]
	i++

	return nil
}

func (cmd *IntervalReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Seconds)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	payload = append(payload, cmd.Nodeid)

	return
}
