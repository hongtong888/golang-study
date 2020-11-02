package main

import "fmt"

/**
空接口，没有任何约束，任何类型都能实现空接口。也可以当做类型来使用，表示任何类型
*/

type A interface{}

func show(a interface{}) {
	fmt.Printf("值:%v 类型:%T\n", a, a)
}

func main() {
	//1、任何类型来实现空接口
	//var a A
	//var str = "你好golang"
	//a = str                          //让字符串实现A这个接口
	//fmt.Printf("值:%v 类型:%T\n", a, a) //值:你好golang 类型:string
	//
	//var num = 20
	//a = num //表示让int类型实现A这个接口
	//fmt.Printf("值:%v 类型:%T\n", a, a)
	//
	//var flag = true
	//a = flag //表示让bool类型实现A这个接口
	//fmt.Printf("值:%v 类型:%T", a, a)

	//2、空接口也可以当做任意类型
	show(20)
	show("你好golang")
	slice := []int{1, 2, 34, 4}
	show(slice)

	//var m1 = make(map[string]string)
	//m1["username"] = "张三"
	//m1["age"] = "20"
	//fmt.Println(m1)

	var m1 = make(map[string]interface{})
	m1["username"] = "张三"
	m1["age"] = 20
	m1["married"] = true
	fmt.Println(m1)

	//var s1 = []interface{}{1, 2, "你好", true} //[1 2 你好 true]
	//fmt.Println(s1)
}
