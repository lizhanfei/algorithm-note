package v1

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
)

//pathSum 向下的路径节点之和
//leetcode : https://leetcode-cn.com/problems/6eUYwP/
//思路：
//深度
func pathSum(root *data.TreeNode, targetSum int) int {
	if nil == root {
		return 0
	}
	path := PathSum{}
	res, resNew := path.search(root)
	fmt.Println(res)
	fmt.Println(resNew)

	return path.getRes(res, resNew, targetSum)
}

type PathSum struct {
}

//深度优先递归
//获取所有的结果集
func (this *PathSum) search(root *data.TreeNode) ([][]int, [][]int) {
	var res = make([][]int, 0)
	var resNew = make([][]int, 0)
	//是否有子节点; 有子节点，则

	if nil != root.Left || nil != root.Right {
		if nil != root.Left {
			resTmp, resNewTmp := this.search(root.Left)
			//resNewTmp 和当前节点结合，形成新的
			res = append(res, resTmp...)
			res = append(res, resNewTmp...)
			for _, v := range resNewTmp {
				resNew = append(resNew, append(v, root.Val))
			}
		}
		if nil != root.Right {
			resTmp, resNewTmp := this.search(root.Right)
			//resNewTmp 和当前节点结合，形成新的
			res = append(res, resTmp...)
			res = append(res, resNewTmp...)
			for _, v := range resNewTmp {
				resNew = append(resNew, append(v, root.Val))
			}
		}
		resNew = append(resNew, []int{root.Val})
		return res, resNew
	} else {
		//没有子节点
		return res, append(resNew, []int{root.Val})
	}
}


func (this *PathSum) getRes(res, resNew [][]int, targetSum int) int {
	resTotal := 0
	for _, v := range res {
		if this.totalArr(v) == targetSum {
			resTotal++
		}
	}
	for _, v := range resNew {
		if this.totalArr(v) == targetSum {
			resTotal++
		}
	}
	return resTotal
}

func (this *PathSum) totalArr(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}