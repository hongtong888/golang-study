package main

import (
	"fmt"
	"sync"
)

/**
go 语言并发安全和锁机制
*/
var wp sync.WaitGroup
var mutex sync.Mutex

func main() {
	for i := 0; i < 10; i++ {
		wp.Add(1)
		go test()
	}
	wp.Wait()
}

var num int

func test() {
	mutex.Lock()
	num++
	fmt.Printf("当前num为：%d \n", num)
	mutex.Unlock()
	wp.Done()
}
