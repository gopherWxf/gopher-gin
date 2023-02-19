package models

import "fmt"

type UserModel struct {
	UserId   int
	UserName string
}

func (u *UserModel) String() string {
	return fmt.Sprintln("userid:", u.UserId, "username", u.UserName)
}
