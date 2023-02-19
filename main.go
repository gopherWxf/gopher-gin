package main

import (
	"github.com/gopherWxf/gopher-gin/src/classes"
	"github.com/gopherWxf/gopher-gin/src/goft"
	"github.com/gopherWxf/gopher-gin/src/middlewares"
)

func main() {
	goft.
		Ignite().
		Attach(
			middlewares.NewUserMid(),
		).
		Mount(
			"/v1",
			classes.NewIndexClass(),
		).
		Mount(
			"/v2",
			classes.NewUserClass(),
		).
		Launch()
}
