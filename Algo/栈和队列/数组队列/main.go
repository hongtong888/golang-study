package main

import (
	"fmt"
	"sync"
)

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

/**
添加元素
*/
func (queue *ArrayQueue) Add(value string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, value)
	queue.size = queue.size + 1
}

//栈大小
func (queue *ArrayQueue) Size() int {
	return queue.size
}

// 栈是否为空
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

/**
移除元素 先进先出
*/
func (queue *ArrayQueue) Remove() (value string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	value = queue.array[0]
	////移动数据
	//for i:=1; i<queue.size; i++ {
	//	queue.array[i-1] = queue.array[i]
	//}
	////缩减容量 但容量不会释放
	//queue.array = queue.array[0 : queue.size-1]

	//创建新数组进行替换
	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray
	queue.size = queue.size - 1
	return
}

func main() {
	arrayQueue := new(ArrayQueue)
	arrayQueue.Add("cat")
	arrayQueue.Add("dog")
	arrayQueue.Add("hen")
	fmt.Println("size:", arrayQueue.size)
	fmt.Println("pop:", arrayQueue.Remove())
	fmt.Println("pop:", arrayQueue.Remove())
	fmt.Println("size:", arrayQueue.Size())
	arrayQueue.Add("drag")
	fmt.Println("pop:", arrayQueue.Remove())
}
