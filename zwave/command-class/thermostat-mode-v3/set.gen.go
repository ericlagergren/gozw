// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev3

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(Set{})
}

// <no value>
type Set struct {
	Level struct {
		NoOfManufacturerDataFields byte

		Mode byte
	}

	ManufacturerData []byte
}

func (cmd Set) CommandClassID() byte {
	return 0x40
}

func (cmd Set) CommandID() byte {
	return byte(CommandSet)
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
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

	cmd.Level.NoOfManufacturerDataFields = (payload[i] & 0xE0) >> 5

	cmd.Level.Mode = (payload[i] & 0x1F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[0+2] >> 5) & 0xE0
		cmd.ManufacturerData = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	{
		var val byte

		val |= (cmd.Level.NoOfManufacturerDataFields << byte(5)) & byte(0xE0)

		val |= (cmd.Level.Mode) & byte(0x1F)

		payload = append(payload, val)
	}

	if cmd.ManufacturerData != nil && len(cmd.ManufacturerData) > 0 {
		payload = append(payload, cmd.ManufacturerData...)
	}

	return
}
