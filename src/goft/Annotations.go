package goft

import "reflect"

// 注解
type Annotation interface {
	SetTag(tag reflect.StructTag)
}

var AnnotationList []Annotation

// 判断当前的 注入对象 是否是注解
func IsAnnotation(t reflect.Type) bool {
	for _, item := range AnnotationList {
		if reflect.TypeOf(item) == t {
			return true
		}
	}
	return false
}
func init() {
	AnnotationList = make([]Annotation, 0)
	AnnotationList = append(AnnotationList, new(Value))
}

type Value struct {
	tag reflect.StructTag
}

func (this *Value) SetTag(tag reflect.StructTag) {
	this.tag = tag
}
func (this *Value) String() string {
	return "21"
}
