package leetcode

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
	"math"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	data.InitRecoverTree()

	recoverTree(&data.Node2)

	val := data.Node2.Val

	fmt.Println(val)
}

func TestIsSymmetric(t *testing.T) {
	data.InitIsSymmetric()

	res := isSymmetric(&data.Node1)
	if res {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	data.InitIsSymmetric()

	for n := 0; n < b.N; n++ {
		_ = isSymmetric(&data.Node1)
		//if res {
		//	fmt.Println("true")
		//} else {
		//	fmt.Println("false")
		//}
	}
}

func BenchmarkIsSymmetricV2(b *testing.B) {
	data.InitIsSymmetric()

	for n := 0; n < b.N; n++ {
		_ = isSymmetricV2(&data.Node1)
		//if res {
		//	fmt.Println("true")
		//} else {
		//	fmt.Println("false")
		//}
	}
}

func TestMaxDepth(t *testing.T) {
	data.InitIsSymmetric()

	fmt.Println(maxDepth(&data.Node1))
}


func TestMaxDepthV2(t *testing.T) {
	data.InitIsSymmetric()

	fmt.Println(maxDepthV2(&data.Node1))
}


func BenchmarkMaxDepthV2(b *testing.B) {
	data.InitIsSymmetric()
	for n := 0; n < b.N; n++ {
		maxDepthV2(&data.Node1)
	}
}


func TestSortedListToBST(t *testing.T) {
	data.InitSortedList()

	res := sortedListToBST(&data.ListNode1)
	fmt.Println(res)
}

func TestSortedListToBSTV1(t *testing.T) {
	fmt.Println(int(math.Ceil(1.0 / 2.0)))
	//fmt.Println(sortedListToBST(&data.ListNode1))
}