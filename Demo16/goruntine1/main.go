package main

import (
	"fmt"
	"sync"
	"time"
)

/**
主线程退出后所有的协程无论有没有执行完毕都会退出，所以我们在主进程中可以通过WaitGroup等待协程执行完毕
*/
var wp sync.WaitGroup

func test() {
	defer wp.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	wp.Add(1)
	go test() //表示开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("main() 你好golang")
		time.Sleep(time.Millisecond * 10)
	}
	wp.Wait()

	//获取当前计算机上面的Cup个数
	//cpuNum := runtime.NumCPU()
	//fmt.Println("cpuNum=", cpuNum)
	//
	////可以自己设置使用多个cpu
	//runtime.GOMAXPROCS(cpuNum - 1)
	//fmt.Println("ok")
}
