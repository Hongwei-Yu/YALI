package engine

import "time"

type Launcher interface {
	send() (bool, int64, uint64, float64, float64, string, time.Time, time.Time)
}
