package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/gopherWxf/gopher-gin/src/goft"
	"github.com/gopherWxf/gopher-gin/src/models"
)

type UserClass struct {
	*goft.GormAdapter
}

func NewUserClass() *UserClass {
	return &UserClass{}
}
func (this *UserClass) UserTest(ctx *gin.Context) string {
	return "User Test"
}

func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/test", this.UserTest)
	goft.Handle("GET", "/user/:id", this.UserDetail)
	goft.Handle("GET", "/userlist", this.UserList)
}
func (this *UserClass) UserDetail(ctx *gin.Context) goft.Model {
	user := models.NewUserModel()
	err := ctx.BindUri(user)
	goft.Error(err, "ID 参数 不合法")
	err = this.Table("users").
		Where("user_id=?", user.UserId).
		Find(user).Error
	goft.Error(err)

	return user
}
func (this *UserClass) UserList(ctx *gin.Context) goft.Models {
	users := []*models.UserModel{
		{UserId: 2, UserName: "wxf2"},
		{UserId: 3, UserName: "wxf3"},
	}
	return goft.MakeModels(users)
}
