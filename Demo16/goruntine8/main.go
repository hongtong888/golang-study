package main

import "fmt"

/**
使用select 从多个管道获取数据	。select 的使用类似于 switch 语句， 它有一系列 case 分支和一个默认的分支。 每个 case 会对
应一个管道的通信（接收或发送） 过程。 select 会一直等待， 直到某个 case 的通信操作完成时， 就会执行 case 分支对应的语句。
*/

func main() {

	//定义两个channel
	ints := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ints <- i
	}

	strings := make(chan string, 10)
	for i := 0; i < 10; i++ {
		strings <- "hello-" + fmt.Sprintf("%d", i)
	}

	//使用select来获取channel里面的数据的时候不需要关闭channel
	for {
		select {
		case value := <-ints:
			fmt.Printf("从ints管道读取的数据%d\n", value)
		case value := <-strings:
			fmt.Printf("从strings管道读取的数据%v\n", value)
		default:
			fmt.Printf("数据获取完毕")
			return //注意退出...
		}
	}
}
