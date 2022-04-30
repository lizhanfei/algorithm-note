package main

import (
	baseStack "algorithm-pattern-repo/dataStruct/stack/base"
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
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

//inorderTraversal 中序递归遍历  左->中->右
//时间复杂度O(N) 要遍历所有的节点
func inorderTraversal(root *treeTraversal.Node) []int {
	res := make([]int, 0)
	if nil == root {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

//inorderTraversalV2 非递归中序遍历
//借助栈实现
//类似深度搜索
//先把左子树全部入栈
//之后逐个弹出，压入结果集
//之后处理右子树
//时间复杂度O(N) 不适用递归，可以避免多层的函数栈
func inorderTraversalV2(root *treeTraversal.Node) []int {
	stack := make([]*treeTraversal.Node, 0)
	res := make([]int, 0)

	nodeNow := root
	for len(stack) != 0 || nodeNow != nil {
		for nodeNow != nil {
			stack = append(stack, nodeNow)
			nodeNow = nodeNow.Left
		}//一直向左
		nodeNow = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, nodeNow.Val)
		nodeNow = nodeNow.Right
	}
	return res
}

//postorderTraversal 递归后序遍历 左->右->中
//leetcode: https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *treeTraversal.Node) []int {
	res := make([]int, 0)
	if nil == root {
		return res
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)

	return res
}
//postorderTraversalV2 非递归版遍历  左->右->中
//基于栈实现
//先一直递归左子树到叶子节点
//弹出节点 判断是否有右子树
//有右子树，则当前节点再次压入。处理右子树 同时记录右子树已处理
//leetcode: https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func postorderTraversalV2(root *treeTraversal.Node) []int {
	res := make([]int, 0)
	if nil == root {
		return res
	}
	stack := make([]*treeTraversal.Node, 0)
	handled := make(map[*treeTraversal.Node]int8)

	for 0 != len(stack) || nil != root {
		for nil != root {
			stack = append(stack, root)
			root = root.Left
		}
		nodeNow := stack[len(stack) - 1]
		if nodeNow.Right != nil  {//有右节点，但是没处理过；则把当前节点再次入栈。处理右节点
			if _, exist := handled[nodeNow.Right]; !exist {
				root = nodeNow.Right
				handled[nodeNow.Right] = 1
				continue
			}
		}
		stack = stack[:len(stack) - 1]
		res = append(res, nodeNow.Val)
	}
	return res
}

//zigzagLevelOrder 锯齿形遍历
// 每层遍历时，采用向前和向后两种方式遍历
// 每次遍历，需要把下级的结果放入到下次遍历的队列中

//奇数层：倒序遍历；倒序遍历上次遍历的结果；下层节点的需要从尾部压入; 压入时，注意右节点先入
//偶数层：需要倒着遍历；下层节点需要从尾部压入

//时间复杂度： O(N)  每个节点只需要遍历一次
//空间复杂度：O(N)  需要有两个暂存空间，存储下级节点列表；如果只有一层，那么暂存空间长度 = 节点数量
//leetcode: https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *treeTraversal.Node) [][]int {
	res := make([][]int, 0)
	if nil == root {
		return res
	}

	queueNow := make([]*treeTraversal.Node, 0)
	queueNow = append(make([]*treeTraversal.Node, 0), root)
	queueTmp := make([]*treeTraversal.Node, 0)
	deep := -1
	for {
		if 0 == len(queueNow) {
			break
		}
		deep++
		resTmp := make([]int, 0)
		if deep % 2 != 0 {//奇数层
			for {
				if 0 == len(queueNow) {
					break
				}
				root = queueNow[len(queueNow) - 1]
				queueNow = queueNow[:len(queueNow) - 1]
				resTmp = append(resTmp, root.Val)
				if root.Right != nil {
					queueTmp = append(queueTmp, root.Right)
				}
				if root.Left != nil {
					queueTmp = append(queueTmp, root.Left)
				}
			}
		} else {//偶数层
			for {
				if 0 == len(queueNow) {
					break
				}
				root = queueNow[len(queueNow) - 1]
				queueNow = queueNow[:len(queueNow) - 1]
				resTmp = append(resTmp, root.Val)
				if root.Left != nil {
					queueTmp = append(queueTmp, root.Left)
				}
				if root.Right != nil {
					queueTmp = append(queueTmp, root.Right)
				}
			}
		}
		queueNow = append(queueNow, queueTmp...) //优化一下空间。每次 queueTmp 不创建新的内存。把结果逐个复制给queueNow
		queueTmp = queueTmp[:0] //queueTmp 清空
		res = append(res, resTmp)
	}
	return res
}

//levelOrder 层序遍历
//逐层遍历。遍历的时候，把下层节点逐个的放入下级结果里
//leetcode: https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *treeTraversal.Node) [][]int {
	res := make([][]int, 0)
	if nil == root {
		return res
	}

	queueNow := make([]*treeTraversal.Node, 0)
	queueNow = append(make([]*treeTraversal.Node, 0), root)
	queueTmp := make([]*treeTraversal.Node, 0)
	for {
		if 0 == len(queueNow) {
			break
		}
		resTmp := make([]int, 0)
		for {
			if 0 == len(queueNow) {
				break
			}
			root = queueNow[0]
			queueNow = queueNow[1:]
			resTmp = append(resTmp, root.Val)
			if root.Left != nil {
				queueTmp = append(queueTmp, root.Left)
			}
			if root.Right != nil {
				queueTmp = append(queueTmp, root.Right)
			}
		}
		queueNow = append(queueNow, queueTmp...) //优化一下空间。每次 queueTmp 不创建新的内存。把结果逐个复制给queueNow
		queueTmp = queueTmp[:0] //queueTmp 清空
		res = append(res, resTmp)
	}
	return res
}

