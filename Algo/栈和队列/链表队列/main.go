package main

import "sync"

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

//新增
func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	//如果队列根节点为nil 则新增为跟节点
	if queue.root == nil {
		newNode := new(LinkNode)
		newNode.Value = v
	} else {
		//不是跟节点，则插入到末尾节点
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		// 新节点放在链表尾部
		nowNode.Next = newNode
	}
	// 队中元素数量+1
	queue.size = queue.size + 1
}

func main() {}
