// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package configurationv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(Report{})
}

// <no value>
type Report struct {
	ParameterNumber byte

	Level struct {
		Size byte
	}

	ConfigurationValue []byte
}

func (cmd Report) CommandClassID() byte {
	return 0x70
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

	cmd.ParameterNumber = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.Size = (payload[i] & 0x07)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[1+2] >> 0) & 0x07
		cmd.ConfigurationValue = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.ParameterNumber)

	{
		var val byte

		val |= (cmd.Level.Size) & byte(0x07)

		payload = append(payload, val)
	}

	if cmd.ConfigurationValue != nil && len(cmd.ConfigurationValue) > 0 {
		payload = append(payload, cmd.ConfigurationValue...)
	}

	return
}
