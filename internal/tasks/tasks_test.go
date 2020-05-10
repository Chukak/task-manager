package tasks

import (
	"encoding/json"
	"math"
	"os"
	"strconv"
	"testing"
	"time"

	db "../database"
	test "github.com/chukak/task-manager/pkg/checks"
)

func TestTaskInitialization(t *testing.T) {
	test.SetT(t)

	task := NewTask(nil)

	test.CheckEqual(task.Description, "")
	test.CheckEqual(task.Priority, 0)
	test.CheckEqual(task.Title, "")
	test.CheckEqual(task.parent, (*Task)(nil))
	test.CheckEqual(task.running, 0)
	test.CheckEqual(len(task.Subtasks), 0)
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
		test.CheckEqual(len(task.Subtasks), 1)
		test.CheckEqual(task.Subtasks[0], subtask1)
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

		allSubtasks := task.Subtasks
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

func TestTaskList(t *testing.T) {
	test.SetT(t)

	task := NewTask(nil)
	task.Title = "Task 1"
	task.Description = "This is task description"
	task.Priority = 4

	task2 := NewTask(nil)
	task2.Title = "Task 2"
	task2.Description = "This is task 2 description"
	task2.Priority = 1

	task3 := NewTask(nil)
	task3.Title = "Task 3"
	task3.Description = "This is task 3 description"
	task3.Priority = 2

	listTask := NewListTask()
	test.CheckNotEqual(listTask, nil)

	listTask.Append(task)
	listTask.Append(task2)
	listTask.Append(task3)
	test.CheckEqual(len(listTask.List), 3)

	listTask.Remove(task2)
	test.CheckEqual(len(listTask.List), 2)
	test.CheckEqual(listTask.List[0], task)
	test.CheckEqual(listTask.List[1], task3)

	listTask.Remove(task)
	listTask.Remove(task3)
	test.CheckEqual(len(listTask.List), 0)
}

func TestTaskFunctionalityUsingDatabase(t *testing.T) {
	test.SetT(t)

	host := os.Getenv("DB_HOST")
	val, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	port := uint16(val)
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	d, _ := db.NewDatabase(host, uint16(port), database, user, password)
	ok, err := d.Open()
	test.CheckTrue(ok)

	SetDatabase(d)
	task := NewTask(nil)
	test.CheckEqual(len(listTaskPointer.List), 1)
	test.CheckEqual(len(taskPointers), 1)
	test.CheckEqual(listTaskPointer.List[0], task)
	test.CheckEqual(taskPointers[task.TaskID], task)

	rows, err := d.Exec(`SELECT 
		parent_id, start_time, end_time, duration_id, is_open, is_active, title, descr, priority 
		FROM tasks WHERE id = $1;`, task.TaskID)
	test.CheckEqual(err, nil)
	test.CheckTrue(rows.Next())
	{
		var parentID int64
		var startTime time.Time
		var endTime time.Time
		var durationID int64
		var isOpen bool
		var isActive bool
		var title string
		var descr string
		var priority int

		test.CheckEqual(rows.Scan(&parentID, &startTime, &endTime, &durationID, &isOpen,
			&isActive, &title, &descr, &priority), nil)

		test.CheckEqual(parentID, -1)
		// string, because checks package have not == operator for time.Time
		test.CheckEqual(startTime.String(), task.Start.String())
		test.CheckEqual(endTime.String(), task.End.String())
		test.CheckNotEqual(durationID, -1)
		test.CheckEqual(isOpen, task.IsOpened)
		test.CheckEqual(isActive, task.IsActive)
		test.CheckEqual(title, task.Title)
		test.CheckEqual(descr, task.Description)
		test.CheckEqual(priority, task.Priority)
	}

	task.Open(true)
	time.Sleep(800 * time.Millisecond)
	rows, err = d.Exec(`SELECT is_open FROM tasks WHERE id = $1;`, task.TaskID)
	test.CheckEqual(err, nil)
	test.CheckTrue(rows.Next())
	{
		var isOpen bool
		test.CheckEqual(rows.Scan(&isOpen), nil)
		test.CheckTrue(isOpen)
	}

	task.SetActive(true)
	time.Sleep(4 * time.Second)

	rows, err = d.Exec(`SELECT is_active, duration_id, start_time FROM tasks WHERE id = $1;`, task.TaskID)
	test.CheckEqual(err, nil)
	test.CheckTrue(rows.Next())
	{
		var isActive bool
		var durationID int64 = -1
		var startTime time.Time
		test.CheckEqual(rows.Scan(&isActive, &durationID, &startTime), nil)
		test.CheckTrue(isActive)
		test.CheckNotEqual(durationID, -1)
		// string, because checks package have not == operator for time.Time
		test.CheckEqual(startTime.String(), task.Start.String())

		rows, err = d.Exec(`SELECT second, minute, hour, day FROM task_duration WHERE id = $1;`,
			durationID)
		test.CheckEqual(err, nil)
		test.CheckTrue(rows.Next())
		{
			var second, minute, hour, day int
			test.CheckEqual(rows.Scan(&second, &minute, &hour, &day), nil)

			test.CheckNotEqual(second, 0)
			test.CheckEqual(minute, task.Duration.Minutes)
			test.CheckEqual(hour, task.Duration.Hours)
			test.CheckEqual(day, task.Duration.Days)
		}
	}

	task.Open(false)
	time.Sleep(800 * time.Millisecond)
	rows, err = d.Exec(`SELECT is_active, is_open, end_time FROM tasks WHERE id = $1;`, task.TaskID)
	test.CheckEqual(err, nil)
	test.CheckTrue(rows.Next())
	{
		var isActive bool
		var isOpen bool
		var endTime time.Time
		test.CheckEqual(rows.Scan(&isActive, &isOpen, &endTime), nil)
		test.CheckFalse(isActive)
		test.CheckFalse(isOpen)
		// string, because checks package have not == operator for time.Time
		test.CheckEqual(endTime.String(), task.End.Truncate(time.Second).String())
	}

	task.Title = "Title"
	task.Description = "Description"
	task.Priority = 1
	task.UpdateInDb()
	time.Sleep(800 * time.Millisecond)
	rows, err = d.Exec(`SELECT title, descr, priority FROM tasks WHERE id = $1;`, task.TaskID)
	test.CheckEqual(err, nil)
	test.CheckTrue(rows.Next())
	{
		var title, description string
		var priority int8
		test.CheckEqual(rows.Scan(&title, &description, &priority), nil)
		test.CheckEqual(title, task.Title)
		test.CheckEqual(description, task.Description)
		test.CheckEqual(priority, task.Priority)
	}

	task.RemoveSelf()
	rows, err = d.Exec(`SELECT id FROM tasks WHERE id = $1;`, task.TaskID)
	test.CheckEqual(err, nil)
	test.CheckFalse(rows.Next())
}

func TestTaskSubtasksUsingDatabase(t *testing.T) {
	test.SetT(t)

	host := os.Getenv("DB_HOST")
	val, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	port := uint16(val)
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	d, _ := db.NewDatabase(host, uint16(port), database, user, password)
	ok, err := d.Open()
	test.CheckTrue(ok)

	SetDatabase(d)
	task := NewTask(nil)

	subtask1 := NewTask(task)
	subtask2 := NewTask(task)
	subtask3 := NewTask(task)

	rows, err := d.Exec(`SELECT id FROM tasks WHERE parent_id = $1;`, task.TaskID)
	test.CheckEqual(err, nil)
	test.CheckTrue(rows.Next())
	{
		var taskID int64
		test.CheckEqual(rows.Scan(&taskID), nil)
		test.CheckEqual(taskID, subtask1.TaskID)
	}
	test.CheckTrue(rows.Next())
	{
		var taskID int64
		test.CheckEqual(rows.Scan(&taskID), nil)
		test.CheckEqual(taskID, subtask2.TaskID)
	}
	test.CheckTrue(rows.Next())
	{
		var taskID int64
		test.CheckEqual(rows.Scan(&taskID), nil)
		test.CheckEqual(taskID, subtask3.TaskID)
	}
	test.CheckFalse(rows.Next())
}
