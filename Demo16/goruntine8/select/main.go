package main

import (
	"fmt"
	"time"
)

/**
使用select 从多个管道获取数据	。select 的使用类似于 switch 语句， 它有一系列 case 分支和一个默认的分支。 每个 case 会对
应一个管道的通信（接收或发送） 过程。 select 会一直等待， 直到某个 case 的通信操作完成时， 就会执行 case 分支对应的语句。
*/

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}
func main() {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
