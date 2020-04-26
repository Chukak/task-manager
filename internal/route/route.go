package route

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// RequestType type of http-request (GET, POST...)
type RequestType int

const (
	// GET request
	GET RequestType = 0
	// POST request
	POST RequestType = 1
)

// RouteManage is a common interface for route
type RouteManage interface {
	addSlash(*string)
	Start()
	Stop() bool
	AddGroup(string)
	AddRequest(string, RequestType, string, ...gin.HandlerFunc)
	DeployDirectory(string, string)
}

// Route is a object for routing on any url
type Route struct {
	engine  *gin.Engine
	groups  map[string]*gin.RouterGroup
	server  *http.Server
	Address string
}

// NewRoute returns is a new Route object
func NewRoute(addr string) *Route {
	route := Route{engine: gin.Default(), groups: map[string]*gin.RouterGroup{},
		server: &http.Server{Addr: addr}, Address: addr}
	route.server.Handler = route.engine
	return &route
}

func (r *Route) addSlash(str *string) {
	if len(*str) > 0 {
		if (*str)[0] != '/' {
			*str = "/" + *str
		}
	}
}

// Start a router
func (r *Route) Start() {
	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Starting server failed: ", err)
		}
	}()
}

// Stop a router
func (r *Route) Stop() bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error = nil
	if err = r.server.Shutdown(ctx); err != nil {
		log.Fatal("Stoping server failed: ", err)
	}
	return err == nil
}

// AddGroup adds a new route group
func (r *Route) AddGroup(group string) {
	r.addSlash(&group)
	r.groups[group] = r.engine.Group(group)
}

// AddRequest adds a new action for this request in group
func (r *Route) AddRequest(groupName string, request RequestType,
	url string, handlers ...gin.HandlerFunc) {
	r.addSlash(&groupName)
	r.addSlash(&url)
	switch request {
	case GET:
		r.groups[groupName].GET(url, handlers...)
	case POST:
		r.groups[groupName].POST(url, handlers...)
	}
}

// DeployDirectory deploy a statisfiles directory
func (r *Route) DeployDirectory(url string, dirpath string) {
	r.engine.Use(static.Serve(url, static.LocalFile(dirpath, true)))
}
