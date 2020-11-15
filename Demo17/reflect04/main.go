package main

import (
	"fmt"
	"reflect"
)

/**
reflect.ValueOf()返回的是reflect.Value类型，可以通过反射修改设置变量的值
因为要修改值，传递的必须是变量值的地址，而变量的kind 是ptr,这时 v.Elem() 可以获取具体的值
*/
func main() {
	var a int64 = 100
	reflectSetValue(&a)

	fmt.Println(a)

	var b string = "你好golang"
	reflectSetValue(&b)
	fmt.Println(b)
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// fmt.Println(v.Kind()) //ptr
	// fmt.Println(v.Elem().Kind()) //int64
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("你好 go语言")
	}
}
