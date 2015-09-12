// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv3

import "fmt"

type CommandID byte

const (
	CommandGet CommandID = 0x04

	CommandReport CommandID = 0x05

	CommandSet CommandID = 0x06

	CommandSupportedGet CommandID = 0x07

	CommandSupportedReport CommandID = 0x08

	CommandEventSupportedGet CommandID = 0x01

	CommandEventSupportedReport CommandID = 0x02
)

func (c CommandID) String() string {
	switch c {

	case CommandGet:
		return "NOTIFICATION_GET"

	case CommandReport:
		return "NOTIFICATION_REPORT"

	case CommandSet:
		return "NOTIFICATION_SET"

	case CommandSupportedGet:
		return "NOTIFICATION_SUPPORTED_GET"

	case CommandSupportedReport:
		return "NOTIFICATION_SUPPORTED_REPORT"

	case CommandEventSupportedGet:
		return "EVENT_SUPPORTED_GET"

	case CommandEventSupportedReport:
		return "EVENT_SUPPORTED_REPORT"

	default:
		return fmt.Sprintf("Unknown (0x%X)", byte(c))
	}
}
