// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

// <no value>

type NetworkKeySet struct {
	NetworkKeyByte []byte
}

func ParseNetworkKeySet(payload []byte) NetworkKeySet {
	val := NetworkKeySet{}

	i := 2

	val.NetworkKeyByte = payload[i:]

	return val
}