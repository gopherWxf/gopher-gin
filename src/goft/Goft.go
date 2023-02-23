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
	config := InitConfig()
	fmt.Printf("http://127.0.0.1:%d\n", config.Server.Port)
	this.Run(fmt.Sprintf(":%d", config.Server.Port))
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
	typeClass := reflect.TypeOf(class).Elem()
	for i := 0; i < valClass.NumField(); i++ {
		vFiled := valClass.Field(i)
		tFiled := typeClass.Field(i)
		if !vFiled.IsNil() || vFiled.Kind() != reflect.Ptr {
			continue
		}
		if p := this.getProp(vFiled.Type()); p != nil {
			vFiled.Set(reflect.New(vFiled.Type().Elem()))
			vFiled.Elem().Set(reflect.ValueOf(p).Elem())
			//判断是否是注解
			if IsAnnotation(vFiled.Type()) {
				p.(Annotation).SetTag(tFiled.Tag)
			}
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
func (this *Goft) Beans(beans ...interface{}) *Goft {
	this.props = append(this.props, beans...)
	return this
}
