package timers

import (
	"testing"
	"time"
)

func TestElapsedTimer(t *testing.T) {
	elapsed := NewElapsedTimer()
	interval := 2

	go elapsed.Start()
	time.Sleep(time.Second * time.Duration(interval))
	elapsed.Stop()
	if int(elapsed.Sec) != interval {
		t.Errorf("elapsed = %v; want: %v", int64(elapsed.Sec),
			interval)
	}
}
