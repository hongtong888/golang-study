package main

import (
	"fmt"
	"sync"
)

/**
需求：使用goroutine和channel协同工作案例
1、开启一个WriteData的的协程给向管道inChan中写入100条数据
2、开启一个ReadData的协程读取inChan中写入的数据
3、注意：WriteData和ReadData同时操作一个管道
4、主线程必须等待操作完成后才可以退出

goroutine结合Channel使用的简单demo,定义两个方法，一个方法给管道里面写数据，一个给管道里面读取数据。要求同步进行。
*/
var wp sync.WaitGroup

func main() {
	ch := make(chan int, 10)
	wp.Add(1)
	go writeDate(ch)
	wp.Add(1)
	go readDate(ch)
	wp.Wait()
	fmt.Println("退出...")

	//匿名函数的写法
	// 1、创建channel
	//var ch1 = make(chan int, 3)
	//wp.Add(1)
	//go func() {
	//	for i := 1; i <= 3; i++ {
	//		num := <-ch1
	//		fmt.Println(num)
	//	}
	//	wp.Done()
	//}()
	//wp.Add(1)
	//go func() {
	//	for i := 1; i <= 3; i++ {
	//		time.Sleep(time.Second)
	//		ch1 <- i
	//	}
	//	wp.Done()
	//}()
	//wp.Wait()
}

func writeDate(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("【写入】数据%v成功\n", i)
	}
	close(ch)
	wp.Done()
}

func readDate(ch chan int) {
	for value := range ch {
		fmt.Printf("【读取】数据%v成功\n", value)
	}
	wp.Done()
}
