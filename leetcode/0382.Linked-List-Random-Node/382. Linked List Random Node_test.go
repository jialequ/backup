package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

func TestProblem382(t *testing.T) {
	header := structures.Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	obj := Constructor(header)
	fmt.Printf("obj = %v\n", structures.List2Ints(header))
	param1 := obj.GetRandom()
	fmt.Printf(literal_7263, param1)
	param1 = obj.GetRandom()
	fmt.Printf(literal_7263, param1)
	param1 = obj.GetRandom()
	fmt.Printf(literal_7263, param1)
	param1 = obj.GetRandom()
	fmt.Printf(literal_7263, param1)
	param1 = obj.GetRandom()
	fmt.Printf(literal_7263, param1)
	param1 = obj.GetRandom()
	fmt.Printf(literal_7263, param1)

	fmt.Printf("\n\n\n")
}

const literal_7263 = "param_1 = %v\n"
