package main

import (
	baseStack "algorithm-pattern-repo/dataStruct/stack/base"
	"algorithm-pattern-repo/dataStruct/tree/traversal/treeTraversal"
	"fmt"
)

//preorderTraversal 前序递归遍历。中->左->右
//时间复杂度 O(N) 要遍历所有的节点
func preorderTraversal(root *treeTraversal.Node) {
	if nil == root {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// preorderTraversal 通过非递归遍历
//非递归可以避免多次的函数栈分配，提高效率
//非递归需要借助栈实现，类似dfs
func preorderTraversalV2(root *treeTraversal.Node) []int {
	stack := baseStack.NewStack()
	res := make([]int, 0)
	stack.Push(root)
	for {
		if 0 == stack.Length() {
			break
		}
		nodeNow := stack.Pop().(*treeTraversal.Node)
		res = append(res, nodeNow.Val)
		if nil != nodeNow.Right {
			stack.Push(nodeNow.Right )
		}
		if nil != nodeNow.Left {
			stack.Push(nodeNow.Left )
		}
	}

	fmt.Println(res)
	return res
}
