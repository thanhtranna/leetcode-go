package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	params
	answers
}

type params struct {
	nums   []int
	target int
}

type answers struct {
	one []int
}

func Test_Problem(t *testing.T) {

	qs := []question{
		{
			params{[]int{3, 2, 4}, 6},
			answers{[]int{1, 2}},
		},

		{
			params{[]int{3, 2, 4}, 5},
			answers{[]int{0, 1}},
		},

		{
			params{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			answers{[]int{1, 3}},
		},

		{
			params{[]int{0, 1}, 1},
			answers{[]int{0, 1}},
		},

		{
			params{[]int{0, 3}, 5},
			answers{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.answers, q.params
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
