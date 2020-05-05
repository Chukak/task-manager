package tasks

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	test "github.com/chukak/task-manager/pkg/checks"
)

func TestTaskInitialization(t *testing.T) {
	test.SetT(t)

	task := NewTask(nil)

	test.CheckEqual(task.Description, "")
	test.CheckEqual(task.Priority, -1)
	test.CheckEqual(task.Title, "")
	test.CheckEqual(task.parent, (*Task)(nil))
	test.CheckEqual(task.running, 0)
	test.CheckEqual(len(task.subtasks), 0)
	test.CheckFalse(task.IsActive)
	test.CheckFalse(task.IsOpened)
}

func TestTaskFunctionality(t *testing.T) {
	test.SetT(t)

	task := NewTask(nil)

	task.Open(true)
	test.CheckTrue(task.IsOpened)
	test.CheckFalse(task.IsActive)
	test.CheckEqual(task.running, 0)

	ticker := time.NewTicker(time.Second * 1)
	task.SetActive(true)
	start := time.Now()
	for i := 0; i < 6; i++ {
		tick := <-ticker.C
		diff := tick.Sub(start)
		sec := int(diff.Seconds()) % 60
		// delay
		test.CheckTrue(sec == task.Duration.Seconds ||
			1 == math.Abs(float64(sec-task.Duration.Seconds)))
	}
	test.CheckTrue(task.IsOpened)
	test.CheckTrue(task.IsActive)
	test.CheckEqual(task.running, 1)
	test.CheckEqual(task.Start.Year(), start.Year())
	test.CheckEqual(int(task.Start.Month()), int(start.Month()))
	test.CheckEqual(task.Start.Day(), start.Day())
	test.CheckEqual(task.Start.Hour(), start.Hour())
	test.CheckEqual(task.Start.Minute(), start.Minute())
	test.CheckEqual(task.Start.Second(), start.Second())

	task.Open(false)
	test.CheckFalse(task.IsOpened)
	test.CheckFalse(task.IsActive)
	test.CheckEqual(task.running, 0)
	end := time.Now()
	test.CheckEqual(task.End.Year(), end.Year())
	test.CheckEqual(int(task.End.Month()), int(end.Month()))
	test.CheckEqual(task.End.Day(), end.Day())
	test.CheckEqual(task.End.Hour(), end.Hour())
	test.CheckEqual(task.End.Minute(), end.Minute())
	test.CheckEqual(task.End.Second(), end.Second())
}

func TestTaskWithSubtasks(t *testing.T) {
	test.SetT(t)
	{
		task := NewTask(nil)

		subtask1 := NewTask(task)
		test.CheckEqual(task.CountSubtasks(), 1)
		test.CheckEqual(len(task.Subtasks()), 1)
		test.CheckEqual(task.Subtasks()[0], subtask1)
		test.CheckEqual(subtask1.parent, task)
	}
	{
		task := NewTask(nil)
		var cached [10]*Task
		for i := 0; i < 10; i++ {
			cached[i] = NewTask(nil)
			task.AddSubtask(cached[i])
			test.CheckEqual(cached[i].parent, task)
			test.CheckEqual(task.CountSubtasks(), i+1)
			cached[i].AddSubtask(NewTask(nil))
			test.CheckEqual(cached[i].CountSubtasks(), 1)
		}

		allSubtasks := task.Subtasks()
		for i := 0; i < 10; i++ {
			test.CheckEqual(cached[i], allSubtasks[i])
		}

		for i := 0; i < 10; i++ {
			task.RemoveSubtask(cached[i])
			test.CheckEqual(cached[i].parent, (*Task)(nil))
			test.CheckEqual(cached[i].CountSubtasks(), 0)
			test.CheckEqual(task.CountSubtasks(), 10-i-1)
		}
	}
}

func TestTaskToJson(t *testing.T) {
	test.SetT(t)

	task := NewTask(nil)
	task.Description = "New task desc!"
	task.Title = "Task 1"
	task.Priority = 3

	task.SetActive(true)
	time.Sleep(time.Millisecond * 10)
	ticker := time.NewTicker(time.Second * 1)
	startTask := task.Start
	for i := 0; i < 3; i++ {
		<-ticker.C
	}
	task.SetActive(false)
	endTask := task.End

	var data []byte
	data, err := json.Marshal(task)
	test.CheckEqual(err, nil)
	test.CheckTrue(len(data) > 0)

	var values map[string]json.RawMessage
	err = json.Unmarshal(data, &values)
	test.CheckEqual(err, nil)
	test.CheckTrue(len(data) > 0)
	test.CheckEqual(string(values["description"]), "\"New task desc!\"")
	test.CheckEqual(string(values["title"]), "\"Task 1\"")

	var priority int = 0
	test.CheckEqual(json.Unmarshal(values["priority"], &priority), nil)
	test.CheckEqual(priority, 3)

	var opened bool = false
	test.CheckEqual(json.Unmarshal(values["opened"], &opened), nil)
	test.CheckTrue(opened)

	var active bool = false
	test.CheckEqual(json.Unmarshal(values["active"], &active), nil)
	test.CheckFalse(active)

	var duration TaskDuration
	test.CheckEqual(json.Unmarshal(values["duration"], &duration), nil)
	test.CheckEqual(duration.Seconds, 3)
	test.CheckEqual(duration.Minutes, 0)
	test.CheckEqual(duration.Hours, 0)
	test.CheckEqual(duration.Days, 0)

	var startTaskUnmarshal, endTaskUnmarshal time.Time
	test.CheckEqual(json.Unmarshal(values["start"], &startTaskUnmarshal), nil)
	test.CheckEqual(json.Unmarshal(values["end"], &endTaskUnmarshal), nil)
	// string, because checks package have not == operator for time.Time
	test.CheckEqual(startTaskUnmarshal.UTC().String(), startTask.UTC().Round(0).String())
	test.CheckEqual(endTaskUnmarshal.UTC().String(), endTask.UTC().Round(0).String())
}
