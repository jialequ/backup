package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

func TestProblem173(t *testing.T) {

	root := structures.Ints2TreeNode([]int{9, 7, 15, 3, structures.NULL, structures.NULL, 20})
	obj := Constructor173(root)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Next()
	fmt.Printf(literal_6814, param1)
	param2 := obj.HasNext()
	fmt.Printf("param_2 = %v\n", param2)
	param1 = obj.Next()
	fmt.Printf(literal_6814, param1)
	param1 = obj.Next()
	fmt.Printf(literal_6814, param1)
	param1 = obj.Next()
	fmt.Printf(literal_6814, param1)
	param1 = obj.Next()
	fmt.Printf(literal_6814, param1)
	param2 = obj.HasNext()
	fmt.Printf("param_2 = %v\n", param2)
}

const literal_6814 = "param_1 = %v\n"
