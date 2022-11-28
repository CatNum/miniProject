package common

import (
	"errors"
	"reflect"
)

// CopyFields 复制结构体
// 用src的所有相同类型的字段覆盖dst
// dst应该为结构体指针
func CopyFields(dst interface{}, src interface{}) (err error) {
	at := reflect.TypeOf(dst)
	av := reflect.ValueOf(dst)
	bt := reflect.TypeOf(src)
	bv := reflect.ValueOf(src)
	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err = errors.New("dst must be a struct pointer")
		return
	}
	if bt.Kind() != reflect.Struct {
		err = errors.New("src must be a struct")
		return
	}
	av = reflect.ValueOf(av.Interface())
	ae := av.Elem()
	// 复制
	for i := 0; i < bv.NumField(); i++ {
		name := bt.Field(i).Name
		f := ae.FieldByName(name)
		bValue := bv.FieldByName(name)
		// 复制同名并类型相同的
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		}
	}
	return
}
