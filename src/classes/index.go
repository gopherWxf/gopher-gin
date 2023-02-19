package classes

import "github.com/gin-gonic/gin"

type IndexClass struct {
	*gin.Engine
}

func NewIndexClass(engine *gin.Engine) *IndexClass {
	return &IndexClass{Engine: engine}
}
func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "index ok",
		})
	}
}
func (this *IndexClass) Build() {
	this.Handle("GET", "/", this.GetIndex())
}