//从叶子层到根，倒序逐层遍历
// 想法1：逐层遍历；最终结果倒排序后返回
//leetcode: https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *treeTraversal.Node) [][]int {
	res := make([][]int, 0)
	if nil == root {
		return res
	}

	queueNow := make([]*treeTraversal.Node, 0)
	queueNow = append(make([]*treeTraversal.Node, 0), root)
	queueTmp := make([]*treeTraversal.Node, 0)
	for {
		if 0 == len(queueNow) {
			break
		}
		resTmp := make([]int, 0)
		for {
			if 0 == len(queueNow) {
				break
			}
			root = queueNow[0]
			queueNow = queueNow[1:]
			resTmp = append(resTmp, root.Val)
			if root.Left != nil {
				queueTmp = append(queueTmp, root.Left)
			}
			if root.Right != nil {
				queueTmp = append(queueTmp, root.Right)
			}
		}
		queueNow = append(queueNow, queueTmp...) //优化一下空间。每次 queueTmp 不创建新的内存。把结果逐个复制给queueNow
		queueTmp = queueTmp[:0] //queueTmp 清空
		res = append(res, resTmp)
	}
	//结果倒序
	lenRes := len(res)
	resBottom := make([][]int, 0, lenRes)
	for {
		if 0 == lenRes {
			break
		}
		resBottom = append(resBottom, res[lenRes - 1])
		res = res[:lenRes - 1]
		lenRes--
	}
	return resBottom
}

//从叶子层到根，倒序逐层遍历
// 想法2：逐层遍历，把待遍历的节点保存到二维数组中。然后倒序逐层遍历结果
//leetcode: https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottomV2(root *treeTraversal.Node) [][]int {
	res := make([][]int, 0)
	if nil == root {
		return res
	}
	nodeList := make([][]*treeTraversal.Node, 0)

	queueNow := make([]*treeTraversal.Node, 0)
	queueNow = append(queueNow, root)
	for {
		if 0 == len(queueNow) {
			break
		}
		nodeList = append(nodeList, queueNow)
		queueTmp := make([]*treeTraversal.Node, 0)
		for {
			if 0 == len(queueNow) {
				break
			}
			root = queueNow[0]
			queueNow = queueNow[1:]
			if root.Left != nil {
				queueTmp = append(queueTmp, root.Left)
			}
			if root.Right != nil {
				queueTmp = append(queueTmp, root.Right)
			}
		}
		queueNow = queueTmp
	}
	//倒序遍历nodeList
	lenRes := len(nodeList)
	for {
		if 0 == lenRes {
			break
		}
		resTmp := make([]int, 0, len(nodeList[lenRes - 1]))
		nodeListNow := nodeList[lenRes - 1]
		nodeList = nodeList[:lenRes - 1]
		//遍历当前节点
		for {
			if 0 == len(nodeListNow) {
				break
			}
			nodeNow := nodeListNow[0]
			nodeListNow = nodeListNow[1:]
			resTmp = append(resTmp, nodeNow.Val)
		}
		res = append(res, resTmp)
		lenRes--
	}

	return res
}

//increasingBST 二叉树展开
//leetcode: https://leetcode-cn.com/problems/NYBBNL/
//思路：
//中序遍历
func increasingBST(root *data.TreeNode) *data.TreeNode {
	var resNode *data.TreeNode
	if root == nil {
		return resNode
	}
	var nodeNow *data.TreeNode
	//
	var scan func(root *data.TreeNode)
	scan = func(root *data.TreeNode) {
		if nil == root {
			return
		}
		scan(root.Left)
		if resNode == nil {
			//第一个节点
			nodeNow = &data.TreeNode{
				Val:   root.Val,
				Left:  nil,
				Right: nil,
			}
			resNode = nodeNow
		} else {
			tmp := &data.TreeNode{
				Val:   root.Val,
				Left:  nil,
				Right: nil,
			}
			nodeNow.Right = tmp
			nodeNow = tmp
		}
		scan(root.Right)
	}
	scan(root)
	return resNode
}

//inorderSuccessor 二叉搜索树中的中序后继
// leetcode : https://leetcode-cn.com/problems/P5rCT8/
//思路：
//中序遍历，找到p值，返回后面的值
func inorderSuccessor(root *data.TreeNode, p *data.TreeNode) *data.TreeNode {
	if root == nil  || p == nil {
		return nil
	}

	var scan func(root *data.TreeNode)
	var needNext bool
	var next *data.TreeNode
	scan = func(root *data.TreeNode) {
		if nil == root {
			return
		}
		if next != nil {
			return
		}

		scan(root.Left)
		if needNext == true && next == nil {
			next = root
			return
		}
		if root == p {
			needNext = true
		}
		scan(root.Right)
	}
	scan(root)
	return next
}