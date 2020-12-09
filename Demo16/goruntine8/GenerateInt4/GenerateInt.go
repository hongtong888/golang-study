package main

import (
	"fmt"
	"math/rand"
)

/**
带并发、缓冲、退出机制的缓冲器
*/
func main() {
	done := make(chan struct{})
	ch := GenerateInt(done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	done <- struct{}{}
}

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)
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

func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
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

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Lable:
		for {
			select {
			case ch <- <-GenerateIntA(send):
			case ch <- <-GenerateIntB(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}
