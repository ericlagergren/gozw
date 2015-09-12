// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(LoggingGet{})
}

// <no value>
type LoggingGet struct {
	BitMask []byte
}

func (cmd LoggingGet) CommandClassID() byte {
	return 0x42
}

func (cmd LoggingGet) CommandID() byte {
	return byte(CommandLoggingGet)
}

func (cmd *LoggingGet) UnmarshalBinary(data []byte) error {
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

func (cmd *LoggingGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.BitMask...)

	return
}
