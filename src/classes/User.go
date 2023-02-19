package classes

import "github.com/gin-gonic/gin"

type UserClass struct {
	*gin.Engine
}

func NewUserClass(engine *gin.Engine) *UserClass {
	return &UserClass{Engine: engine}
}
func (this *UserClass) Build() {
	this.Handle("GET", "/user", this.GetUser())
}

func (this *UserClass) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	}
}
