package treeTraversal

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var Node1 = Node{Val: 1,}
var Node2 = Node{Val: 2,}
var Node3 = Node{Val: 3,}
var Node4 = Node{Val: 4,}
var Node5 = Node{Val: 5,}
var Node6 = Node{Val: 6,}
var Node7 = Node{Val: 7,}
var Node8 = Node{Val: 8,}

func init() {
	Node5.Left = &Node3
	Node5.Right = &Node7

	Node7.Left = &Node6
	Node7.Right = &Node8

	Node3.Left = &Node2
	Node3.Right = &Node4

	Node2.Left = &Node1
}
