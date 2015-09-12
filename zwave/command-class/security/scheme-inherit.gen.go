// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SchemeInherit{})
}

// <no value>
type SchemeInherit struct {
	SupportedSecuritySchemes byte
}

func (cmd SchemeInherit) CommandClassID() byte {
	return 0x98
}

func (cmd SchemeInherit) CommandID() byte {
	return byte(CommandSchemeInherit)
}

func (cmd *SchemeInherit) UnmarshalBinary(data []byte) error {
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

	cmd.SupportedSecuritySchemes = payload[i]
	i++

	return nil
}

func (cmd *SchemeInherit) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.SupportedSecuritySchemes)

	return
}
