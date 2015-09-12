// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package doorlock

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ConfigurationSet{})
}

// <no value>
type ConfigurationSet struct {
	OperationType byte

	Properties1 struct {
		InsideDoorHandlesState byte

		OutsideDoorHandlesState byte
	}

	LockTimeoutMinutes byte

	LockTimeoutSeconds byte
}

func (cmd ConfigurationSet) CommandClassID() byte {
	return 0x62
}

func (cmd ConfigurationSet) CommandID() byte {
	return byte(CommandConfigurationSet)
}

func (cmd *ConfigurationSet) UnmarshalBinary(data []byte) error {
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

	cmd.OperationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.InsideDoorHandlesState = (payload[i] & 0x0F)

	cmd.Properties1.OutsideDoorHandlesState = (payload[i] & 0xF0) >> 4

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LockTimeoutMinutes = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LockTimeoutSeconds = payload[i]
	i++

	return nil
}

func (cmd *ConfigurationSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.OperationType)

	{
		var val byte

		val |= (cmd.Properties1.InsideDoorHandlesState) & byte(0x0F)

		val |= (cmd.Properties1.OutsideDoorHandlesState << byte(4)) & byte(0xF0)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.LockTimeoutMinutes)

	payload = append(payload, cmd.LockTimeoutSeconds)

	return
}
