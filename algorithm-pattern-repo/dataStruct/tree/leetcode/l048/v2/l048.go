package v2

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	step int
	val  []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *data.TreeNode) string {
	res := ""

	stack := make([]*data.TreeNode, 0)
	stack = append(stack, root)

	for 0 != len(stack) {
		nodeNow := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if nil == nodeNow {
			res += "null,"
			continue
		}
		res += fmt.Sprintf("%d,", nodeNow.Val)
		stack = append(stack, nodeNow.Right)
		stack = append(stack, nodeNow.Left)
	}

	return strings.TrimRight(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data1 string) *data.TreeNode {
	if "" == data1 {
		return nil
	}
	this.step = 0
	this.val = strings.Split(data1, ",")
	if "null" == this.val[this.step] {
		return nil
	}
	root := data.TreeNode{
		Val:   this.str2Int(this.val[this.step]),
		Left:  nil,
		Right: nil,
	}

	root.Left = this.makeNode()
	root.Right = this.makeNode()
	return &root
}

func (this *Codec) makeNode() *data.TreeNode {
	this.step++//当前节点位置
	if "null" == this.val[this.step] {
		return nil
	}
	nodeNow := data.TreeNode{
		Val:   this.str2Int(this.val[this.step]),
		Left:  nil,
		Right: nil,
	}
	nodeNow.Left = this.makeNode()
	nodeNow.Right = this.makeNode()

	return &nodeNow
}

func (this *Codec) str2Int(intStr string) int {
	res, _ := strconv.Atoi(intStr)
	return res
}
