package middlewares

import "fmt"

type UserMid struct {
}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (u *UserMid) OnRequest() error {
	fmt.Println("这是新的用户中间件")
	return nil
}
