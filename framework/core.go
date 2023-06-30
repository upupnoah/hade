package framework

import "net/http"

// Core is the core of the framework
type Core struct {
}

// NewCore create a new core
func NewCore() *Core {
	return new(Core)
}

// ServeHTTP implement http.Handler interface
func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
