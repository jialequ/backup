package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem225(t *testing.T) {
	obj := Constructor225()
	fmt.Printf(literal_8739, obj)
	param5 := obj.Empty()
	fmt.Printf("param_5 = %v\n", param5)
	obj.Push(2)
	fmt.Printf(literal_8739, obj)
	obj.Push(10)
	fmt.Printf(literal_8739, obj)
	param2 := obj.Pop()
	fmt.Printf("param_2 = %v\n", param2)
	param3 := obj.Top()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj.Empty()
	fmt.Printf("param_4 = %v\n", param4)
}

const literal_8739 = "obj = %v\n"
