package tasks

import (
	"sync/atomic"
	"time"

	"github.com/chukak/task-manager/pkg/timers"
)

// TaskManage is a common interface for tasks
type TaskManage interface {
	AddSubtask(*Task)
	RemoveSubtask(*Task)
	CountSubtasks() int
	Open(bool)
	SetActive(bool)
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

// NewTask returns a new Task object
func NewTask(par *Task) *Task {
	task := Task{
		parent: par, Start: time.Time{}, End: time.Time{}, ticker: timers.NewCountdownTimer(),
		running: 0, Subtasks: []*Task{},
		Duration: TaskDuration{Seconds: 0, Minutes: 0, Hours: 0, Days: 0},
		IsActive: false, IsOpened: false, Title: "", Description: "", Priority: -1}
	if par != nil {
		par.AddSubtask(&task)
	}
	return &task
}

// AddSubtask adds a subtask to this Task. This task become a parent of added subtask
func (t *Task) AddSubtask(newSubtask *Task) {
	newSubtask.parent = t
	t.Subtasks = append(t.Subtasks, newSubtask)
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
}

// SetActive set this task as active
func (t *Task) SetActive(active bool) {
	if active != t.IsActive {
		t.IsActive = active
		if active {
			if !t.IsOpened {
				t.Open(active)
			}
			atomic.StoreInt32(&t.running, 1)
			t.Start = time.Now()
			t.ticker.Run()

			go func() {
				for atomic.LoadInt32(&t.running) > 0 {
					select {
					case <-t.ticker.Tick:
						t.Duration.Seconds = t.ticker.Sec
						t.Duration.Minutes = t.ticker.Min
						t.Duration.Hours = t.ticker.Hours
						t.Duration.Days = t.ticker.Days
					}
				}
			}()
		} else {
			atomic.StoreInt32(&t.running, 0)
			t.ticker.Finish()
			t.End = time.Now()
		}
	}
}

// NewListTask returns a new ListTask object
func NewListTask() *ListTask {
	return &ListTask{}
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
