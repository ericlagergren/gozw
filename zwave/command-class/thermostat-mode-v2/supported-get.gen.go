// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev2

import "encoding/gob"

func init() {
	gob.Register(SupportedGet{})
}

// <no value>
type SupportedGet struct {
}

func (cmd SupportedGet) CommandClassID() byte {
	return 0x40
}

func (cmd SupportedGet) CommandID() byte {
	return byte(CommandSupportedGet)
}

func (cmd *SupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	return
}
