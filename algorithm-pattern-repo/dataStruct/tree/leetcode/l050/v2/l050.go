package v2

import "algorithm-pattern-repo/dataStruct/tree/leetcode/data"

func pathSum(root *data.TreeNode, targetSum int) int {
	if nil == root {
		return 0
	}
	ret := 0
	ret += search(root, targetSum)
	ret += pathSum(root.Left, targetSum)
	ret += pathSum(root.Right, targetSum)
	return ret
}
//
func search(root *data.TreeNode, targetSum int) int {
	if nil == root {
		return 0
	}
	ret := 0
	if root.Val == targetSum {
		ret++
	}
	ret += search(root.Left, targetSum - root.Val)
	ret += search(root.Right, targetSum - root.Val)
	return ret
}