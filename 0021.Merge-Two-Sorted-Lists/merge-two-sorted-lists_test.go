package mergetwosortedlists

import (
	"fmt"
	"leetcode-go/structures"
	"testing"
)

type question struct {
	parameter
	answer
}

type parameter struct {
	one     []int
	another []int
}

type answer struct {
	one []int
}

func Test_Problem21(t *testing.T) {

	qs := []question{

		{
			parameter{[]int{}, []int{}},
			answer{[]int{}},
		},

		{
			parameter{[]int{1}, []int{1}},
			answer{[]int{1, 1}},
		},

		{
			parameter{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			answer{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			parameter{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			answer{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			parameter{[]int{1}, []int{9, 9, 9, 9, 9}},
			answer{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			parameter{[]int{9, 9, 9, 9, 9}, []int{1}},
			answer{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			parameter{[]int{2, 3, 4}, []int{4, 5, 6}},
			answer{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			parameter{[]int{1, 3, 8}, []int{1, 7}},
			answer{[]int{1, 1, 3, 7, 8}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 21------------------------\n")

	for _, q := range qs {
		_, p := q.answer, q.parameter
		fmt.Printf("【input】:%v【output】:%v\n", p, structures.List2Ints(mergeTwoLists(structures.Ints2List(p.one), structures.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
