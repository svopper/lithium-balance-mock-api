package structs

import "time"

// BmsSocTrimmed struct representing BmsSocTrimmed last 1
type BmsSocTrimmed struct {
	UTCTime       time.Time `json:"UtcTime"`
	BmsSocTrimmed int64     `json:"BmsSocTrimmed"`
}

// InverterActivePower struct representing InverterActivePower last 1
type InverterActivePower struct {
	UTCTime             time.Time  `json:"UtcTime"`            
	InverterActivePower float64    `json:"InverterActivePower"`
}