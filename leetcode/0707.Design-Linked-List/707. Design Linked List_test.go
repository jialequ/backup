package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem707(t *testing.T) {
	obj := Constructor()
	fmt.Printf(literal_7013, MList2Ints(&obj))
	param1 := obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	obj.AddAtHead(38)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtHead(45)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.DeleteAtIndex(2)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtIndex(1, 24)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtTail(36)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtIndex(3, 72)
	obj.AddAtTail(76)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtHead(7)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtHead(36)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtHead(34)
	fmt.Printf(literal_7013, MList2Ints(&obj))
	obj.AddAtTail(91)
	fmt.Printf(literal_7013, MList2Ints(&obj))

	fmt.Printf("\n\n\n")

	obj1 := Constructor()
	fmt.Printf(literal_9653, MList2Ints(&obj1))
	param2 := obj1.Get(0)
	fmt.Printf(literal_6984, param2, MList2Ints(&obj1))
	obj1.AddAtIndex(1, 2)
	fmt.Printf(literal_9653, MList2Ints(&obj1))
	param2 = obj1.Get(0)
	fmt.Printf(literal_6984, param2, MList2Ints(&obj1))
	param2 = obj1.Get(1)
	fmt.Printf(literal_6984, param2, MList2Ints(&obj1))
	obj1.AddAtIndex(0, 1)
	fmt.Printf(literal_9653, MList2Ints(&obj1))
	param2 = obj1.Get(0)
	fmt.Printf("param_1 = %v obj1 = %v\n", param2, MList2Ints(&obj1))
	param2 = obj1.Get(1)
	fmt.Printf(literal_6984, param2, MList2Ints(&obj1))
}

func MList2Ints(head *MyLinkedList) []int {
	res := []int{}
	cur := head.head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

const literal_7013 = "obj = %v\n"

const literal_9653 = "obj1 = %v\n"

const literal_6984 = "param_2 = %v obj1 = %v\n"
