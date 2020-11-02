package main

import "fmt"

/**
golang中接口是一种抽象的数据类型，定义了对象的行为规范，只定义不实现，由具体的对象来实现
*/

//定义usb的一个接口
type Usb interface {
	start()
	stop()
}

//接口的实现类
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

type Camera struct {
}

func main() {
	//初始化手机结构体
	phone := Phone{
		Name: "苹果",
	}
	var usb Usb = phone
	usb.start()
}
