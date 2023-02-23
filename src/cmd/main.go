package main

import (
	"github.com/gopherWxf/gopher-gin/classes"
	"github.com/gopherWxf/gopher-gin/goft"
	"github.com/gopherWxf/gopher-gin/middlewares"
)

func main() {
	goft.
		Ignite().
		Beans(goft.NewGormAdapter(), goft.NewXOrmAdapter()).
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
