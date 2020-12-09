package main

import (
	"fmt"
	"math/rand"
)

/**
最简单的带缓冲的生成器
*/
func main() {
	ch := GenerateInt()
	for v := range ch {
		fmt.Println(v)
	}
}

func GenerateInt() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
