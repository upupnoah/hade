package framework

import (
	"log"
	"net/http"
)

// Core is the core of the framework
type Core struct {
	router map[string]ControllerHandler
}

// NewCore create a new core
func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// ServeHTTP implement http.Handler interface
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serverHTTP")
	ctx := NewContext(request, response)

	// 一个简单的路由选择器，这里直接写死为测试路由 foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	err := router(ctx)
	if err != nil {
		return
	}
}
