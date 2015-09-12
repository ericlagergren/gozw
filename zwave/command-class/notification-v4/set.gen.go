// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv4

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(Set{})
}

// <no value>
type Set struct {
	NotificationType byte

	NotificationStatus byte
}

func (cmd Set) CommandClassID() byte {
	return 0x71
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

	cmd.NotificationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationStatus = payload[i]
	i++

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.NotificationType)

	payload = append(payload, cmd.NotificationStatus)

	return
}
