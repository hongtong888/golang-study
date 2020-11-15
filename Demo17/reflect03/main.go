package main

import (
	"fmt"
	"reflect"
)

/**
reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的信息，reflect.Value提供了获取原始值的一些方法
*/
func main() {
	var a = 13
	reflectValue(a)
}

//reflect.Value 提供了很多api来获取原始值
func reflectValue(x interface{}) {
	//reflect := reflect.ValueOf(x)
	//result := reflect.Int() + 10
	//fmt.Println(result)

	var a int64 = 100
	var b float32 = 12.3
	var c string = "你好golang"
	reflectValueBySwitch(a)
	reflectValueBySwitch(b)
	reflectValueBySwitch(c)
}

//通过switch 来写一个通用的获取原始值的方法
func reflectValueBySwitch(x interface{}) {
	valueOf := reflect.ValueOf(x)
	kind := valueOf.Kind()
	switch kind {
	case reflect.Float64:
		fmt.Printf("Float64类型的原始值%v\n", valueOf.Float())
	case reflect.Float32:
		fmt.Printf("Float32类型的原始值%v\n", valueOf.Float())
	case reflect.Int64:
		fmt.Printf("Int64类型的原始值%v\n", valueOf.Int())
	case reflect.String:
		fmt.Printf("string类型的原始值%v\n", valueOf.String())
	default:
		fmt.Printf("未找到类型\n")
	}
}
