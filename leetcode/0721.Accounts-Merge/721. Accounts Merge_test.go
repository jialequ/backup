package leetcode

import (
	"fmt"
	"testing"
)

type question721 struct {
	para721
	ans721
}

// para 是参数
// one 代表第一个参数
type para721 struct {
	w [][]string
}

// ans 是答案
// one 代表第一个答案
type ans721 struct {
	one [][]string
}

func TestProblem721(t *testing.T) {

	qs := []question721{

		{
			para721{[][]string{{"John", literal_7319, "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", literal_7319, "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"}}},
			ans721{[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", literal_7319},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"}}},
		},

		{
			para721{[][]string{{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
				{"Ethan", literal_9108, literal_9108, "Ethan0@m.co"},
				{"Kevin", "Kevin4@m.co", literal_4786, literal_4786},
				{"Gabe", "Gabe0@m.co", literal_1739, literal_2165},
				{"Gabe", literal_1739, "Gabe4@m.co", literal_2165}}},
			ans721{[][]string{{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", literal_9108},
				{"Gabe", "Gabe0@m.co", literal_2165, literal_1739, "Gabe4@m.co"},
				{"Kevin", literal_4786, "Kevin4@m.co"}}},
		},

		{
			para721{[][]string{{"David", literal_6829, literal_3640, literal_7064},
				{"David", literal_8501, literal_8501, literal_6829},
				{"David", literal_9238, literal_3640, literal_6829},
				{"David", literal_6829, literal_9238, literal_7064},
				{"David", literal_3640, literal_9238, literal_7064}}},
			ans721{[][]string{{"David", literal_6829, literal_9238, literal_7064, literal_3640, literal_8501}}},
		},

		{
			para721{[][]string{}},
			ans721{[][]string{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 721------------------------\n")

	for _, q := range qs {
		_, p := q.ans721, q.para721
		fmt.Printf("【input】:%v       【output】:%v\n", p, accountsMerge(p.w))
	}
	fmt.Printf("\n\n\n")
}

const literal_7319 = "johnsmith@mail.com"

const literal_9108 = "Ethan3@m.co"

const literal_4786 = "Kevin2@m.co"

const literal_1739 = "Gabe3@m.co"

const literal_2165 = "Gabe2@m.co"

const literal_6829 = "David0@m.co"

const literal_3640 = "David4@m.co"

const literal_7064 = "David3@m.co"

const literal_8501 = "David5@m.co"

const literal_9238 = "David1@m.co"
