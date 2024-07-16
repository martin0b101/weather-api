package weather

import (
	"sync"
	"time"
	"github.com/weather-api/types"
)


var (
	cacheMutex sync.RWMutex
	cache      *[]types.City
)

func init() {
	cache = nil
	go func() {
		for {
			time.Sleep(20 * time.Minute) 
		}
	}()
}