package tasks

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	route "../route"
)

// TaskHTTPWrapper is a wrapper interface for http requests for Task object
type TaskHTTPWrapper interface {
	InitRoutes(*route.Route, string)
	NewTaskRequest(*gin.Context)
	OpenTaskRequest(*gin.Context)
	ActiveTaskRequest(*gin.Context)
	RemoveTaskRequest(*gin.Context)
	UpdateTaskRequest(*gin.Context)
	AllTasksRequest(*gin.Context)
	GetTaskDataRequest(*gin.Context)
	updateList()
}

// TaskHTTPContext store list tasks abd release methods
type TaskHTTPContext struct {
}

type JSONBody map[string]interface{}

func getJSONValueByKey(e *interface{}, body *JSONBody, key string) bool {
	var ok bool
	*e, ok = (*body)[key]
	if !ok {
		msg := fmt.Sprintf("JSON body does have a key '%s'.", key)
		log.Println(msg)
	}
	return ok
}

// InitRoutes inits routes for tasks
func (tc *TaskHTTPContext) InitRoutes(r *route.Route, group string) {
	r.AddRequest(group, route.GET, "/task/all", tc.AllTasksRequest)
	r.AddRequest(group, route.POST, "/task/new", tc.NewTaskRequest)
	r.AddRequest(group, route.POST, "/task/open", tc.OpenTaskRequest)
	r.AddRequest(group, route.POST, "/task/active", tc.ActiveTaskRequest)
	r.AddRequest(group, route.POST, "/task/remove", tc.RemoveTaskRequest)
	r.AddRequest(group, route.POST, "/task/update", tc.UpdateTaskRequest)
	r.AddRequest(group, route.POST, "/task/get", tc.GetTaskDataRequest)
}

// NewTaskRequest is a wrapper to create a task
func (tc *TaskHTTPContext) NewTaskRequest(c *gin.Context) {
	body := JSONBody{}
	_ = c.Bind(&body)

	var parentID interface{}
	if !getJSONValueByKey(&parentID, &body, "parent") {
		parentID = nil
	}

	var parent *Task = nil
	var ok bool = false
	if parentID != nil {
		if parent, ok = taskPointers[int64(parentID.(float64))]; !ok {
			parent = nil
		}
	}

	task := NewTask(parent)
	c.JSON(http.StatusOK, gin.H{
		"id": task.TaskID,
	})
}

// OpenTaskRequest is a wrapper to open or close a task
func (tc *TaskHTTPContext) OpenTaskRequest(c *gin.Context) {
	body := JSONBody{}
	c.BindJSON(&body)

	var id, open interface{}
	if !getJSONValueByKey(&id, &body, "id") {
		return
	}
	if !getJSONValueByKey(&open, &body, "open") {
		return
	}

	if task, ok := taskPointers[int64(id.(float64))]; ok {
		task.Open(open.(bool))
		c.JSON(http.StatusOK, gin.H{
			"id": task.TaskID,
		})
	}
}

// ActiveTaskRequest is a wrapper to run a task
func (tc *TaskHTTPContext) ActiveTaskRequest(c *gin.Context) {
	body := JSONBody{}
	c.BindJSON(&body)

	var id, active interface{}
	if !getJSONValueByKey(&id, &body, "id") {
		return
	}
	if !getJSONValueByKey(&active, &body, "active") {
		return
	}

	if task, ok := taskPointers[int64(id.(float64))]; ok {
		task.SetActive(active.(bool))
		c.JSON(http.StatusOK, gin.H{
			"id": task.TaskID,
		})
	}
}

// RemoveTaskRequest is a wrapper to remote task or subtask
func (tc *TaskHTTPContext) RemoveTaskRequest(c *gin.Context) {
	body := JSONBody{}
	c.BindJSON(&body)

	var id interface{}
	if !getJSONValueByKey(&id, &body, "id") {
		return
	}

	if task, ok := taskPointers[int64(id.(float64))]; ok {
		if task.parent != nil {
			task.parent.RemoveSubtask(task)
		} else {
			task.RemoveSelf()
		}
	}
}

// UpdateTaskRequest is a wrapper to update task data
func (tc *TaskHTTPContext) UpdateTaskRequest(c *gin.Context) {
	body := JSONBody{}
	_ = c.Bind(&body)

	var id, title, desc, priority interface{}
	if !getJSONValueByKey(&id, &body, "id") {
		return
	}
	if !getJSONValueByKey(&title, &body, "title") {
		title = nil
	}
	if !getJSONValueByKey(&desc, &body, "description") {
		desc = nil
	}
	if !getJSONValueByKey(&priority, &body, "priority") {
		priority = nil
	}

	if task, ok := taskPointers[int64(id.(float64))]; ok {
		if title != nil {
			task.Title = title.(string)
		}
		if desc != nil {
			task.Description = desc.(string)
		}
		if priority != nil {
			task.Priority = int8(priority.(float64))
		}
		task.UpdateInDb()

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

// AllTasksRequest is a wrapper to get all the tasks
func (tc *TaskHTTPContext) AllTasksRequest(c *gin.Context) {
	tc.updateList()

	list := listTaskPointer
	if list == nil {
		list = &ListTask{}
	}
	c.JSON(http.StatusOK, gin.H{
		"listTasks": list.List,
	})
}

// GetTaskDataRequest get data from task
func (tc *TaskHTTPContext) GetTaskDataRequest(c *gin.Context) {
	body := JSONBody{}
	_ = c.Bind(&body)

	var id interface{}
	if !getJSONValueByKey(&id, &body, "id") {
		return
	}
	if task, ok := taskPointers[int64(id.(float64))]; ok {
		c.JSON(http.StatusOK, gin.H{
			"id":          task.TaskID,
			"title":       task.Title,
			"description": task.Description,
			"priority":    task.Priority,
			"active":      task.IsActive,
			"opened":      task.IsOpened,
			"start":       task.Start.Truncate(time.Second),
			"end":         task.End.Truncate(time.Second),
		})
	}
}

func (tc *TaskHTTPContext) updateList() {
	list := listTaskPointer
	if list == nil {
		list = &ListTask{}
	}
	gin.Bind(*list)
}
