package main

import (
	"fmt"
	"reflect"
)

/**
reflect.TypeOf()获取任意值的类型对象，在反射中类型还分为两种 type和name
*/
func main() {
	//基本类型
	a := 10         //int
	b := 23.4       //float64
	c := true       //bool
	d := "你好golang" //string

	//引用类型
	person := Person{
		Name: "你好",
		Age:  15,
	}
	var e myInt = 34

	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)
	reflectFn(person)
	reflectFn(e)

	var h = 25
	reflectFn(&h)

	var i = [3]int{1, 2, 3}
	reflectFn(i) //[3]int 类型名称: 类型种类:array

	var j = []int{11, 22, 33}
	reflectFn(j) //[]int 类型名称: 类型种类:slice
}

func reflectFn(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("TypeOf:%v Name:%v Kind:%v\n", t, t.Name(), t.Kind())
}

type myInt int
type Person struct {
	Name string
	Age  int
}
