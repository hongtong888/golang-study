package main

import (
	"fmt"
	"sync"
)

/**
多协程案例
*/
var wp sync.WaitGroup

func test(num int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("协程(%v)打印的第%v条数据\n", num, i)
	}
	defer wp.Done()
}

func main() {
	for i := 1; i <= 5; i++ {
		wp.Add(1)
		test(i)
	}
	wp.Wait()
}
