// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package version

import "encoding/gob"

func init() {
	gob.Register(Get{})
}

// <no value>
type Get struct {
}

func (cmd Get) CommandClassID() byte {
	return 0x86
}

func (cmd Get) CommandID() byte {
	return byte(CommandGet)
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	return
}
