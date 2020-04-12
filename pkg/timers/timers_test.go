package timers

import (
	"testing"
	"time"

	test "github.com/chukak/task-manager/pkg/test"
	utility "github.com/chukak/task-manager/pkg/utility"
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

		time.Sleep(time.Second * 1)
		test.CheckFalse(deadline.IsRunning())

		test.CheckEqual(resultValue, expectedValue)

		deadline.ExpiresFromNow(interval, utility.Bind(SetValueFn2, 1, 7, -3))
		test.CheckTrue(deadline.IsRunning())

		time.Sleep(time.Second * 1)
		test.CheckFalse(deadline.IsRunning())

		test.CheckEqual(resultValue, expectedValue)

		deadline.ExpiresFromNow(interval, utility.Bind(SetValueFn3))
		test.CheckTrue(deadline.IsRunning())

		time.Sleep(time.Second * 1)
		test.CheckFalse(deadline.IsRunning())

		test.CheckEqual(resultValue, expectedValue)

		deadline.Cancel()
	}
}
