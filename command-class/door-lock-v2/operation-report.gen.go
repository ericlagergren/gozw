// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package doorlockv2

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandOperationReport commandclass.CommandID = 0x03

func init() {
	gob.Register(OperationReport{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x62),
		Command:      commandclass.CommandID(0x03),
		Version:      2,
	}, NewOperationReport)
}

func NewOperationReport() commandclass.Command {
	return &OperationReport{}
}

// <no value>
type OperationReport struct {
	DoorLockMode byte

	Properties1 struct {
		InsideDoorHandlesMode byte

		OutsideDoorHandlesMode byte
	}

	DoorCondition byte

	LockTimeoutMinutes byte

	LockTimeoutSeconds byte
}

func (cmd OperationReport) CommandClassID() commandclass.CommandClassID {
	return 0x62
}

func (cmd OperationReport) CommandID() commandclass.CommandID {
	return CommandOperationReport
}

func (cmd OperationReport) CommandIDString() string {
	return "DOOR_LOCK_OPERATION_REPORT"
}

func (cmd *OperationReport) UnmarshalBinary(data []byte) error {
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

	cmd.DoorLockMode = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.InsideDoorHandlesMode = (payload[i] & 0x0F)

	cmd.Properties1.OutsideDoorHandlesMode = (payload[i] & 0xF0) >> 4

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DoorCondition = payload[i]
	i++

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

func (cmd *OperationReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.DoorLockMode)

	{
		var val byte

		val |= (cmd.Properties1.InsideDoorHandlesMode) & byte(0x0F)

		val |= (cmd.Properties1.OutsideDoorHandlesMode << byte(4)) & byte(0xF0)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.DoorCondition)

	payload = append(payload, cmd.LockTimeoutMinutes)

	payload = append(payload, cmd.LockTimeoutSeconds)

	return
}
