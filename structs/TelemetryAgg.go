package structs

// TelemetryAgg struct representing mock telemetry
type TelemetryAgg struct {
	UTCEndTime              string  `json:"UtcEndTime"`
	MeterPVActivePowerCount int64   `json:"MeterPvActivePowerCount"`
	BmsSocTrimmedCount      int64   `json:"BmsSocTrimmedCount"`
	BmsCellVoltageAvg       float64 `json:"BmsCellVoltageAvg"`
}

// TelemetryAggTemp struct representing mock telemetry
type TelemetryAggTemp struct {
	UTCEndTime                     string  `json:"UtcEndTime"`
	BmsAirInletTemperatureMinCount int64   `json:"BmsAirInletTemperatureMinCount"`
	BmsAirInletTemperatureMin      float64 `json:"BmsAirInletTemperatureMin"`
	BmsAirInletTemperatureMax      float64 `json:"BmsAirInletTemperatureMax"`
	BmsAirInletTemperatureMaxCount int64   `json:"BmsAirInletTemperatureMaxCount"`
}
