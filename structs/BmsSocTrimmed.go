package structs

import "time"

// BmsSocTrimmed struct representing BmsSocTrimmed last 1
type BmsSocTrimmed struct {
	UTCTime       time.Time `json:"UtcTime"`
	BmsSocTrimmed int64     `json:"BmsSocTrimmed"`
}
