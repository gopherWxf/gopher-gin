package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
}

func Ignite() *Goft {
	return &Goft{Engine: gin.New()}
}
func (this *Goft) Launch() {
	fmt.Println("http://127.0.0.1")
	this.Run(":80")
}
func (this *Goft) Mount(classes ...IClass) *Goft {
	for _, class := range classes {
		class.Build(this)
	}
	return this
}
