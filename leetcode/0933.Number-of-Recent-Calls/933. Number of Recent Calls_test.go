package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem933(t *testing.T) {
	obj := Constructor933()
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Ping(1)
	fmt.Printf(literal_3509, param1)
	param1 = obj.Ping(100)
	fmt.Printf(literal_3509, param1)
	param1 = obj.Ping(3001)
	fmt.Printf(literal_3509, param1)
	param1 = obj.Ping(3002)
	fmt.Printf(literal_3509, param1)
}

const literal_3509 = "param = %v\n"
