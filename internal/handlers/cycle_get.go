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
	data := storage.LatestCycle
	storage.Mu.RUnlock()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(data)
}
