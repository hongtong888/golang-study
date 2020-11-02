package main

import "fmt"

/**
值接收者： 如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
指针接收者： 如果结构体中的方法是指针接收者，那么实例化后结构体指针类型都可以赋值给接口变量， 结构体值类型没法赋值给接口变量。
*/

//定义usb的一个接口
type Usb interface {
	start()
	stop()
}

//接口的实现类 手机结构体
type Phone struct {
	Name string
}

//手机类实现接口的方法
func (phone Phone) start() {
	fmt.Println(phone.Name, "启动")
}

func (phone Phone) stop() {
	fmt.Println(phone.Name, "关闭")
}

func main() {
	phone := Phone{
		Name: "华为手机",
	}
	phone.stop()
}
