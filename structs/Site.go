package structs

// Site struct representing any site
type Site struct {
	SiteID                   string  `json:"siteId"`
	SiteName                 string  `json:"siteName"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	PostalAddress            string  `json:"postalAddress"`
	Country                  string  `json:"country"`
	InstallationTypeInternal string  `json:"installationTypeInternal"`
	InstallationTypeExternal string  `json:"installationTypeExternal"`
	Status                   string  `json:"status"`
	RackTotal                int64   `json:"rackTotal"`
	EnergyKWh                int64   `json:"energyKWh"`
	PowerMaxKW               int64   `json:"powerMaxKW"`
	CommissioningDate        string  `json:"commissioningDate"`
	DecommissioningDate      string  `json:"decommissioningDate"`
	PartitionKey             string  `json:"partitionKey"`
	RowKey                   string  `json:"rowKey"`
	Timestamp                string  `json:"timestamp"`
	ETag                     string  `json:"eTag"`
}
