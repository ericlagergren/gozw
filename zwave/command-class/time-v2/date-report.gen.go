// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package timev2

import "encoding/binary"

// <no value>

type DateReport struct {
	Year uint16

	Month byte

	Day byte
}

func ParseDateReport(payload []byte) DateReport {
	val := DateReport{}

	i := 2

	val.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.Month = payload[i]
	i++

	val.Day = payload[i]
	i++

	return val
}