// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

// <no value>

type SecuritySchemeReport struct {
	SupportedSecuritySchemes byte
}

func ParseSecuritySchemeReport(payload []byte) SecuritySchemeReport {
	val := SecuritySchemeReport{}

	i := 2

	val.SupportedSecuritySchemes = payload[i]
	i++

	return val
}