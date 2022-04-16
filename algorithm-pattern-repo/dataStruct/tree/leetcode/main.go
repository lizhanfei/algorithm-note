package leetcode

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
)

//recoverTree 恢复二叉树
//leetcode: https://leetcode-cn.com/problems/recover-binary-search-tree/
//思路：中序遍历二叉树
//遍历过程中，如果出现了降序，则记录当前位置
//有一个降序点，则将降序点的两个元素置换
//有两个降序点，则将第一个降序点的第一个元素和第二个降序点的第二个元素置换
func recoverTree(root *data.TreeNode)  {
	//使用栈实现中序遍历

	stack := make([]*data.TreeNode, 0)
	levelPoint := make([][2]*data.TreeNode, 0) //记录降序点
	res := make([]*data.TreeNode, 0)

	for 0 != len(stack) || nil != root {
		//左节点遍历到叶子
		for nil != root {
			stack = append(stack, root)
			root = root.Left
		}
		//弹出节点
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if 0 != len(res) &&  root.Val < res[len(res) - 1].Val {
			//降序对
			levelPoint = append(levelPoint, [2]*data.TreeNode{res[len(res) - 1], root})
		}
		res = append(res, root)
		//右子树不为空，则处理右子树
		root = root.Right
	}
	if 1 == len(levelPoint) {
		//一个降序对
		tmpVal := levelPoint[0][0].Val
		levelPoint[0][0].Val = levelPoint[0][1].Val
		levelPoint[0][1].Val = tmpVal
	} else if 2 == len(levelPoint) {
		tmpVal := levelPoint[0][0].Val
		levelPoint[0][0].Val = levelPoint[1][1].Val
		levelPoint[1][1].Val = tmpVal
	}
}