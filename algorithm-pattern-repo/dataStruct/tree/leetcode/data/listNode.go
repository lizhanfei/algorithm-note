package data

type ListNode struct {
	Val  int
	Next *ListNode
}

var ListNode1 = ListNode{Val: 0,}
var ListNode2 = ListNode{Val: 1,}
var ListNode3 = ListNode{Val: 2,}
var ListNode4 = ListNode{Val: 3,}
var ListNode5 = ListNode{Val: 4,}
var ListNode6 = ListNode{Val: 5,}

func InitSortedList() {
	ListNode1.Next = &ListNode2
	ListNode2.Next = &ListNode3
	ListNode3.Next = &ListNode4
	ListNode4.Next = &ListNode5
	ListNode5.Next = &ListNode6
}
