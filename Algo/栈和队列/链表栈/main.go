package main

import (
	"fmt"
	"sync"
)

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
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 如果栈顶为空，那么增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否则新元素插入链表的头部
		// 原来的链表
		preNode := stack.root
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v
		// 原来的链表链接到新元素后面
		newNode.Next = preNode
		// 将新节点放在头部
		stack.root = newNode
	}
	// 栈中元素数量+1
	stack.size = stack.size + 1
}

/**
出栈操作
*/
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}
	// 顶部元素要出栈
	value := stack.root.Value
	stack.root = stack.root.Next
	// 栈中元素数量-1
	stack.size = stack.size - 1
	return value
}

// 获取栈顶元素
func (stack *LinkStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}
	// 顶部元素值
	v := stack.root.Value
	return v
}

// 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}
