package main

import (
	"os"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile(os.Getenv("CURRENT_SOURCE_PATH")+"bin/web/src/", true)))
	api := r.Group("/api")
	api.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now().Format("3:4:5"),
		})
	})

	r.Run()
}
