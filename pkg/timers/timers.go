package timers

import (
	"time"
)

// TimeManage is a common timer interface.
type TimeManage interface {
	SetIntervalMs(int64)
	Start()
	Stop()
	Cancel()
}

// ElapsedTimer counts elapsed time from start point and save this.
type ElapsedTimer struct {
	interval int64
	start    time.Time
	elapsed  time.Duration
	Microsec int64
	Millisec int64
	Sec      float64
}

// NewElapsedTimer object returns from this function
func NewElapsedTimer() ElapsedTimer {
	return ElapsedTimer{Microsec: 0, Millisec: 0, Sec: 0}
}

// Start the timer from now.
func (timer *ElapsedTimer) Start() {
	timer.start = time.Now()
}

// Stop the timer and save elapsed time from start point.
func (timer *ElapsedTimer) Stop() {
	timer.elapsed = time.Since(timer.start)
	timer.Microsec = timer.elapsed.Microseconds()
	timer.Millisec = timer.elapsed.Milliseconds()
	timer.Sec = timer.elapsed.Seconds()
}
