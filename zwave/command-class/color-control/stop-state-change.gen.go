// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrol

// <no value>

type StopStateChange struct {
	CapabilityId byte
}

func ParseStopStateChange(payload []byte) StopStateChange {
	val := StopStateChange{}

	i := 2

	val.CapabilityId = payload[i]
	i++

	return val
}