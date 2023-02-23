package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
}

func Ignite() *Goft {
	goft := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory()}
	//强迫加载的异常中间件
	goft.Use(ErrorHandler())
	//整个配置加载进bean中
	goft.beanFactory.setBean(InitConfig())
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
		this.beanFactory.inject(class)
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
func (this *Goft) Beans(beans ...interface{}) *Goft {
	this.beanFactory.setBean(beans...)
	return this
}
