// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev3

import "fmt"

type CommandID byte

const (
	CommandGet CommandID = 0x02

	CommandReport CommandID = 0x03

	CommandSet CommandID = 0x01

	CommandSupportedGet CommandID = 0x04

	CommandSupportedReport CommandID = 0x05
)

func (c CommandID) String() string {
	switch c {

	case CommandGet:
		return "THERMOSTAT_MODE_GET"

	case CommandReport:
		return "THERMOSTAT_MODE_REPORT"

	case CommandSet:
		return "THERMOSTAT_MODE_SET"

	case CommandSupportedGet:
		return "THERMOSTAT_MODE_SUPPORTED_GET"

	case CommandSupportedReport:
		return "THERMOSTAT_MODE_SUPPORTED_REPORT"

	default:
		return fmt.Sprintf("Unknown (0x%X)", byte(c))
	}
}
