// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrol

import "encoding/gob"

func init() {
	gob.Register(CapabilityGet{})
}

// <no value>
type CapabilityGet struct {
}

func (cmd CapabilityGet) CommandClassID() byte {
	return 0x33
}

func (cmd CapabilityGet) CommandID() byte {
	return byte(CommandCapabilityGet)
}

func (cmd *CapabilityGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *CapabilityGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	return
}