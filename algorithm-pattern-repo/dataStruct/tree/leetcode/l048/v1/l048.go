package v1

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
	"strconv"
	"strings"
)

//二叉树序列化和反序列化
//leetcode: https://leetcode-cn.com/problems/h54YBf/

//思路1
//数组标识二叉树  适合完全二叉树
//层序遍历二叉树
//根节点的数组位置时0
//左子节点的数组位置 = 父节点位置 * 2 + 1
//右子节点的数组位置 = 父节点位置 * 2 + 2
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *data.TreeNode) string {
	if nil == root {
		return ""
	}

	res := ""
	resTmp := ""
	queueNow := make([]*data.TreeNode, 0)
	queueNow = append(queueNow, root)
	queueTmp := make([]*data.TreeNode, 0)
	for {
		if 0 == len(queueNow) {
			break
		}
		hasNoNil := false
		for {
			if 0 == len(queueNow) {
				break
			}
			nodeNow := queueNow[0]
			queueNow = queueNow[1:]
			if nil == nodeNow {//当前节点时nil节点
				resTmp += "null,"
				queueTmp = append(queueTmp, nil)
				queueTmp = append(queueTmp, nil)
				continue
			} else {
				hasNoNil = true
				res += resTmp
				resTmp = ""
				res += fmt.Sprintf("%d,", nodeNow.Val)

				if nodeNow.Left != nil {
					queueTmp = append(queueTmp, nodeNow.Left)
				} else {
					queueTmp = append(queueTmp, nil)
				}
				if nodeNow.Right != nil {
					queueTmp = append(queueTmp, nodeNow.Right)
				} else {
					queueTmp = append(queueTmp, nil)
				}
			}
		}
		if hasNoNil {//有一个值不为nil，那么继续遍历下一层
			queueNow = append(queueNow, queueTmp...)
			queueTmp = queueTmp[:0]
		}
	}

	return strings.TrimRight(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data1 string) *data.TreeNode {
	if "" == data1 {
		return nil
	}
	intArr := strings.Split(data1, ",")

	fmt.Println(intArr)

	firstVal := this.str2Int(intArr[0])
	root := data.TreeNode{
		Val: firstVal,
	}
	queueNow := make([]*nodeWithIndex, 0)
	queueNow = append(queueNow, &nodeWithIndex{
		node:  &root,
		index: 0,
	})

	for {
		if 0 == len(queueNow) {
			break
		}
		nodeNow := queueNow[0]
		queueNow = queueNow[1:]
		leftIndex := nodeNow.index*2 + 1
		rightIndex := nodeNow.index*2 + 2

		if leftIndex < len(intArr) && intArr[leftIndex] != "null" { //左子节点有值
			leftTmp := data.TreeNode{
				Val:   this.str2Int(intArr[leftIndex]),
				Left:  nil,
				Right: nil,
			}
			nodeNow.node.Left = &leftTmp
			queueNow = append(queueNow, &nodeWithIndex{
				node:  &leftTmp,
				index: leftIndex,
			})
		}
		if rightIndex < len(intArr) && intArr[rightIndex] != "null" { //左子节点有值
			rightTmp := data.TreeNode{
				Val:   this.str2Int(intArr[rightIndex]),
				Left:  nil,
				Right: nil,
			}
			nodeNow.node.Right = &rightTmp
			queueNow = append(queueNow, &nodeWithIndex{
				node:  &rightTmp,
				index: rightIndex,
			})
		}
	}
	return &root
}

func (this *Codec) str2Int(intStr string) int {
	res, _ := strconv.Atoi(intStr)
	return res
}


type nodeWithIndex struct {
	node  *data.TreeNode
	index int
}