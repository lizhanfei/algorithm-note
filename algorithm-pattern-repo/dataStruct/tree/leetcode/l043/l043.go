package l043

import "algorithm-pattern-repo/dataStruct/tree/leetcode/data"

type CBTInserter struct {
	root  *data.TreeNode
	queue []*data.TreeNode
}

func Constructor(root *data.TreeNode) CBTInserter {
	queue := make([]*data.TreeNode, 0)
	queueTmp := make([]*data.TreeNode, 0)
	queueTmp = append(queueTmp, root)
	var hasLeft, hasRight bool
	for {
		if 0 == len(queueTmp) {
			break
		}
		hasLeft = false
		hasRight = false
		nodNow := queueTmp[0]
		queueTmp = queueTmp[1:]

		if nodNow.Left != nil {
			hasLeft = true
			queueTmp = append(queueTmp, nodNow.Left)
		}
		if nodNow.Right != nil {
			hasRight = true
			queueTmp = append(queueTmp, nodNow.Right)
		}
		if !hasLeft || !hasRight {
			queue = append(queue, nodNow)
		}
	}
	return CBTInserter{root: root, queue: queue}
}

func (this *CBTInserter) Insert(v int) int {
	newNode := &data.TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
	firstNode := this.queue[0]
	if nil == firstNode.Left {
		firstNode.Left = newNode
	} else {
		firstNode.Right = newNode
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, newNode)
	return firstNode.Val
}

func (this *CBTInserter) Get_root() *data.TreeNode {
	return this.root
}
