package main

import (
	"fmt"
	"sync"
	"time"
)

/**
需求：要统计1-120000的数字中那些是素数？for循环实现  和协程实现
*/
var wg sync.WaitGroup

func main() {
	//1、使用普通的for循环
	start := time.Now().Unix()
	//for num := 2; num <=120000; num++ {
	//	var flag =true
	//	for i := 2; i< num; i++ {
	//		if num%i == 0 {
	//			flag = false
	//			break
	//		}
	//	}
	//	if flag {
	//		fmt.Println(num, "是素数")
	//	}
	//}

	//2、使用协程实现
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start) //4毫秒
}

func test(n int) {
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				fmt.Println(num, "是素数")
			}
		}
	}
	wg.Done()
}
