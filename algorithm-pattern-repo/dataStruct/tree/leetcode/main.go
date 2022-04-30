package leetcode

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
	"math"
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
	if nil == left && nil == right { //左右有一个等于nil，则返回true
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
		if nil == nodeNow[0] && nil == nodeNow[1] { //左右有一个等于nil，则返回true
			return true
		}
		if (nil == nodeNow[0] && nil != nodeNow[1]) || (nil != nodeNow[0] && nil == nodeNow[1]) { //只有一个等于nil，则返回false
			return false
		}
		if nodeNow[0].Val != nodeNow[1].Val {
			return false
		}
		queue = append(queue, [2]*data.TreeNode{nodeNow[0].Left, nodeNow[1].Right})
		queue = append(queue, [2]*data.TreeNode{nodeNow[0].Right, nodeNow[1].Left})
	}

	return true
}

//maxDepth  二叉树的最大深度
//leetcode: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
//思路： 递归解决
//每递归一层，深度增加一
//时间复杂度 O(N) 需要遍历所有的节点
// 空间复杂度
func maxDepth(root *data.TreeNode) int {
	if nil == root {
		return 0
	}
	deepNow := 1
	var maxDeep = 1
	recursionDepthTree(root, deepNow, &maxDeep)
	return maxDeep
}

func recursionDepthTree(nodeNow *data.TreeNode, deepNow int, maxDeep *int) {
	if deepNow > *maxDeep {
		*maxDeep = deepNow
	}
	if nil != nodeNow.Left {
		deepNowLeft := deepNow + 1
		recursionDepthTree(nodeNow.Left, deepNowLeft, maxDeep)
	}
	if nil != nodeNow.Right {
		deepNowRight := deepNow + 1
		recursionDepthTree(nodeNow.Right, deepNowRight, maxDeep)
	}
}

//maxDepthV2 二叉树的最大深度
//leetcode: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// 思路： 迭代
// 套用深度优先搜索的模式
// 遍历的节点在入栈时，保存一个二维数组，一个保存节点，一个保存深度
// 每次入栈时，使用当前节点的深度+1，保存新节点的深度。
// 同时获取最大的深度
type NodeWithDeep struct {
	deep int
	node *data.TreeNode
}

func maxDepthV2(root *data.TreeNode) int {
	if nil == root {
		return 0
	}
	var maxDeep = 1
	var stack = make([]*NodeWithDeep, 0) //基于队列实现栈
	stack = append(stack, &NodeWithDeep{
		deep: 1,
		node: root,
	})

	var nodeNow *NodeWithDeep
	for {
		if 0 == len(stack) {
			break
		}
		nodeNow = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if nodeNow.deep > maxDeep {
			maxDeep = nodeNow.deep
		}
		if nodeNow.node.Left != nil {
			stack = append(stack, &NodeWithDeep{
				deep: nodeNow.deep + 1,
				node: nodeNow.node.Left,
			})
		}
		if nodeNow.node.Right != nil {
			stack = append(stack, &NodeWithDeep{
				deep: nodeNow.deep + 1,
				node: nodeNow.node.Right,
			})
		}
	}
	return maxDeep
}

//sortedListToBST 有序链表转换二叉搜索树
//leetcode: https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
//思路： 寻找链表的中位节点作为根节点，将尾节点作为跟节点的右节点
//递归从头结点到 中位节点，建立左子树
//从中位节点到尾节点，建立右子树
type queueTask struct {
	parent     *data.TreeNode
	parentType int // 1->left 2->right
	left       int //左边界
	right      int //右边界
}

func sortedListToBST(head *data.ListNode) *data.TreeNode {
	if head == nil {
		return nil
	}
	//借助一个切片，存储链表的每一个节点
	listNodeArr := make([]*data.ListNode, 0)
	for {
		if head == nil {
			break
		}
		listNodeArr = append(listNodeArr, head)
		head = head.Next
	}
	//迭代。创建一个切片。每次取节点的中位节点作为根节点
	queue := make([]*queueTask, 0)
	//取queue的中位节点作为根节点
	middle := len(listNodeArr) / 2
	root := &data.TreeNode{
		Val: listNodeArr[middle].Val,
	}
	queue = append(queue, &queueTask{
		parent:     root,
		parentType: 1,
		left:       0,
		right:      middle - 1,
	})
	queue = append(queue, &queueTask{
		parent:     root,
		parentType: 2,
		left:       middle + 1,
		right:      len(listNodeArr) - 1,
	})
	for {
		//退出条件。
		if 0 == len(queue) {
			break
		}
		nodeTmp := queue[0]
		queue = queue[1:]
		//根据 nodeTmp 构造节点
		if nodeTmp.left > nodeTmp.right {
			continue
		}
		middle := int(math.Ceil(float64(nodeTmp.right-nodeTmp.left)/2.0)) + nodeTmp.left
		//当前节点只有一个
		tmp := data.TreeNode{
			Val: listNodeArr[middle].Val,
		}
		switch nodeTmp.parentType {
		case 1:
			nodeTmp.parent.Left = &tmp
		case 2:
			nodeTmp.parent.Right = &tmp
		}
		queue = append(queue, &queueTask{
			parent:     &tmp,
			parentType: 1,
			left:       nodeTmp.left,
			right:      middle - 1,
		})
		queue = append(queue, &queueTask{
			parent:     &tmp,
			parentType: 2,
			left:       middle + 1,
			right:      nodeTmp.right,
		})
	}

	return root
}

