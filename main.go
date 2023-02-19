package main

import (
	"github.com/gopherWxf/gopher-gin/src/classes"
	"github.com/gopherWxf/gopher-gin/src/goft"
)

func main() {
	goft.Ignite().Mount(
		classes.NewIndexClass(),
		classes.NewUserClass(),
	).Launch()
}
