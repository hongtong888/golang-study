package main

import (
	"fmt"
	"math/rand"
)

/**
带退出机制的缓冲器
*/
func main() {
	done := make(chan struct{})
	ch := GenerateInt(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}
