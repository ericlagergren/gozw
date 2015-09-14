// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensorconfiguration

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandSensorTriggerLevelSet commandclass.CommandID = 0x01

func init() {
	gob.Register(SensorTriggerLevelSet{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x9E),
		Command:      commandclass.CommandID(0x01),
		Version:      1,
	}, NewSensorTriggerLevelSet)
}

func NewSensorTriggerLevelSet() commandclass.Command {
	return &SensorTriggerLevelSet{}
}

// <no value>
type SensorTriggerLevelSet struct {
	Properties1 struct {
		Current bool

		Default bool
	}

	SensorType byte

	Properties2 struct {
		Size byte

		Scale byte

		Precision byte
	}

	TriggerValue []byte
}

func (cmd SensorTriggerLevelSet) CommandClassID() commandclass.CommandClassID {
	return 0x9E
}

func (cmd SensorTriggerLevelSet) CommandID() commandclass.CommandID {
	return CommandSensorTriggerLevelSet
}

func (cmd SensorTriggerLevelSet) CommandIDString() string {
	return "SENSOR_TRIGGER_LEVEL_SET"
}

func (cmd *SensorTriggerLevelSet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.Current = payload[i]&0x40 == 0x40

	cmd.Properties1.Default = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.Size = (payload[i] & 0x07)

	cmd.Properties2.Scale = (payload[i] & 0x18) >> 3

	cmd.Properties2.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[2+2] >> 0) & 0x07
		cmd.TriggerValue = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *SensorTriggerLevelSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		if cmd.Properties1.Current {
			val |= byte(0x40) // flip bits on
		} else {
			val &= ^byte(0x40) // flip bits off
		}

		if cmd.Properties1.Default {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Properties2.Size) & byte(0x07)

		val |= (cmd.Properties2.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Properties2.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.TriggerValue != nil && len(cmd.TriggerValue) > 0 {
		payload = append(payload, cmd.TriggerValue...)
	}

	return
}
