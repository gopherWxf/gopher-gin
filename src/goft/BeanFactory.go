package goft

import "reflect"

type BeanFactory struct {
	beans []interface{}
}

func NewBeanFactory() *BeanFactory {
	bf := &BeanFactory{beans: make([]interface{}, 0)}
	bf.beans = append(bf.beans, bf)
	return bf
}

// 往内存中塞入bean
func (this *BeanFactory) setBean(beans ...interface{}) {
	this.beans = append(this.beans, beans...)
}

// GetBean 获取依赖
func (this *BeanFactory) GetBean(bean interface{}) interface{} {
	return this.getBean(reflect.TypeOf(bean))
}

// 得到 内存中预先设置好的 bean对象
func (this *BeanFactory) getBean(t reflect.Type) interface{} {
	for _, p := range this.beans {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

// Inject 外部方法 注入 用户写的（后面还要改,这个方法不处理注解)
func (this *BeanFactory) Inject(object interface{}) {
	vObject := reflect.ValueOf(object)
	if vObject.Kind() == reflect.Ptr { //由于不是控制器 ，所以传过来的值 不一定是指针。因此要做判断
		vObject = vObject.Elem()
	}
	for i := 0; i < vObject.NumField(); i++ {
		f := vObject.Field(i)
		if f.Kind() != reflect.Ptr || !f.IsNil() {
			continue
		}
		if p := this.getBean(f.Type()); p != nil && f.CanInterface() {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

// inject 内部方法 把bean 注入到 控制器中 (内部方法,用户控制器注入) ,同时处理注解
func (this *BeanFactory) inject(class IClass) {
	valClass := reflect.ValueOf(class).Elem()
	typeClass := reflect.TypeOf(class).Elem()
	for i := 0; i < valClass.NumField(); i++ {
		vFiled := valClass.Field(i)
		tFiled := typeClass.Field(i)
		if !vFiled.IsNil() || vFiled.Kind() != reflect.Ptr {
			continue
		}
		//判断是否是注解
		if IsAnnotation(vFiled.Type()) {
			vFiled.Set(reflect.New(vFiled.Type().Elem()))
			vFiled.Interface().(Annotation).SetTag(tFiled.Tag)
			this.Inject(vFiled.Interface())
			continue
		}
		//用户控制器注入
		if p := this.getBean(vFiled.Type()); p != nil {
			vFiled.Set(reflect.New(vFiled.Type().Elem()))
			vFiled.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
