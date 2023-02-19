package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/gopherWxf/gopher-gin/src/goft"
	"github.com/gopherWxf/gopher-gin/src/models"
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
	goft.Handle("GET", "/user2", this.UserDetail)

}
func (this *UserClass) UserDetail(ctx *gin.Context) goft.Model {
	return &models.UserModel{UserId: 2, UserName: "wxf"}
}
