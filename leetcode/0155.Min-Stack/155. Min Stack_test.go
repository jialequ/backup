package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem155(t *testing.T) {
	obj1 := Constructor155()
	obj1.Push(1)
	fmt.Printf(literal_1546, obj1)
	obj1.Push(0)
	fmt.Printf(literal_1546, obj1)
	obj1.Push(10)
	fmt.Printf(literal_1546, obj1)
	obj1.Pop()
	fmt.Printf(literal_1546, obj1)
	param3 := obj1.Top()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj1.GetMin()
	fmt.Printf("param_4 = %v\n", param4)
}

const literal_1546 = "obj1 = %v\n"
