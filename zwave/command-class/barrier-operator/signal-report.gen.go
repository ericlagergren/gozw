// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package barrieroperator

// <no value>

type BarrierOperatorSignalReport struct {
	SubsystemType byte

	SubsystemState byte
}

func ParseBarrierOperatorSignalReport(payload []byte) BarrierOperatorSignalReport {
	val := BarrierOperatorSignalReport{}

	i := 2

	val.SubsystemType = payload[i]
	i++

	val.SubsystemState = payload[i]
	i++

	return val
}