package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem307(t *testing.T) {
	obj := Constructor307([]int{1, 3, 5})
	fmt.Printf(literal_7123, obj)
	fmt.Printf(literal_9271, obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Printf(literal_7123, obj)
	fmt.Printf(literal_9271, obj.SumRange(0, 2))

	obj = Constructor307([]int{-1})
	fmt.Printf(literal_7123, obj)
	fmt.Printf(literal_9271, obj.SumRange(0, 0))
	obj.Update(0, 1)
	fmt.Printf(literal_7123, obj)
	fmt.Printf(literal_9271, obj.SumRange(0, 0))

	obj = Constructor307([]int{7, 2, 7, 2, 0})
	fmt.Printf(literal_7123, obj)
	obj.Update(4, 6)
	obj.Update(0, 2)
	obj.Update(0, 9)
	fmt.Printf(literal_9271, obj.SumRange(4, 4))
	obj.Update(3, 8)
	fmt.Printf(literal_7123, obj)
	fmt.Printf(literal_9271, obj.SumRange(0, 4))
	obj.Update(4, 1)
	fmt.Printf(literal_9271, obj.SumRange(0, 3))
	fmt.Printf(literal_9271, obj.SumRange(0, 4))
	obj.Update(0, 4)
}

const literal_7123 = "obj = %v\n"

const literal_9271 = "SumRange(0,2) = %v\n"
