package v1

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"math"
)

func maxPathSum(root *data.TreeNode) int {
	var maxSum = math.MinInt32

	var scan func(root *data.TreeNode) int
	scan = func(root *data.TreeNode) int {
		if root == nil {
			return 0 //如果当前节点为空，则最大贡献为0
		}
		maxLeft := scan(root.Left)
		maxRight := scan(root.Right)

		maxNow := root.Val

		if maxLeft > 0 {
			maxNow += maxLeft
		}
		if maxRight > 0 {
			maxNow += maxRight
		}

		if maxNow > maxSum {
			maxSum = maxNow
		}

		if maxLeft > 0 || maxRight > 0 {
			if maxLeft > maxRight {
				return root.Val + maxLeft
			} else {
				return root.Val + maxRight
			}
		}
		return root.Val
	}
	scan(root)

	return maxSum
}
