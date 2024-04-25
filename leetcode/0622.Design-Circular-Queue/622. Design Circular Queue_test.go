package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem622(t *testing.T) {
	obj := Constructor(3)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.EnQueue(1)
	fmt.Printf(literal_3905, param1, obj)
	param2 := obj.DeQueue()
	fmt.Printf(literal_3905, param2, obj)
	param3 := obj.Front()
	fmt.Printf(literal_3905, param3, obj)
	param4 := obj.Rear()
	fmt.Printf(literal_3905, param4, obj)
	param5 := obj.IsEmpty()
	fmt.Printf(literal_3905, param5, obj)
	param6 := obj.IsFull()
	fmt.Printf(literal_3905, param6, obj)
}

const literal_3905 = "param_1 = %v obj = %v\n"
