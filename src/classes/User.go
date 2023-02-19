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
func (this *UserClass) GetUser(ctx *gin.Context) string {
	return "wxf"
}
func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.GetUser)
}
