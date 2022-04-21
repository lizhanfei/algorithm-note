package leetcode

import "algorithm-pattern-repo/dataStruct/tree/leetcode/data"

type CBTInserter struct {
	root *data.TreeNode
}

func Constructor(root *data.TreeNode) CBTInserter {
	return CBTInserter{root: root}
}

//思路
//广度优先遍历
//发现第一个子节点为空的节点
//则构造一个节点，挂载该节点下
func (this *CBTInserter) Insert(v int) int {
	//借助队列实现
	queue := data.NewQueue()
	queue.Push(&data.QueueNode{
		Val:  this.root,
		Next: nil,
	})
	for {
		if 0 == queue.Len() {
			break
		}
		nodeNow := queue.Shift()
		if nil == nodeNow.Val.Left {
			nodeNow.Val.Left = &data.TreeNode{
				Val:   v,
				Left:  nil,
				Right: nil,
			}
			return nodeNow.Val.Val
		}
		if nil == nodeNow.Val.Right {
			nodeNow.Val.Right = &data.TreeNode{
				Val:   v,
				Left:  nil,
				Right: nil,
			}
			return nodeNow.Val.Val
		}
		queue.Push(&data.QueueNode{
			Val:  nodeNow.Val.Left,
			Next: nil,
		})
		queue.Push(&data.QueueNode{
			Val:  nodeNow.Val.Right,
			Next: nil,
		})
	}
	return 0
}

func (this *CBTInserter) Get_root() *data.TreeNode {
	return this.root
}

