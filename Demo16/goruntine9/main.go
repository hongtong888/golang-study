package main

import (
	"fmt"
	"sync"
	"time"
)

/**
协程中的异常捕获
*/
var wp sync.WaitGroup

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello,world")
	}
	wp.Done()
}

//函数
func test() {
	//这里我们可以使用defer + recover 来捕获异常
	defer func() {
		wp.Done()
		err := recover()
		if err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error

}

func main() {
	wp.Add(1)
	go sayHello()

	wp.Add(1)
	go test()

	wp.Wait()

}
