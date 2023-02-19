package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/gopherWxf/gopher-gin/src/classes"
)

func main() {
	r := gin.New()
	NewIndexClass(r).Build()

	fmt.Println("http://127.0.0.1")
	r.Run(":80")
}
