package data


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Node1 = TreeNode{Val: 1,}
var Node2 = TreeNode{Val: 2,}
var Node3 = TreeNode{Val: 3,}
var Node4 = TreeNode{Val: 4,}
var Node5 = TreeNode{Val: 5,}
var Node6 = TreeNode{Val: 6,}
var Node7 = TreeNode{Val: 7,}
var Node8 = TreeNode{Val: 8,}


func initAny() {
	Node5.Left = &Node3
	Node5.Right = &Node7

	Node7.Left = &Node6
	Node7.Right = &Node8

	Node3.Left = &Node2
	Node3.Right = &Node4

	Node2.Left = &Node1
}



func InitRecoverTree() {
	Node2.Left = &Node3
	Node2.Right = &Node1
}
