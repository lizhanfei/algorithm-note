package data

type QueueNode struct {
	Val  *TreeNode
	Next *QueueNode
}
type queue struct {
	head   *QueueNode
	end    *QueueNode
	length int
}

func NewQueue() queue {
	return queue{
		head:   nil,
		end:    nil,
		length: 0,
	}
}

func (this *queue) Push(node *QueueNode) {
	if this.head == nil {
		this.head = node
	}
	if this.end == nil {
		this.end = node
	} else { //如果end不是初始状态，那么把node挂到 end后。end指向最后一个
		this.end.Next = node
		this.end = this.end.Next
	}
	this.length++
}

//获取头节点
func (this *queue) Shift() *QueueNode {
	res := this.head
	if this.length > 0 { //需要向后移动
		if 1 == this.length { //只有一个元素
			this.head = nil
			this.end = nil
			this.length = 0
		} else {
			this.head = this.head.Next
			this.length--
		}
	}
	return res
}

func (this *queue) Len() int {
	return this.length
}
