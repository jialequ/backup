package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem677(t *testing.T) {
	obj := Constructor()
	fmt.Printf(literal_0462, obj)
	obj.Insert("apple", 3)
	fmt.Printf(literal_0462, obj)
	fmt.Printf("obj.sum = %v\n", obj.Sum("ap"))
	obj.Insert("app", 2)
	fmt.Printf(literal_0462, obj)
	fmt.Printf("obj.sum = %v\n", obj.Sum("ap"))
}

const literal_0462 = "obj = %v\n"
