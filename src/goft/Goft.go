package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Goft struct {
	*gin.Engine
	g   *gin.RouterGroup
	dba interface{}
}

func Ignite() *Goft {
	goft := &Goft{Engine: gin.New()}
	//强迫加载的异常中间件
	goft.Use(ErrorHandler())
	//goft.Attach(middlewares.NewRecoverMid())
	return goft
}
func (this *Goft) Launch() {
	fmt.Println("http://127.0.0.1")
	this.Run(":80")
}
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
		valClass := reflect.ValueOf(class).Elem()
		if valClass.NumField() > 0 {
			if this.dba != nil {
				valClass.Field(0).Set(reflect.New(valClass.Field(0).Type().Elem()))
				valClass.Field(0).Elem().Set(reflect.ValueOf(this.dba).Elem())
			}
		}
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
func (this *Goft) DB(dba interface{}) *Goft {
	this.dba = dba
	return this
}
