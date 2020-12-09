package main

import (
	"fmt"
	"math/rand"
)

/**
通知退出机制
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
