package main

import "fmt"

/**
当向管道中发送完数据时， 我们可以通过 close 函数来关闭管道。当管道被关闭时， 再往该管道发送值会引发 panic， 从该管道取值的操作会先取完管道中的
值， 再然后取到的值一直都是对应类型的零值。
*/

func main() {
	//1、使用for range遍历通道，当通道被关闭的时候就会退出for range,如果没有关闭管道就会报个错误fatal error: all goroutines are asleep - deadlock!
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1) //关闭管道

	//for range循环遍历管道的值  ,注意：管道没有key
	for v := range ch1 {
		fmt.Println(v)
	}
	//2、通过for循环遍历管道的时候管道可以不关闭
	var ch2 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}

	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}
}
