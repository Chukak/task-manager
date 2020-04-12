package timers

import (
	"sync/atomic"
	"time"

	utility "github.com/chukak/task-manager/pkg/utility"
)

// ElapsedTimerManage is a common interface for elapsed timers.
type ElapsedTimerManage interface {
	Start()
	Stop()
}

// ElapsedTimer counts elapsed time from start point and save this.
type ElapsedTimer struct {
	start    time.Time
	elapsed  time.Duration
	Microsec int64
	Millisec int64
	Sec      float64
}

// NewElapsedTimer returns a new ElapsedTimer object.
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

// DeadlineTimerManage is a common interface for deadline timers.
type DeadlineTimerManage interface {
	ExpiresFromNow(time.Duration, utility.BindedFunction)
	Cancel()
	AsLoop(bool)
	IsRunning()
	IsLoop()
}

// DeadlineTimer calls function after expired time.
type DeadlineTimer struct {
	t         *time.Timer
	fun       utility.BindedFunction
	duration  time.Duration
	loop      bool
	running   int32
	canceled  chan bool
	CallTimes int
}

// NewDeadlineTimer returns a new DeadlineTimer object.
func NewDeadlineTimer() DeadlineTimer {
	return DeadlineTimer{CallTimes: 0, loop: false, running: 0}
}

// ExpiresFromNow calls function after expired time.
func (dtimer *DeadlineTimer) ExpiresFromNow(interval time.Duration, fn utility.BindedFunction) {
	if atomic.LoadInt32(&dtimer.running) == 0 {
		dtimer.t = time.NewTimer(interval)
		dtimer.fun = fn
		atomic.StoreInt32(&dtimer.running, 1)
		dtimer.canceled = make(chan bool)
	}

	go func() {
		for atomic.LoadInt32(&dtimer.running) > 0 {
			select {
			case <-dtimer.t.C:
				dtimer.fun.F()
				dtimer.CallTimes++
				dtimer.t.Reset(interval)
				if !dtimer.loop {
					atomic.StoreInt32(&dtimer.running, 0)
					break
				}
			case <-dtimer.canceled:
			}
		}
	}()
}

// Cancel timer operations.
func (dtimer *DeadlineTimer) Cancel() bool {
	ok := false
	if atomic.LoadInt32(&dtimer.running) > 0 {
		dtimer.canceled <- true
		ok = dtimer.t.Stop()
	}
	atomic.StoreInt32(&dtimer.running, 0)
	close(dtimer.canceled)
	dtimer.CallTimes = 0

	return ok
}

// AsLoop sets deadline timer as cyclic timer or not.
func (dtimer *DeadlineTimer) AsLoop(on bool) {
	dtimer.loop = on
}

// IsRunning returns running status of deadline timer.
func (dtimer *DeadlineTimer) IsRunning() bool {
	return atomic.LoadInt32(&dtimer.running) > 0
}

// IsLoop returns loop status of deadline timer
func (dtimer *DeadlineTimer) IsLoop() bool {
	return dtimer.loop
}
