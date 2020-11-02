package main

import "fmt"

/**
类型断言：一个接口的值由具体类型和具体类型的值两部分组成，如果我们想要判断空接口的类型，可以使用类型断言
x.(T),x表示类型为空接口interface{}的变量   T表示断言的类型
*/
func main() {
	testType(10)
	testType1("10")
}

//用if测试类型断言
func testType(a interface{}) {
	if _, ok := a.(string); ok {
		fmt.Println("类型为string类型")
	} else if _, ok := a.(int); ok {
		fmt.Println("类型为int类型")
	} else if _, ok := a.(bool); ok {
		fmt.Println("类型为bool类型")
	}
}

func testType1(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("类型为int类型")
	case string:
		fmt.Println("类型为string类型")
	case bool:
		fmt.Println("类型为bool类型")
	}

}
