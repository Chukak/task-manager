package tasks

import (
	"log"
	"sync/atomic"
	"time"

	db "../database"
	"github.com/chukak/task-manager/pkg/timers"
)

// TaskManage is a common interface for tasks
type TaskManage interface {
	AddSubtask(*Task)
	RemoveSubtask(*Task)
	CountSubtasks() int
	Open(bool)
	SetActive(bool)
	RemoveSelf()
	UpdateInDb()
}

// TaskDuration store a duration of task
type TaskDuration struct {
	Seconds int `json:"seconds"`
	Minutes int `json:"minutes"`
	Hours   int `json:"hours"`
	Days    int `json:"days"`
}

// Task stores targets, time, duration and all dependencies of this target
type Task struct {
	parent      *Task
	ticker      timers.CountdownTimer
	Start       time.Time    `json:"start"`
	End         time.Time    `json:"end"`
	Duration    TaskDuration `json:"duration"`
	Subtasks    []*Task      `json:"subtasks"`
	running     int32
	IsActive    bool   `json:"active"`
	IsOpened    bool   `json:"opened"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
	TaskID      int64  `json:"taskID"`
}

// ListTaskManage is a common interface for task lists
type ListTaskManage interface {
	Append(*Task)
	Remove(*Task)
}

// ListTask stores all tasks
type ListTask struct {
	List []*Task `json:"listTasks"`
}

var databasePointer *db.Database = nil
var listTaskPointer *ListTask = nil
var taskPointers = map[int64]*Task{}

// SetDatabase sets the database pointer
func SetDatabase(d *db.Database) {
	databasePointer = d
}

// execSql exes SQL query
func execSQL(t db.QueryExecType, q string, args ...interface{}) (db.QueryResult, error) {
	var rows db.QueryResult
	var err error
	if databasePointer != nil {
		conn, err := databasePointer.GetConnection()
		rows, err = databasePointer.Exec(t, conn, q, args...)
		if err != nil {
			log.Println("SQL error: ", err.Error())
		}
		databasePointer.CloseConnection(conn)
	}
	return rows, err
}

// NewTask returns a new Task object
func NewTask(par *Task) *Task {
	task := Task{
		parent: par, Start: time.Time{}, End: time.Time{}, ticker: timers.NewCountdownTimer(),
		running: 0, Subtasks: []*Task{},
		Duration: TaskDuration{Seconds: 0, Minutes: 0, Hours: 0, Days: 0},
		IsActive: false, IsOpened: false, Title: "", Description: "", Priority: 0, TaskID: -1}

	if listTaskPointer != nil && par == nil {
		listTaskPointer.Append(&task)
	}

	if databasePointer != nil {
		result, err := execSQL(db.SELECT, "INSERT INTO task_duration DEFAULT VALUES RETURNING id;")
		if err == nil && result.Next() {
			var durationID int = -1
			result.Scan(&durationID)
			result, _ = execSQL(db.SELECT,
				"INSERT INTO tasks (duration_id) VALUES ($1) RETURNING id, start_time, end_time", durationID)
			if err == nil && result.Next() {
				var taskID int64 = -1
				result.Scan(&taskID, &task.Start, &task.End)
				task.TaskID = taskID
				taskPointers[taskID] = &task
			}
		}
	}

	if par != nil {
		par.AddSubtask(&task)
	}
	return &task
}

// AddSubtask adds a subtask to this Task. This task become a parent of added subtask
func (t *Task) AddSubtask(newSubtask *Task) {
	newSubtask.parent = t
	t.Subtasks = append(t.Subtasks, newSubtask)
	execSQL(db.UPDATE,
		"UPDATE tasks SET parent_id = $1 WHERE id = $2 RETURNING id;", t.TaskID, newSubtask.TaskID)
}

// RemoveSubtask removes a subtask from this Task,
// also removes all the subtasks from this subtask
func (t *Task) RemoveSubtask(oldSubtask *Task) {
	oldSubtask.Subtasks = nil
	oldSubtask.parent = nil
	index := -1
	for i, e := range t.Subtasks {
		if e == oldSubtask {
			index = i

		}
	}
	if index > -1 {
		t.Subtasks = append(t.Subtasks[:index], t.Subtasks[index+1:]...)
	}

	execSQL(db.DELETE, "DELETE FROM tasks WHERE id = $1;", oldSubtask.TaskID)
}

// CountSubtasks count number of subtasks
func (t *Task) CountSubtasks() int {
	return len(t.Subtasks)
}

// Open this task
func (t *Task) Open(o bool) {
	t.IsOpened = o
	if !o {
		t.SetActive(false)
	}
	execSQL(db.UPDATE, "UPDATE tasks SET is_open = $1 WHERE id = $2;", o, t.TaskID)
}

// SetActive set this task as active
func (t *Task) SetActive(active bool) {
	if active != t.IsActive {
		t.IsActive = active
		execSQL(db.UPDATE, "UPDATE tasks SET is_active = $1 WHERE id = $2;", active, t.TaskID)

		if active {
			if !t.IsOpened {
				t.Open(active)
			}
			atomic.StoreInt32(&t.running, 1)
			t.Start = time.Now().Truncate(time.Second)
			execSQL(db.UPDATE, "UPDATE tasks SET start_time = $1 WHERE id = $2;", t.Start, t.TaskID)
			t.ticker = timers.NewCountdownTimer()
			t.ticker.Run()

			go func() {
				for atomic.LoadInt32(&t.running) > 0 {
					select {
					case <-t.ticker.Tick:
						t.Duration.Seconds = t.ticker.Sec
						t.Duration.Minutes = t.ticker.Min
						t.Duration.Hours = t.ticker.Hours
						t.Duration.Days = t.ticker.Days
						execSQL(db.UPDATE,
							`UPDATE task_duration SET second = $1, minute = $2, hour = $3, day = $4 
							WHERE id IN (SELECT duration_id FROM tasks WHERE id = $5);`,
							t.Duration.Seconds, t.Duration.Minutes, t.Duration.Hours, t.Duration.Days, t.TaskID)
					}
				}
			}()
		} else {
			atomic.StoreInt32(&t.running, 0)
			t.ticker.Finish()
			t.End = time.Now().Truncate(time.Second)
			execSQL(db.UPDATE, "UPDATE tasks SET end_time = $1 WHERE id = $2;", t.End, t.TaskID)
		}
	}
}

// RemoveSelf removes this task
func (t *Task) RemoveSelf() {
	t.Subtasks = nil
	t.parent = nil

	execSQL(db.DELETE, "DELETE FROM tasks WHERE id = $1 OR parent_id = $1;", t.TaskID)
}

func (t *Task) UpdateInDb() {
	execSQL(db.UPDATE,
		"UPDATE tasks SET title = $1, descr = $2, priority = $3 WHERE id = $4;",
		t.Title, t.Description, t.Priority, t.TaskID)
}

// NewListTask returns a new ListTask object
func NewListTask() *ListTask {
	listTaskPointer = &ListTask{}
	return listTaskPointer
}

// Append a new Task to list
func (l *ListTask) Append(t *Task) {
	l.List = append(l.List, t)
}

// Remove a task from list
func (l *ListTask) Remove(t *Task) {
	index := -1
	for i, e := range l.List {
		if e == t {
			index = i
		}
	}
	if index > -1 {
		l.List = append(l.List[:index], l.List[index+1:]...)
	}
}
