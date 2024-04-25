package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem478(t *testing.T) {
	obj := Constructor(1, 0, 0)
	fmt.Printf(literal_4108, obj.RandPoint())
	fmt.Printf(literal_4108, obj.RandPoint())
	fmt.Printf(literal_4108, obj.RandPoint())

	obj = Constructor(10, 5, -7.5)
	fmt.Printf(literal_4108, obj.RandPoint())
	fmt.Printf(literal_4108, obj.RandPoint())
	fmt.Printf(literal_4108, obj.RandPoint())
}

const literal_4108 = "RandPoint() = %v\n"
