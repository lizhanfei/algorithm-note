package data

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push(&QueueNode{
		Val:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Next: nil,
	})
	queue.Push(&QueueNode{
		Val:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Next: nil,
	})
	queue.Push(&QueueNode{
		Val:  &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Next: nil,
	})

	fmt.Println(queue.Len())

	queue.Shift()
	queue.Shift()

	fmt.Println(queue.Len())
}
