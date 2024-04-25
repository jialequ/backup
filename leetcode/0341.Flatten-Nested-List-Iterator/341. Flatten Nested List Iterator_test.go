package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem338(t *testing.T) {
	obj := Constructor([]*NestedInteger{})
	fmt.Printf(literal_1892, obj)
	fmt.Printf(literal_1892, obj.Next())
	fmt.Printf(literal_1892, obj.HasNext())
}

const literal_1892 = "obj = %v\n"
