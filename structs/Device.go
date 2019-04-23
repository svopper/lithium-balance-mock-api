package structs

// Device struct representing any device
type Device struct {
	DeviceID           string  `json:"DeviceId"`
	Latitude           float64 `json:"Latitude"`
	Longitude          float64 `json:"Longitude"`
	SiteName           string  `json:"SiteName"`
	SiteID             string  `json:"SiteId"`
	RackTotal          int64   `json:"RackTotal"`
	CellVoltageV       float64 `json:"CellVoltageV"`
	InverterPowerMaxKW float64 `json:"InverterPowerMaxKW"`
}
