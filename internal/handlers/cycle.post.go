package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rajbond/data2/internal/models"
	"github.com/rajbond/data2/internal/storage"
)

func CreateCycle(
	w http.ResponseWriter,
	r *http.Request,
) {

	start := time.Now()

	defer func() {

		fmt.Printf(
			"Handler Duration: %d µs\n",
			time.Since(start).Microseconds(),
		)

		fmt.Println()
	}()

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"POST only",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var req models.CycleRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	now := time.Now()

	var gapMs int64 = 0

	storage.Mu.Lock()

	lastTime, exists :=
		storage.LastRequestTime[req.DeviceID]

	if exists {

		gapMs =
			now.Sub(lastTime).
				Milliseconds()
	}

	storage.LastRequestTime[req.DeviceID] = now

	storage.LatestCycle = req

	storage.Mu.Unlock()

	fmt.Println(
		"--------------------------------",
	)

	fmt.Printf(
		"Time: %s\n",
		now.Format(
			"2006-01-02 15:04:05.000",
		),
	)

	fmt.Printf(
		"Device: %s\n",
		req.DeviceID,
	)

	fmt.Printf(
		"Cycle Time: %.2f sec\n",
		req.CycleTimeSec,
	)

	fmt.Printf(
		"Production Count: %d\n",
		req.ProductionCount,
	)

	if gapMs > 0 {

		fmt.Printf(
			"Gap Since Previous Request: %d ms (%.2f sec)\n",
			gapMs,
			float64(gapMs)/1000,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]any{
			"success": true,
		},
	)
}

// package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/rajbond/data2/internal/models"
// 	"github.com/rajbond/data2/internal/storage"
// )

// func CreateCycle(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {

// 	start := time.Now()

// 	defer func() {
// 		fmt.Printf(
// 			"Handler Duration: %d µs\n",
// 			time.Since(start).Microseconds(),
// 		)
// 	}()

// 	if r.Method != http.MethodPost {

// 		http.Error(
// 			w,
// 			"POST only",
// 			http.StatusMethodNotAllowed,
// 		)

// 		return
// 	}

// 	var req models.CycleRequest

// 	err := json.NewDecoder(
// 		r.Body,
// 	).Decode(&req)

// 	if err != nil {

// 		http.Error(
// 			w,
// 			err.Error(),
// 			http.StatusBadRequest,
// 		)

// 		return
// 	}

// 	// Calculate network/device gap
// 	receivedAt := time.Now()

// 	gap := receivedAt.UnixMilli() - req.Timestamp

// 	storage.Mu.Lock()
// 	storage.LatestCycle = req
// 	storage.Mu.Unlock()

// 	fmt.Println("--------------------------------")

// 	fmt.Printf(
// 		"Received At: %s\n",
// 		receivedAt.Format("2006-01-02 15:04:05.000"),
// 	)

// 	fmt.Printf(
// 		"Device: %s\n",
// 		req.DeviceID,
// 	)

// 	fmt.Printf(
// 		"Cycle Time: %.2f sec\n",
// 		req.CycleTimeSec,
// 	)

// 	fmt.Printf(
// 		"Production Count: %d\n",
// 		req.ProductionCount,
// 	)

// 	fmt.Printf(
// 		"Network Gap: %d ms\n",
// 		gap,
// 	)

// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	json.NewEncoder(w).Encode(
// 		map[string]any{
// 			"success": true,
// 		},
// 	)
// }

// func CreateCycle(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {

// 	if r.Method != http.MethodPost {

// 		http.Error(
// 			w,
// 			"POST only",
// 			http.StatusMethodNotAllowed,
// 		)

// 		return
// 	}

// 	var req models.CycleRequest

// 	err := json.NewDecoder(
// 		r.Body,
// 	).Decode(&req)

// 	if err != nil {

// 		http.Error(
// 			w,
// 			err.Error(),
// 			http.StatusBadRequest,
// 		)

// 		return
// 	}

// 	storage.Mu.Lock()
// 	storage.LatestCycle = req
// 	storage.Mu.Unlock()

// 	// storage.Mu.Lock()
// 	// storage.LatestCycle = req
// 	// storage.Mu.Unlock()

// 	fmt.Println("--------------------------------")

// 	fmt.Printf(
// 		"Time: %s\n",
// 		time.Now().Format(
// 			"2006-01-02 15:04:05",
// 		),
// 	)

// 	fmt.Printf(
// 		"Device: %s\n",
// 		req.DeviceID,
// 	)

// 	fmt.Printf(
// 		"Cycle Time: %.2f sec\n",
// 		req.CycleTimeSec,
// 	)

// 	fmt.Printf(
// 		"Production Count: %d\n",
// 		req.ProductionCount,
// 	)

// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	json.NewEncoder(w).Encode(
// 		map[string]any{
// 			"success": true,
// 		},
// 	)
// }
