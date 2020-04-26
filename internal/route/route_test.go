package route

import (
	"os"
	"os/signal"
	"testing"
	"time"

	test "github.com/chukak/task-manager/pkg/checks"
	"github.com/gin-gonic/gin"
)

func TestRouteInitializationr(t *testing.T) {
	test.SetT(t)
	gin.SetMode(gin.ReleaseMode)

	address := "localhost:8080"
	route := NewRoute(address)

	test.CheckNotEqual(route.engine, nil)
	test.CheckNotEqual(route.server, nil)
	test.CheckEqual(len(route.groups), 0)
	test.CheckEqual(route.Address, address)
}

func TestRouteFunctionality(t *testing.T) {
	test.SetT(t)
	gin.SetMode(gin.ReleaseMode)

	address := "localhost:8080"
	route := NewRoute(address)

	testStr := "aaa"
	route.addSlash(&testStr)
	test.CheckEqual(testStr, "/aaa")
	testStr = "/bbb"
	route.addSlash(&testStr)
	test.CheckEqual(testStr, "/bbb")

	apiGroup := "/api"
	fileGroup := "file"
	route.AddGroup(apiGroup)
	route.AddGroup(fileGroup)

	test.CheckEqual(len(route.groups), 2)
	test.CheckNotEqual(route.groups[apiGroup], nil)
	test.CheckNotEqual(route.groups["/"+fileGroup], nil) // slash

	handler := func(c *gin.Context) {}

	route.AddRequest(apiGroup, GET, "time", handler)
	test.CheckNotEqual(route.groups[apiGroup], nil)

	savedLen := len(route.engine.RouterGroup.Handlers)
	route.DeployDirectory("/", "/staticfiles")
	test.CheckEqual(len(route.engine.RouterGroup.Handlers), savedLen+1)
}

func TestRouteLaunchServer(t *testing.T) {
	test.SetT(t)
	gin.SetMode(gin.ReleaseMode)

	address := "localhost:8080"
	route := NewRoute(address)

	ok := false
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		ok = route.Stop()
	}()

	route.Start()
	time.Sleep(time.Second * 2)
	// raise signal
	pid, err := os.FindProcess(os.Getpid())
	test.CheckEqual(err, nil)
	pid.Signal(os.Interrupt)

	time.Sleep(time.Second * 1)
	test.CheckTrue(ok)
}
