// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementinclusion

// <no value>

type NodeNeighborUpdateStatus struct {
	SeqNo byte

	Status byte
}

func ParseNodeNeighborUpdateStatus(payload []byte) NodeNeighborUpdateStatus {
	val := NodeNeighborUpdateStatus{}

	i := 2

	val.SeqNo = payload[i]
	i++

	val.Status = payload[i]
	i++

	return val
}