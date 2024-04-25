package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem1656(t *testing.T) {
	obj := Constructor(5)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Insert(3, "ccccc")
	fmt.Printf(literal_0156, param1, obj)
	param1 = obj.Insert(1, "aaaaa")
	fmt.Printf(literal_0156, param1, obj)
	param1 = obj.Insert(2, "bbbbb")
	fmt.Printf(literal_0156, param1, obj)
	param1 = obj.Insert(5, "eeeee")
	fmt.Printf(literal_0156, param1, obj)
	param1 = obj.Insert(4, "ddddd")
	fmt.Printf(literal_0156, param1, obj)
}

const literal_0156 = "param_1 = %v obj = %v\n"
