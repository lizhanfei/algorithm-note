package main

import (
	"algorithm-pattern-repo/dataStruct/tree/traversal/treeTraversal"
	"fmt"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	preorderTraversal(&treeTraversal.Node5)
}


func TestPreorderTraversalV2(t *testing.T) {
	preorderTraversalV2(&treeTraversal.Node5)
}

func TestInorderTraversal(t *testing.T) {
	res := inorderTraversal(&treeTraversal.Node5)
	fmt.Println(res)

	res = inorderTraversal(nil)
	fmt.Println(res)
}

func TestInorderTraversalV2(t *testing.T) {
	res := inorderTraversalV2(&treeTraversal.Node5)
	fmt.Println(res)
}

func TestPostorderTraversal(t *testing.T) {
	res := postorderTraversal(&treeTraversal.Node5)
	fmt.Println(res)
}

func TestPostorderTraversalV2(t *testing.T) {
	res := postorderTraversalV2(&treeTraversal.Node5)
	fmt.Println(res)
}

func TestZigzagLevelOrder(t *testing.T) {
	res := zigzagLevelOrder(&treeTraversal.Node5)
	fmt.Println(res)
}

func TestLevelOrder(t *testing.T) {
	res := levelOrder(&treeTraversal.Node5)
	fmt.Println(res)
}


func TestLevelOrderBottom(t *testing.T) {
	res := levelOrderBottom(&treeTraversal.Node5)
	fmt.Println(res)
}


func TestLevelOrderBottomV2(t *testing.T) {
	res := levelOrderBottomV2(&treeTraversal.Node5)
	fmt.Println(res)
}

