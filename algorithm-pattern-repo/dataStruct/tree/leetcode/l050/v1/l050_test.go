package v1

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
	"testing"
)

func TestL050(t *testing.T) {
	data.InitL050()

	fmt.Println(pathSum(&data.Node1, 8))
}


func TestL050V1(t *testing.T) {
	data.Node1.Val = 0
	data.Node2.Val = 1
	data.Node3.Val = 1
	data.Node1.Left = &data.Node2
	data.Node1.Right = &data.Node3

	fmt.Println(pathSum(&data.Node1, 1))
}


func TestL050V2(t *testing.T) {
	data.InitL050V2()

	fmt.Println(pathSum(&data.Node1, 3))
}
