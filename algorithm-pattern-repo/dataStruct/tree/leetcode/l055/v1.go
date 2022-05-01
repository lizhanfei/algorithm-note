package l055

import "algorithm-pattern-repo/dataStruct/tree/leetcode/data"

type BSTIterator struct {
	val []int
	length int
	step int
}


func Constructor(root *data.TreeNode) BSTIterator {
	//中序遍历二叉树，获取
	bst := BSTIterator{val:make([]int, 0)}
	bst.InitVal(root)
	bst.length = len(bst.val)
	return bst
}

func (this *BSTIterator) InitVal(root *data.TreeNode) {
	if nil == root {
		return
	}
	this.InitVal(root.Left)
	this.val = append(this.val, root.Val)
	this.InitVal(root.Right)
}

func (this *BSTIterator) Next() int {
	res := this.val[this.step]
	this.step++
	return res
}


func (this *BSTIterator) HasNext() bool {
	if this.step < this.length{
		return  true
	}
	return false
}


