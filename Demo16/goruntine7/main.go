package main

import (
	"fmt"
	"sync"
)

/**
协程结合channal 编写0-120000的素数
*/
var wp sync.WaitGroup

func main() {

	//定义一个channel，用于存放数据,并启用存放数据的协程
	intChan := make(chan int, 10000)
	wp.Add(1)
	go putNum(intChan)

	//在定义一个channal，用于存放素数，并开启多个协程
	primeNum := make(chan int, 10000)
	flagChan := make(chan bool, 20)
	for i := 0; i < 16; i++ {
		wp.Add(1)
		go putPrimemNum(primeNum, intChan, flagChan)
	}

	wp.Add(1)
	go printPrimeNum(primeNum)

	//循环取出关闭标识，并关闭primeNum  channel
	wp.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-flagChan
		}
		close(primeNum)
		wp.Done()
	}()

	wp.Wait()
}

func printPrimeNum(primeNum chan int) {
	for num := range primeNum {
		fmt.Println(num)
	}
	wp.Done()
}

func putPrimemNum(primeNum chan int, intChan chan int, flagChan chan bool) {
	//从intChan遍历取出数据，在进行素数判断
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeNum <- num
		}
	}
	//close(primeNum)
	flagChan <- true
	wp.Done()
}

func putNum(intChan chan int) {
	for i := 2; i <= 100; i++ {
		intChan <- i
	}
	close(intChan)
	wp.Done()
}
