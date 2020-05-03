package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	route "../../internal/route"
	task "../../internal/tasks"
)

// All code below only for example!

type ListTask struct {
	List []*task.Task `json:"listTasks"`
}

func main() {
	var tasks ListTask
	t := task.NewTask(nil)
	t.Title = "Task 1"
	t.Description = "This is task description"
	t.Priority = 4

	t2 := task.NewTask(nil)
	t2.Title = "Task 2"
	t2.Description = "This is task 2 description"
	t2.Priority = 1

	tasks.List = append(tasks.List, t)
	tasks.List = append(tasks.List, t2)
	gin.Bind(tasks)

	r := route.NewRoute(":8080")
	r.DeployDirectory("/", os.Getenv("CURRENT_SOURCE_PATH")+"bin/web/")
	r.AddGroup("api")
	r.AddRequest("api", route.GET, "/task/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"taskCount": len(tasks.List),
			"listTasks": tasks.List,
		})
	})

	r.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			_ = sig
			log.Println("Stopping server...")
			r.Stop()
			os.Exit(0)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
