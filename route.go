package main

import "github.com/upupnoah/hade/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
