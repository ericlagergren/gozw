// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package versionv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(Report{})
}

// <no value>
type Report struct {
	ZWaveLibraryType byte

	ZWaveProtocolVersion byte

	ZWaveProtocolSubVersion byte

	Firmware0Version byte

	Firmware0SubVersion byte

	HardwareVersion byte

	NumberOfFirmwareTargets byte
}

func (cmd Report) CommandClassID() byte {
	return 0x86
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

	cmd.ZWaveLibraryType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZWaveProtocolVersion = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZWaveProtocolSubVersion = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Firmware0Version = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Firmware0SubVersion = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HardwareVersion = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfFirmwareTargets = payload[i]
	i++

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.ZWaveLibraryType)

	payload = append(payload, cmd.ZWaveProtocolVersion)

	payload = append(payload, cmd.ZWaveProtocolSubVersion)

	payload = append(payload, cmd.Firmware0Version)

	payload = append(payload, cmd.Firmware0SubVersion)

	payload = append(payload, cmd.HardwareVersion)

	payload = append(payload, cmd.NumberOfFirmwareTargets)

	return
}
