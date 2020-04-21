package lib

type GetIntelDataResponse struct {
	Event string `json:"event"`
	Report Report `json:"report"`
}

type Report struct {
	Stats []ReportDataPoint `json:"stats"`
}

type ReportDataPoint struct {
	Tick int `json:"tick"`
	Players []PlayerIntel `json:"players"`
}