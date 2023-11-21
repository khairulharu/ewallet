package dto

type IpChecker struct {
	Status   string  `json:"status"`
	Query    string  `json:"query"`
	Timezone string  `json:"timezone"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}