//largestValues 二叉树每层的最大值
//leetcode: https://leetcode-cn.com/problems/hPov7L/
//思路：
// 参考二叉树层序遍历 https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//每一层遍历时，记录下最大值
func largestValues(root *data.TreeNode) []int {
	maxRes := make([]int, 0)
	if nil == root {
		return maxRes
	}
	queueNow := make([]*data.TreeNode, 0)
	queueNow = append(make([]*data.TreeNode, 0), root)

	queueTmp := make([]*data.TreeNode, 0)
	for {
		if 0 == len(queueNow) {
			break
		}
		maxValue := math.MinInt32
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
			if root.Val > maxValue {
				maxValue = root.Val
			}
		}
		queueNow = append(queueNow, queueTmp...)
		queueTmp = queueTmp[:0]
		maxRes = append(maxRes, maxValue)
	}
	return maxRes
}

//findBottomLeftValue  二叉树最底层最左边的值
//leetcode: https://leetcode-cn.com/problems/LwUNpT/
//思路：
//深度优先遍历
//遍历中记录当前深度和节点值
//遍历到叶子节点时，判断深度是否大于已知深度，大于已知深度，则将当前叶子节点的值设为结果
type nodeWithDeep struct {
	node *data.TreeNode
	deep int
}

func findBottomLeftValue(root *data.TreeNode) int {
	res := 0
	maxDeep := math.MinInt32

	stack := make([]*nodeWithDeep, 0)
	stack = append(stack, &nodeWithDeep{
		node: root,
		deep: 1,
	})

	for 0 != len(stack) {
		nodeNow := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if nodeNow.node.Left != nil || nodeNow.node.Right != nil { //非叶子节点
			if nodeNow.node.Right != nil {
				stack = append(stack, &nodeWithDeep{
					node: nodeNow.node.Right,
					deep: nodeNow.deep + 1,
				})
			}
			if nodeNow.node.Left != nil {
				stack = append(stack, &nodeWithDeep{
					node: nodeNow.node.Left,
					deep: nodeNow.deep + 1,
				})
			}
		} else {                        //叶子节点
			if nodeNow.deep > maxDeep { //超过当前最大深度
				res = nodeNow.node.Val
				maxDeep = nodeNow.deep
			}
		}
	}

	return res
}

//rightSideView 二叉树的右侧视图
//leetcode: https://leetcode-cn.com/problems/WNC0Lk/
//思路：
//层序遍历，每层从右向左遍历
//每层取第一个压入结果集
func rightSideView(root *data.TreeNode) []int {
	res := make([]int, 0)
	if nil == root {
		return res
	}

	queueNow := make([]*data.TreeNode, 0)
	queueNow = append(queueNow, root)

	queueTmp := make([]*data.TreeNode, 0)
	var resTmp int
	for {
		if 0 == len(queueNow) {
			break
		}
		resTmp = -200
		for {
			if 0 == len(queueNow) {
				break
			}
			root = queueNow[0]
			queueNow = queueNow[1:]
			if resTmp == -200 {
				resTmp = root.Val
			}
			if root.Right != nil {
				queueTmp = append(queueTmp, root.Right)
			}
			if root.Left != nil {
				queueTmp = append(queueTmp, root.Left)
			}
		}
		queueNow = append(queueNow, queueTmp...)
		queueTmp = queueTmp[:0]
		res = append(res, resTmp)
	}
	return res
}

//pruneTree 二叉树剪枝
//leetcode https://leetcode-cn.com/problems/pOCWxh/
// 思路：
//计算每个节点子树的节点和
//如果某个节点的左右子节点不全为0，那么把为0的节点减掉
//如果左右节点均为0.当前节点不为0，那么把左右节点均减掉;否则，不处理
//
func pruneTree(root *data.TreeNode) *data.TreeNode {
	if 0 == handleTreeNode(root) {
		root = nil
	}
	return root
}

//handleTreeNode 处理子树，并返回从root往下所有节点的和
func handleTreeNode(root *data.TreeNode) int {
	var leftVal, rightVal int
	if nil != root.Left {
		leftVal = handleTreeNode(root.Left)
	}
	if nil != root.Right {
		rightVal = handleTreeNode(root.Right)
	}
	if leftVal != 0 || rightVal != 0 { //子树不全为0，那么剪除为0 的子树
		if leftVal == 0 {
			root.Left = nil
		}
		if rightVal == 0 {
			root.Right = nil
		}

	} else if root.Val != 0 {
		root.Left = nil
		root.Right = nil
	}

	return root.Val + leftVal + rightVal
}

type NodeWithValNow struct {
	node   *data.TreeNode
	valNow int //当前的值
}

func sumNumbers(root *data.TreeNode) int {
	if nil == root {
		return 0
	}
	res := make([]int, 0)
	var stack = make([]*NodeWithValNow, 0) //基于队列实现栈
	stack = append(stack, &NodeWithValNow{
		valNow: root.Val,
		node:   root,
	})

	var nodeNow *NodeWithValNow
	for {
		if 0 == len(stack) {
			break
		}
		nodeNow = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if nodeNow.node.Left != nil {
			stack = append(stack, &NodeWithValNow{
				valNow: nodeNow.valNow*10 + nodeNow.node.Left.Val,
				node:   nodeNow.node.Left,
			})
		}
		if nodeNow.node.Right != nil {
			stack = append(stack, &NodeWithValNow{
				valNow: nodeNow.valNow*10 + nodeNow.node.Right.Val,
				node:   nodeNow.node.Right,
			})
		}
		if nil == nodeNow.node.Left && nil == nodeNow.node.Right { //叶子节点
			res = append(res, nodeNow.valNow)
		}
	}
	fmt.Println(res)
	resReturn := 0
	for _, v := range res {
		resReturn += v
	}
	return resReturn
}
