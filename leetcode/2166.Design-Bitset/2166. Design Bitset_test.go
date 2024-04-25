package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem2166(t *testing.T) {
	obj := Constructor(5)
	fmt.Printf(literal_8437, obj)

	obj.Fix(3)
	fmt.Printf(literal_8437, obj)
	obj.Fix(1)
	fmt.Printf(literal_8437, obj)
	obj.Flip()
	fmt.Printf(literal_8437, obj)

	fmt.Printf("all = %v\n", obj.All())
	obj.Unfix(0)
	fmt.Printf(literal_8437, obj)
	obj.Flip()
	fmt.Printf(literal_8437, obj)

	fmt.Printf("one = %v\n", obj.One())
	obj.Unfix(0)
	fmt.Printf(literal_8437, obj)

	fmt.Printf("count = %v\n", obj.Count())
	fmt.Printf("toString = %v\n", obj.ToString())
}

const literal_8437 = "obj = %v\n"
