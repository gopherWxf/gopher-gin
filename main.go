package main

import (
	"github.com/gopherWxf/gopher-gin/src/classes"
	"github.com/gopherWxf/gopher-gin/src/goft"
)

func main() {
	goft.Ignite().
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
