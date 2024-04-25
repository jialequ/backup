package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem1603(t *testing.T) {
	obj := Constructor(1, 1, 0)
	fmt.Printf(literal_6140, obj)
	fmt.Printf(literal_6140, obj.AddCar(1))
	fmt.Printf(literal_6140, obj.AddCar(2))
	fmt.Printf(literal_6140, obj.AddCar(3))
	fmt.Printf(literal_6140, obj.AddCar(1))
}

const literal_6140 = "obj = %v\n"
