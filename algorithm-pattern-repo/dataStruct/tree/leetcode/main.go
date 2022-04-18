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
func recoverTree(root *data.TreeNode) {
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
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if 0 != len(res) && root.Val < res[len(res)-1].Val {
			//降序对
			levelPoint = append(levelPoint, [2]*data.TreeNode{res[len(res)-1], root})
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

//isSymmetric 二叉树是否对称
//leetcode: https://leetcode-cn.com/problems/symmetric-tree/
//思路：
// 根节点，直接判断左右是否相同
// 左右节点判断 左右节点是否相同 且 左节点的左子节点 是否和 右节点的右子节点相同  && 左节点的右子节点 是否和 右节点的左子节点相同
//时间复杂度：O(N)  要遍历整棵树
//空间复杂度：O(N)  递归过程中，会复制每个节点到下次递归
func isSymmetric(root *data.TreeNode) bool {
	return isSymmetricLevel(root.Left, root.Right)
}

func isSymmetricLevel(left *data.TreeNode, right *data.TreeNode) bool {
	if nil == left && nil == right{//左右有一个等于nil，则返回true
		return true
	}
	if (nil == left && nil != right) || (nil != left && nil == right) { //只有一个等于nil，则返回false
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricLevel(left.Left, right.Right) && isSymmetricLevel(left.Right, right.Left)
}

//非递归处理二叉树是否对称
//leetcode: https://leetcode-cn.com/problems/symmetric-tree/
// 思路：
//借助一个二维队列，队列的元素是一个长度为2的数组
//起始，把两个root节点组成数组压入队列
//每次遍历，从队列取出元素，进行对比。 对比不相同，则返回false
//对比之后，把0节点的左子节点和1元素的右子节点组成数组压入队列; 把0元素的右子节点和 1元素的左子节点组成数组压入队列
func isSymmetricV2(root *data.TreeNode) bool {
	queue := make([][2]*data.TreeNode, 0)

	queue = append(queue, [2]*data.TreeNode{root, root})
	for {
		if 0 == len(queue) {
			break
		}
		nodeNow := queue[0]
		queue = queue[1:]

		//对比 nodeNow 元素
		if nil == nodeNow[0] && nil == nodeNow[1]{//左右有一个等于nil，则返回true
			return true
		}
		if (nil == nodeNow[0] && nil != nodeNow[1]) || (nil != nodeNow[0] && nil == nodeNow[1]) { //只有一个等于nil，则返回false
			return false
		}
		if nodeNow[0].Val != nodeNow[1].Val {
			return false
		}
		queue =append(queue, [2]*data.TreeNode{nodeNow[0].Left, nodeNow[1].Right})
		queue =append(queue, [2]*data.TreeNode{nodeNow[0].Right, nodeNow[1].Left})
	}

	return true
}
