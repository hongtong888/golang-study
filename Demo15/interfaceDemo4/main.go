package main

import "fmt"

/**
方法可以把接口作为参数传递，传递接口的具体实现类,并使用类型断言来对传递类型进行判断
*/

//定义usb的一个接口
type Usb interface {
	start()
	stop()
}

//定义一个computer结构体
type Computer struct {
}

func (computer Computer) start() {
	fmt.Println("启动")
}

func (computer Computer) stop() {
	fmt.Println("关闭")
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

//定义一个方法，接收接口作为参数
func (computer Computer) work(usb Usb) {
	if _, ok := usb.(Phone); ok {
		usb.start()
	} else {
		usb.stop()
	}

}

func main() {
	//初始化computer结构体
	computer := Computer{}

	phone := Phone{
		Name: "华为手机",
	}

	computer.work(phone)
	computer.work(computer)
}
