// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sceneactuatorconf

// <no value>

type SceneActuatorConfReport struct {
	SceneId byte

	Level byte

	DimmingDuration byte
}

func ParseSceneActuatorConfReport(payload []byte) SceneActuatorConfReport {
	val := SceneActuatorConfReport{}

	i := 2

	val.SceneId = payload[i]
	i++

	val.Level = payload[i]
	i++

	val.DimmingDuration = payload[i]
	i++

	return val
}