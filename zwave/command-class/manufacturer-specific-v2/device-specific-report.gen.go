// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package manufacturerspecificv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(DeviceSpecificReport{})
}

// <no value>
type DeviceSpecificReport struct {
	Properties1 struct {
		DeviceIdType byte
	}

	Properties2 struct {
		DeviceIdDataLengthIndicator byte

		DeviceIdDataFormat byte
	}

	DeviceIdData []byte
}

func (cmd DeviceSpecificReport) CommandClassID() byte {
	return 0x72
}

func (cmd DeviceSpecificReport) CommandID() byte {
	return byte(CommandDeviceSpecificReport)
}

func (cmd *DeviceSpecificReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.DeviceIdType = (payload[i] & 0x07)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.DeviceIdDataLengthIndicator = (payload[i] & 0x1F)

	cmd.Properties2.DeviceIdDataFormat = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DeviceIdData = payload[i:]

	return nil
}

func (cmd *DeviceSpecificReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	{
		var val byte

		val |= (cmd.Properties1.DeviceIdType) & byte(0x07)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.DeviceIdDataLengthIndicator) & byte(0x1F)

		val |= (cmd.Properties2.DeviceIdDataFormat << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.DeviceIdData...)

	return
}
