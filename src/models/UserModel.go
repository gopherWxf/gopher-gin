package models

import "fmt"

type UserModel struct {
	UserId   int    `gorm:"column:user_id" uri:"id" binding:"required,gt=0"`
	UserName string `gorm:"column:user_name"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel) String() string {
	return fmt.Sprintln("userid:", u.UserId, "username", u.UserName)
}
