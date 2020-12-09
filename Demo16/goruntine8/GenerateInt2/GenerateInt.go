package main

import (
	"fmt"
	"math/rand"
)

/**
多个goruntine增强生成器
*/
func main() {
	ch := GenerateInt()
	for v := range ch {
		fmt.Println(v)
	}
}

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func GenerateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

/**
通过select扇入合并管道 增加生成的随机源
*/
func GenerateInt() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			select {
			case ch <- <-GenerateIntA():
			case ch <- <-GenerateIntB():
			}
		}
	}()
	return ch
}
