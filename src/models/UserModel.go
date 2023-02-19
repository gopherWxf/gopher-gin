package models

import "fmt"

type UserModel struct {
	UserId   int `uri:"id" binding:"required,gt=0"`
	UserName string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel) String() string {
	return fmt.Sprintln("userid:", u.UserId, "username", u.UserName)
}
