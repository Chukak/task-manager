package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	db "../../internal/database"
	route "../../internal/route"
	task "../../internal/tasks"
)

func main() {
	host := os.Getenv("DB_HOST")
	val, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	port := uint16(val)
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	d, err := db.NewDatabase(host, port, database, user, password)
	if err != nil {
		log.Println("Failed to connect to database: ", err.Error())
	}
	ok, err := d.Open()
	if !ok {
		log.Println("Failed to open database: ", err.Error())
	}

	task.SetDatabase(d)

	tasks := task.NewListTask()
	tasks.LoadFromDb()

	var taskContext task.TaskHTTPContext

	r := route.NewRoute(":8080")
	r.DeployDirectory("/", os.Getenv("CURRENT_SOURCE_PATH")+"bin/web/")
	r.AddGroup("api")
	taskContext.InitRoutes(r, "api")

	r.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			_ = sig
			log.Println("Stopping server...")
			r.Stop()

			tasks.InactiveAllTasks()
			os.Exit(0)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
