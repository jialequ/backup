package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem703(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Printf(literal_1785, obj.Add(3))
	fmt.Printf(literal_1785, obj.Add(5))
	fmt.Printf(literal_1785, obj.Add(10))
	fmt.Printf(literal_1785, obj.Add(9))
	fmt.Printf(literal_1785, obj.Add(4))
}

const literal_1785 = "Add 7 = %v\n"
