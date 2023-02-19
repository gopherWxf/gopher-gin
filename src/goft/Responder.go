package goft

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var ResponderList []Responder

func init() {
	ResponderList = []Responder{
		new(StringResponder),
		new(ModelResponder),
		new(ModelsResponder),
	}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	//StringResponder(  handler.(func(*gin.Context) string)  ).RespondTo()
	h_ref := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		r_ref := reflect.ValueOf(r).Elem()
		if h_ref.Type().ConvertibleTo(r_ref.Type()) {
			r_ref.Set(h_ref)
			return r_ref.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(ctx *gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200, this(ctx))
	}
}

type ModelResponder func(ctx *gin.Context) Model

func (this ModelResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, this(ctx))
	}
}

type ModelsResponder func(ctx *gin.Context) Models

func (this ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteString(string(this(ctx)))
	}
}
