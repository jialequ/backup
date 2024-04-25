package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem895(t *testing.T) {
	obj := Constructor895()
	fmt.Printf(literal_3167, obj)
	obj.Push(5)
	fmt.Printf(literal_3167, obj)
	obj.Push(7)
	fmt.Printf(literal_3167, obj)
	obj.Push(5)
	fmt.Printf(literal_3167, obj)
	obj.Push(7)
	fmt.Printf(literal_3167, obj)
	obj.Push(4)
	fmt.Printf(literal_3167, obj)
	obj.Push(5)
	fmt.Printf(literal_3167, obj)
	param1 := obj.Pop()
	fmt.Printf(literal_9305, param1)
	param1 = obj.Pop()
	fmt.Printf(literal_9305, param1)
	param1 = obj.Pop()
	fmt.Printf(literal_9305, param1)
	param1 = obj.Pop()
	fmt.Printf(literal_9305, param1)
	param1 = obj.Pop()
	fmt.Printf(literal_9305, param1)
	param1 = obj.Pop()
	fmt.Printf(literal_9305, param1)
}

const literal_3167 = "obj = %v\n"

const literal_9305 = "param_1 = %v\n"
