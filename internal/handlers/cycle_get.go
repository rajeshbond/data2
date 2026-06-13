package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rajbond/data2/internal/storage"
)

func GetLatestCycle(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"GET only",
			http.StatusMethodNotAllowed,
		)

		return
	}

	storage.Mu.RLock()

	response := map[string]any{
		"device_id": storage.LatestCycle.DeviceID,

		"cycle_time_sec": storage.LatestCycle.CycleTimeSec,

		"production_count": storage.LatestCycle.ProductionCount,

		"server_time": storage.LatestServerTime,

		"gap_ms": storage.LatestGapMs,

		"gap_sec": float64(storage.LatestGapMs) / 1000,

		"handler_duration_us": storage.LatestHandlerDurationUs,
	}

	storage.Mu.RUnlock()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}

// package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/rajbond/data2/internal/storage"
// )

// func GetLatestCycle(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {

// 	if r.Method != http.MethodGet {

// 		http.Error(
// 			w,
// 			"GET only",
// 			http.StatusMethodNotAllowed,
// 		)

// 		return
// 	}

// 	storage.Mu.RLock()
// 	data := storage.LatestCycle
// 	storage.Mu.RUnlock()

// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	json.NewEncoder(w).Encode(data)
// }
