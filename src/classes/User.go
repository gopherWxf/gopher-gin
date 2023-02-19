package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/gopherWxf/gopher-gin/src/goft"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.GetUser())
}

func (this *UserClass) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	}
}
