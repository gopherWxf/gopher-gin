package main

import (
	"github.com/gopherWxf/gopher-gin/src/classes"
	"github.com/gopherWxf/gopher-gin/src/goft"
	"github.com/gopherWxf/gopher-gin/src/middlewares"
	"log"
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
		Mount(
			"/v3",
			classes.NewArticleClass(),
		).
		Task("0/3 * * * * *", func() {
			log.Println("执行定时任务")
		}).
		Launch()
}
