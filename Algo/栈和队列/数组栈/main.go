package main

import (
	"fmt"
	"sync"
)

/**
栈：先进后出，先进队的数据最后才出来  数组形式下的压栈 出栈
*/

type ArrayStack struct {
	array []string   //底层可变数组切片
	size  int        //数据容量大小
	lock  sync.Mutex //并发安全锁
}

/**
入栈操作
*/
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	//刚入切片中
	stack.array = append(stack.array, v)
	//容量+1
	stack.size = stack.size + 1
}

/**
出栈操作
*/
func (stack *ArrayStack) Pop() (v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}
	//栈顶元素
	v = stack.array[stack.size-1]
	//切片容量变化
	//1、切片偏移量，表明最后元素不属于该切片，但是切片的容量并不会回收，时间复杂度为O(1),空间复杂度较高
	stack.array = stack.array[0 : stack.size-1]

	//2、创建新的数组进行替换
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray
	//元素个数减一
	stack.size = stack.size - 1
	return v
}

/**
获取栈顶元素
*/
func (stack *ArrayStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}
	value := stack.array[stack.size-1]
	return value
}

//栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

//测试
func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
