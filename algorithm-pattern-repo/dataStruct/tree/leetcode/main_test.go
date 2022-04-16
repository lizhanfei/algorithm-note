package leetcode

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
	"testing"
)


func TestRecoverTree(t *testing.T) {
	data.InitRecoverTree()

	recoverTree(&data.Node2)

	val := data.Node2.Val

	fmt.Println(val)
}
