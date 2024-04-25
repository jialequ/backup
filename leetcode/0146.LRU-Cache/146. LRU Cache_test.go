package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem146(t *testing.T) {
	obj := Constructor(2)
	fmt.Printf(literal_0874, MList2Ints(&obj))
	obj.Put(1, 1)
	fmt.Printf(literal_0874, MList2Ints(&obj))
	obj.Put(2, 2)
	fmt.Printf(literal_0874, MList2Ints(&obj))
	param1 := obj.Get(1)
	fmt.Printf(literal_1978, param1, MList2Ints(&obj))
	obj.Put(3, 3)
	fmt.Printf(literal_0874, MList2Ints(&obj))
	param1 = obj.Get(2)
	fmt.Printf(literal_1978, param1, MList2Ints(&obj))
	obj.Put(4, 4)
	fmt.Printf(literal_0874, MList2Ints(&obj))
	param1 = obj.Get(1)
	fmt.Printf(literal_1978, param1, MList2Ints(&obj))
	param1 = obj.Get(3)
	fmt.Printf(literal_1978, param1, MList2Ints(&obj))
	param1 = obj.Get(4)
	fmt.Printf(literal_1978, param1, MList2Ints(&obj))
}

func MList2Ints(lru *LRUCache) [][]int {
	res := [][]int{}
	head := lru.head
	for head != nil {
		tmp := []int{head.Key, head.Val}
		res = append(res, tmp)
		head = head.Next
	}
	return res
}

const literal_0874 = "obj = %v\n"

const literal_1978 = "param_1 = %v obj = %v\n"
