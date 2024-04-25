package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem232(t *testing.T) {
	obj := Constructor232()
	fmt.Printf(literal_6234, obj)
	obj.Push(2)
	fmt.Printf(literal_6234, obj)
	obj.Push(10)
	fmt.Printf(literal_6234, obj)
	param2 := obj.Pop()
	fmt.Printf("param_2 = %v\n", param2)
	param3 := obj.Peek()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj.Empty()
	fmt.Printf("param_4 = %v\n", param4)
}

const literal_6234 = "obj = %v\n"
