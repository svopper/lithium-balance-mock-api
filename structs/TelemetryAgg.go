package structs

// TelemetryAgg struct representing mock telemetry
type TelemetryAgg struct {
	UTCEndTime              string  `json:"UtcEndTime"`
	MeterPVActivePowerCount int64   `json:"MeterPvActivePowerCount"`
	BmsSocTrimmedCount      int64   `json:"BmsSocTrimmedCount"`
	BmsCellVoltageAvg       float64 `json:"BmsCellVoltageAvg"`
}
