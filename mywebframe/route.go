package main

import (
	"webframe/farmework"
)

func registerRouter(core *farmework.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	core.Get("foo", FooControllerHandler)
}