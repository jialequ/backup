package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem208(t *testing.T) {
	obj := Constructor208()
	fmt.Printf(literal_2657, obj)
	obj.Insert("apple")
	fmt.Printf(literal_2657, obj)
	param1 := obj.Search("apple")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.Search("app")
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
	param3 := obj.StartsWith("app")
	fmt.Printf("param_3 = %v obj = %v\n", param3, obj)
	obj.Insert("app")
	fmt.Printf(literal_2657, obj)
	param4 := obj.Search("app")
	fmt.Printf("param_4 = %v obj = %v\n", param4, obj)
}

const literal_2657 = "obj = %v\n"
