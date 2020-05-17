package timers

import (
	"math"
	"testing"
	"time"

	test "github.com/chukak/task-manager/pkg/checks"
	utility "github.com/chukak/task-manager/pkg/utility"
)

const EPSILON float64 = 1.0

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

func TestDeadlineTimer(t *testing.T) {
	test.SetT(t)
	{
		deadline := NewDeadlineTimer()
		test.CheckFalse(deadline.IsRunning())

		interval := time.Millisecond * 800
		resultValue, expectedValue := 0, 5

		SetValueFn1 := func(v int) {
			resultValue = v
		}
		SetValueFn2 := func(k1 int, k2 int, k3 int) {
			resultValue = k1 + k2 + k3
		}
		SetValueFn3 := func() {
			resultValue = expectedValue
		}

		deadline.ExpiresFromNow(interval, utility.Bind(SetValueFn1, expectedValue))
		test.CheckTrue(deadline.IsRunning())
		test.CheckNotEqual(resultValue, expectedValue)

		time.Sleep(time.Second * 1)
		test.CheckFalse(deadline.IsRunning())
		test.CheckEqual(deadline.CallTimes, 1)

		test.CheckEqual(resultValue, expectedValue)
		resultValue = 0

		deadline.ExpiresFromNow(interval, utility.Bind(SetValueFn2, 1, 7, -3))
		test.CheckTrue(deadline.IsRunning())
		test.CheckNotEqual(resultValue, expectedValue)

		time.Sleep(time.Second * 1)
		test.CheckFalse(deadline.IsRunning())

		test.CheckEqual(resultValue, expectedValue)
		test.CheckEqual(deadline.CallTimes, 2)
		resultValue = 0

		deadline.ExpiresFromNow(interval, utility.Bind(SetValueFn3))
		test.CheckTrue(deadline.IsRunning())
		test.CheckNotEqual(resultValue, expectedValue)

		time.Sleep(time.Second * 1)
		test.CheckFalse(deadline.IsRunning())

		test.CheckEqual(resultValue, expectedValue)
		test.CheckEqual(deadline.CallTimes, 3)

		deadline.Cancel()
	}
	{
		deadline := NewDeadlineTimer()
		deadline.AsLoop(true)
		test.CheckFalse(deadline.IsRunning())
		test.CheckTrue(deadline.IsLoop())

		interval := time.Millisecond * 800
		resultValue, expectedValue := 0, 5

		SetValueFn1 := func(v int) {
			resultValue = v
		}

		deadline.ExpiresFromNow(interval, utility.Bind(SetValueFn1, expectedValue))
		for i := 0; i < 17; i++ {
			test.CheckTrue(deadline.IsRunning())
			test.CheckNotEqual(resultValue, expectedValue)

			time.Sleep(time.Second * 1)
			test.CheckTrue(deadline.IsRunning())

			test.CheckEqual(resultValue, expectedValue)
			resultValue = 0
		}
		test.CheckTrue(deadline.CallTimes > 17)
		test.CheckTrue(deadline.IsRunning())
		deadline.Cancel()
		test.CheckFalse(deadline.IsRunning())
		test.CheckEqual(deadline.CallTimes, 0)
	}
}

func TestCountdownTimer(t *testing.T) {
	test.SetT(t)

	countdownTimer := NewCountdownTimer()

	countdownTimer.Run()
	for i := 1; i < 65; i++ {
		sec := i % 60

		min := i / 60 % 60
		select {
		case <-countdownTimer.Tick:
			// check seconds with epsilon
			test.CheckTrue((countdownTimer.Sec == sec) ||
				(math.Abs(float64(countdownTimer.Sec-sec)) == EPSILON))

			test.CheckEqual(countdownTimer.Min, min)
			test.CheckEqual(countdownTimer.Hours, 0)
			test.CheckEqual(countdownTimer.Days, 0)
		}
	}
	countdownTimer.Finish()
}
