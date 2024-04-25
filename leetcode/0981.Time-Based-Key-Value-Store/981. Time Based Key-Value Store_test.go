package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem981(t *testing.T) {
	obj := Constructor981()
	obj.Set("foo", "bar", 1)
	fmt.Printf(literal_8973, obj.Get("foo", 1))
	fmt.Printf(literal_8973, obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Printf(literal_8973, obj.Get("foo", 4))
	fmt.Printf(literal_8973, obj.Get("foo", 5))
}

const literal_8973 = "Get = %v\n"
