package main

import (
	"algorithm-pattern-repo/dataStruct/tree/traversal/treeTraversal"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	preorderTraversal(&treeTraversal.Node5)
}


func TestPreorderTraversalV2(t *testing.T) {
	preorderTraversalV2(&treeTraversal.Node5)
}