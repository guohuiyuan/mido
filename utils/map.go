package utils

import (
	"fmt"
	"reflect"
)

// ToMap 结构体转map[string]interface{}
func ToMap(in interface{}, tagName string) (out map[string]interface{}, err error) {
	out = make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		err = fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return
}
