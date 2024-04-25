package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem1600(t *testing.T) {
	obj := Constructor("king")
	fmt.Printf(literal_3146, obj)
	obj.Birth("king", "andy")
	fmt.Printf(literal_3146, obj)
	obj.Birth("king", "bob")
	fmt.Printf(literal_3146, obj)
	obj.Birth("king", "catherine")
	fmt.Printf(literal_3146, obj)
	obj.Birth("andy", "matthew")
	fmt.Printf(literal_3146, obj)
	obj.Birth("bob", "alex")
	fmt.Printf(literal_3146, obj)
	obj.Birth("bob", "asha")
	fmt.Printf(literal_3146, obj)
	param2 := obj.GetInheritanceOrder()
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
	obj.Death("bob")
	fmt.Printf(literal_3146, obj)
	param2 = obj.GetInheritanceOrder()
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
}

const literal_3146 = "obj = %v\n"
