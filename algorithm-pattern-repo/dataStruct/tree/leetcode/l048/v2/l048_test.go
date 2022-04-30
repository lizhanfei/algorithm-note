package v2

import (
	"algorithm-pattern-repo/dataStruct/tree/leetcode/data"
	"fmt"
	"testing"
)

func TestL048(t *testing.T) {
	data.InitL048()
	codec := Constructor()
	resStr := codec.serialize(&data.Node1)
	fmt.Println(resStr)

	res := codec.deserialize(resStr)
	fmt.Println(res)
}

