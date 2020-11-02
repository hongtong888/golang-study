package main

import "fmt"

/**
空接口，没有任何约束，任何类型都能实现空接口
*/

type A interface{}

func main() {
	var a A
	var str = "你好golang"
	a = str                          //让字符串实现A这个接口
	fmt.Printf("值:%v 类型:%T\n", a, a) //值:你好golang 类型:string

	var num = 20
	a = num //表示让int类型实现A这个接口
	fmt.Printf("值:%v 类型:%T\n", a, a)

	var flag = true
	a = flag //表示让bool类型实现A这个接口
	fmt.Printf("值:%v 类型:%T", a, a)
}
