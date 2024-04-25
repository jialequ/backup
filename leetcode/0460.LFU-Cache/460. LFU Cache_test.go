package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem460(t *testing.T) {
	obj := Constructor(5)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(1, 1)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(2, 2)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(3, 3)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(4, 4)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(5, 5)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)

	param1 := obj.Get(4)
	fmt.Printf(literal_0145, param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(4)
	fmt.Printf(literal_0145, param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(4)
	fmt.Printf(literal_0145, param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(5)
	fmt.Printf(literal_0145, param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(5)
	fmt.Printf(literal_0145, param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(5)
	fmt.Printf(literal_0145, param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(6, 6)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(7, 7)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(8, 8)
	fmt.Printf(literal_5160, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
}

func MList2Ints(lfu *LFUCache) map[int][][]int {
	res := map[int][][]int{}
	for k, v := range lfu.nodes {
		node := v.Value.(*node)
		arr := [][]int{}
		tmp := []int{node.key, node.value, node.frequency}
		arr = append(arr, tmp)
		res[k] = arr
	}
	return res
}

func MLists2Ints(lfu *LFUCache) map[int][]int {
	res := map[int][]int{}
	for k, v := range lfu.lists {
		tmp := []int{}
		for head := v.Front(); head != nil; head = head.Next() {
			tmp = append(tmp, head.Value.(*node).value)
		}
		res[k] = tmp
	}
	return res
}

const literal_5160 = "obj.list = %v obj.map = %v obj.min = %v\n"

const literal_0145 = "param_1 = %v obj.list = %v obj.map = %v obj.min = %v\n"
