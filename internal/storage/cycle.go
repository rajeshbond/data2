package storage

import (
	"sync"
	"time"

	"github.com/rajbond/data2/internal/models"
)

var (
	Mu sync.RWMutex

	LatestCycle models.CycleRequest

	LatestGapMs int64

	LatestServerTime string

	LatestHandlerDurationUs int64

	LastRequestTime = make(
		map[string]time.Time,
	)
)

// package storage

// import (
// 	"sync"
// 	"time"

// 	"github.com/rajbond/data2/internal/models"
// )

// var (
// 	Mu sync.RWMutex

// 	LatestCycle models.CycleRequest

// 	LastRequestTime = make(
// 		map[string]time.Time,
// 	)
// )

// package storage

// import (
// 	"sync"

// 	"github.com/rajbond/data2/internal/models"
// )

// var (
// 	LatestCycle models.CycleRequest
// 	Mu          sync.RWMutex
// )
