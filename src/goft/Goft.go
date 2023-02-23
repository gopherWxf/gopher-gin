package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gopherWxf/gopher-gin/funcs"
	"log"
)

type Goft struct {
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
}

func Ignite() *Goft {
	//gin.SetMode(gin.ReleaseMode)
	goft := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory()}
	//强迫加载的异常中间件
	goft.Use(ErrorHandler())
	//整个配置加载进bean中
	config := InitConfig()
	goft.beanFactory.setBean(config)
	if config.Server.Html != "" {
		goft.FuncMap = funcs.FuncMap
		goft.LoadHTMLGlob(config.Server.Html)
	}
	return goft
}
func (this *Goft) Launch() {
	var port int32 = 80
	config := this.beanFactory.GetBean(new(SysConfig))
	if config != nil {
		port = config.(*SysConfig).Server.Port
	}
	getCronTask().Start()
	fmt.Printf("http://127.0.0.1:%d\n", port)
	this.Run(fmt.Sprintf(":%d", port))
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

// corn表达式 增加定时任务
func (this *Goft) Task(expr string, f func()) *Goft {
	_, err := getCronTask().AddFunc(expr, f)
	if err != nil {
		log.Println(err)
	}
	return this
}
