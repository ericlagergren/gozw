// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package manufacturerspecific

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(Report{})
}

// <no value>
type Report struct {
	ManufacturerId uint16

	ProductTypeId uint16

	ProductId uint16
}

func (cmd Report) CommandClassID() byte {
	return 0x72
}

func (cmd Report) CommandID() byte {
	return byte(CommandReport)
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductTypeId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ManufacturerId)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ProductTypeId)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ProductId)
		payload = append(payload, buf...)
	}

	return
}
