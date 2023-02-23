package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Goft struct {
	*gin.Engine
	g     *gin.RouterGroup
	props []interface{}
}

func Ignite() *Goft {
	goft := &Goft{Engine: gin.New(), props: make([]interface{}, 0)}
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
		this.setProp(class)
	}
	return this
}
func (this *Goft) getProp(t reflect.Type) interface{} {
	for _, p := range this.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}
func (this *Goft) setProp(class IClass) {
	valClass := reflect.ValueOf(class).Elem()
	for i := 0; i < valClass.NumField(); i++ {
		f := valClass.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		if p := this.getProp(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
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
	this.props = append(this.props, dba)
	return this
}
