package main

import "fmt"

/**
管道是 Golang 在语言级别上提供的 goroutine 间的通讯方式， 我们可以使用 channel 在
多个 goroutine 之间传递消息。 如果说 goroutine 是 Go 程序并发的执行体， channel 就是它们
之间的连接。 channel 是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
管道创建：make(chan 元素类型, 容量)
*/

func main() {
	//1、创建管道
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30

	//2、从管道取值
	a := <-ch
	fmt.Println(a)

	<-ch //从管道里面取值   //20
	c := <-ch
	fmt.Println(c) //30
	ch <- 56
	ch <- 66
	//4、打印管道的长度和容量
	fmt.Printf("值：%v 容量：%v 长度%v\n", ch, cap(ch), len(ch)) //值：0xc0000d0080 容量：3 长度2

	// 5、管道的类型（引用数据类型）

	ch1 := make(chan int, 4)

	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 69
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d) //69

	//6、管道阻塞

	ch6 := make(chan int, 1)
	ch6 <- 34
	ch6 <- 64 //all goroutines are asleep - deadlock!

	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	ch7 := make(chan string, 2)
	ch7 <- "数据1"
	ch7 <- "数据2"
	m1 := <-ch7
	m2 := <-ch7
	m3 := <-ch7
	fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock!

	//正确的写法
	ch8 := make(chan int, 1)
	ch8 <- 34
	<-ch8
	ch8 <- 67
	<-ch8
	ch8 <- 78
	m4 := <-ch8
	fmt.Println(m4)

	//7、定义一个空接口存放任意数据类型的管道
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat

	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的Cat是什么?

	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	//下面的写法是错误的!编译不通过
	//fmt.Printf("newCat.Name	=%v", newCat.Name)
	//使用类型断言
	cat1 := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", cat1.Name)
}

type Cat struct {
	Name string
	age  int
}
