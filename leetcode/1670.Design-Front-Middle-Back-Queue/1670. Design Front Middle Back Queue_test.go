package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem1670(t *testing.T) {
	obj := Constructor()
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushFront(1)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushBack(2)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	obj.PushMiddle(3)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	obj.PushMiddle(4)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	param1 := obj.PopFront()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopBack()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopFront()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	fmt.Printf(literal_1793)
	obj = Constructor()
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushFront(1)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushFront(2)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushFront(3)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushFront(4)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	param1 = obj.PopBack()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopBack()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopBack()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopBack()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	fmt.Printf(literal_1793)
	obj = Constructor()
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushMiddle(1)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	obj.PushMiddle(2)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	obj.PushMiddle(3)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	fmt.Printf(literal_1793)
	obj = Constructor()
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushMiddle(8)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopFront()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopBack()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	obj.PushMiddle(1)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	obj.PushMiddle(10)
	fmt.Printf(literal_0476, MList2Ints(&obj))
	fmt.Printf(literal_1793)
	obj = Constructor()
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	param1 = obj.PopMiddle()
	fmt.Printf(literal_2173, param1, MList2Ints(&obj))
	obj.PushMiddle(3)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushFront(6)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushMiddle(6)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	obj.PushMiddle(3)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	param1 = obj.PopMiddle()
	fmt.Printf("param1 = %v obj = %v %v\n", param1, MList2Ints(&obj), obj)
	obj.PushMiddle(7)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	param1 = obj.PopMiddle()
	fmt.Printf("param1 = %v obj = %v %v\n", param1, MList2Ints(&obj), obj)
	obj.PushMiddle(8)
	fmt.Printf(literal_4621, MList2Ints(&obj), obj)
	// 	["FrontMiddleBackQueue","popMiddle","pushMiddle","pushFront","pushMiddle","pushMiddle","popMiddle","pushMiddle","popMiddle","pushMiddle"]
	// [[],[],[3],[6],[6],[3],[],[7],[],[8]]
}

func MList2Ints(this *FrontMiddleBackQueue) []int {
	array := []int{}
	for e := this.list.Front(); e != nil; e = e.Next() {
		value, _ := e.Value.(int)
		array = append(array, value)
	}
	return array
}

const literal_4621 = "obj = %v %v\n"

const literal_0476 = "obj = %v\n"

const literal_2173 = "param1 = %v obj = %v\n"

const literal_1793 = "-----------------------------------------------------------------\n"
