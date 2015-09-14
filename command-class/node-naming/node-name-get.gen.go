// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package nodenaming

import (
	"encoding/gob"

	"github.com/helioslabs/gozw/command-class"
)

const CommandNodeNameGet commandclass.CommandID = 0x02

func init() {
	gob.Register(NodeNameGet{})
	commandclass.Register(commandclass.CommandIdentifier{
		CommandClass: commandclass.CommandClassID(0x77),
		Command:      commandclass.CommandID(0x02),
		Version:      1,
	}, NewNodeNameGet)
}

func NewNodeNameGet() commandclass.Command {
	return &NodeNameGet{}
}

// <no value>
type NodeNameGet struct {
}

func (cmd NodeNameGet) CommandClassID() commandclass.CommandClassID {
	return 0x77
}

func (cmd NodeNameGet) CommandID() commandclass.CommandID {
	return CommandNodeNameGet
}

func (cmd NodeNameGet) CommandIDString() string {
	return "NODE_NAMING_NODE_NAME_GET"
}

func (cmd *NodeNameGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *NodeNameGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
