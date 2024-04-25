package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem352(t *testing.T) {
	obj := Constructor352()
	fmt.Printf("obj = %v\n", obj)
	obj.AddNum(1)
	fmt.Printf(literal_8930, obj.GetIntervals()) // [1,1]
	obj.AddNum(3)
	fmt.Printf(literal_8930, obj.GetIntervals()) // [1,1] [3,3]
	obj.AddNum(7)
	fmt.Printf(literal_8930, obj.GetIntervals()) // [1, 1], [3, 3], [7, 7]
	obj.AddNum(2)
	fmt.Printf(literal_8930, obj.GetIntervals()) // [1, 3], [7, 7]
	obj.AddNum(6)
	fmt.Printf(literal_8930, obj.GetIntervals()) // [1, 3], [6, 7]
}

const literal_8930 = "Intervals = %v\n"
