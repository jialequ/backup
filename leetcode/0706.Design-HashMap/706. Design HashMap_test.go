package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem706(t *testing.T) {
	obj := Constructor706()
	obj.Put(7, 10)
	fmt.Printf("Get 7 = %v\n", obj.Get(7))
	obj.Put(7, 20)
	fmt.Printf("Contains 7 = %v\n", obj.Get(7))
	param1 := obj.Get(100)
	fmt.Printf(literal_9365, param1)
	obj.Remove(100007)
	param1 = obj.Get(7)
	fmt.Printf(literal_9365, param1)
	obj.Remove(7)
	param1 = obj.Get(7)
	fmt.Printf(literal_9365, param1)
}

const literal_9365 = "param1 = %v\n"
