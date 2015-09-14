// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"
	"errors"

	"github.com/helioslabs/gozw/command-class"
)

const CommandNetworkKeySet commandclass.CommandID = 0x06

func init() {
	gob.Register(NetworkKeySet{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x98),
		Command:      commandclass.CommandID(0x06),
		Version:      1,
	}, NewNetworkKeySet)
}

func NewNetworkKeySet() commandclass.Command {
	return &NetworkKeySet{}
}

// <no value>
type NetworkKeySet struct {
	NetworkKeyByte []byte
}

func (cmd NetworkKeySet) CommandClassID() commandclass.CommandClassID {
	return 0x98
}

func (cmd NetworkKeySet) CommandID() commandclass.CommandID {
	return CommandNetworkKeySet
}

func (cmd NetworkKeySet) CommandIDString() string {
	return "NETWORK_KEY_SET"
}

func (cmd *NetworkKeySet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return nil
	}

	cmd.NetworkKeyByte = payload[i:]

	return nil
}

func (cmd *NetworkKeySet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.NetworkKeyByte...)

	return
}
