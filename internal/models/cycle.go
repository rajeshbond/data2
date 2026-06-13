package models

type CycleRequest struct {
	DeviceID        string  `json:"device_id"`
	CycleTimeSec    float64 `json:"cycle_time_sec"`
	ProductionCount uint64  `json:"production_count"`
}
