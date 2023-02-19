package middlewares

import (
	"github.com/gin-gonic/gin"
)

type RecoverMid struct {
}

func NewRecoverMid() *RecoverMid {
	return &RecoverMid{}
}

func (u *RecoverMid) OnRequest(ctx *gin.Context) error {
	defer func() {
		if err := recover(); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err})
		}
	}()
	ctx.Next()
	return nil
}
