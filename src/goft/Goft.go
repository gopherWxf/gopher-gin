package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Ignite() *Goft {
	return &Goft{Engine: gin.New()}
}
func (this *Goft) Launch() {
	fmt.Println("http://127.0.0.1")
	this.Run(":80")
}
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}
func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(ctx *gin.Context) {
		err := f.OnRequest(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		} else {
			ctx.Next()
		}
	})
	return this
}
