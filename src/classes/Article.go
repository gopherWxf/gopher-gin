package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/gopherWxf/gopher-gin/src/goft"
	"github.com/gopherWxf/gopher-gin/src/models"
)

/*
	使用：
	1. 创建一个控制器type ArticleClass struct
	2. 创建build和detail
	3. 创建数据库模型 ArticleModel
	4. 进行业务处理,返回responder类型,如果没有则自己注册
	5. 控制器注入到启动函数里面
*/

type ArticleClass struct {
	*goft.GormAdapter
}

func NewArticleClass() *ArticleClass {
	return &ArticleClass{}
}
func (this *ArticleClass) ArticleDetail(ctx *gin.Context) goft.Model {
	news := models.NewArticleModel()
	goft.Error(ctx.ShouldBindUri(news))
	goft.Error(this.Table("mynews").Where("id=?", news.NewsID).Find(news).Error)
	return news
}
func (this *ArticleClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/article/:id", this.ArticleDetail)
}
