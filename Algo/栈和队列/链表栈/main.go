package main

import "sync"

/**
链表形式下的下压栈
*/

type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

/**
入栈操作
*/
func (stack *LinkStack) Push(v string) {

}

func main() {

}
