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
var Node9 = TreeNode{Val: 9,}


func InitAny() {
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


func InitIsSymmetric() {
	Node1.Val = 9
	Node2.Val = -42
	Node3.Val = -42
	Node4.Val = 76
	Node5.Val = 76
	Node6.Val = 13
	Node7.Val = 13


	Node1.Left = &Node2
	Node1.Right = &Node3
	Node2.Left = &Node4
	Node3.Right = &Node5
	Node4.Right = &Node6
	Node5.Left = &Node7
}


func InitRightSideView() {
	Node1.Left = &Node2
	Node1.Right = &Node3
	Node3.Right = &Node4
}


func InitSumNumbers() {
	Node2.Val = 0
	Node1.Left = &Node2
	//Node1.Right = &Node3
}

func InitL048() {
	Node1.Val = 1
	Node2.Val = 2
	Node3.Val = 3
	Node4.Val = 4
	Node5.Val = 5
	Node6.Val = 6
	Node7.Val = 7

	Node1.Left = &Node2
	Node1.Right = &Node3
	Node3.Left = &Node4
	Node3.Right = &Node5
	Node4.Left = &Node6
	Node4.Right = &Node7
}

func InitL050() {
	Node1.Val = 10
	Node2.Val = 5
	Node3.Val = -3
	Node4.Val = 3
	Node5.Val = 2
	Node6.Val = 11
	Node7.Val = 3
	Node8.Val = -2
	Node9.Val = 1


	Node1.Left = &Node2
	Node1.Right = &Node3
	Node3.Right = &Node6

	Node2.Left = &Node4
	Node2.Right = &Node5
	Node4.Left = &Node7
	Node4.Right = &Node8

	Node5.Right = &Node9
}

func InitL050V2() {
	Node1.Val = 1
	Node2.Val = 2
	Node3.Val = 3
	Node4.Val = 4
	Node5.Val = 5

	Node1.Right = &Node2
	Node2.Right = &Node3
	Node3.Right = &Node4
	Node4.Right = &Node5
}