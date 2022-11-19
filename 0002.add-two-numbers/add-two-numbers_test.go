package leetcode

import (
	"fmt"
	"leetcode-go/structures"
	"testing"
)

type question2 struct {
	param
	answer
}

// param Is the parameter
// one Represents the first parameter
type param struct {
	one     []int
	another []int
}

// answer Is the answer
// one Represents the first answer
type answer struct {
	one []int
}

func Test_Problem2(t *testing.T) {

	qs := []question2{
		{
			param{[]int{}, []int{}},
			answer{[]int{}},
		},
		{
			param{[]int{1}, []int{1}},
			answer{[]int{2}},
		},
		{
			param{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			answer{[]int{2, 4, 6, 8}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			answer{[]int{2, 4, 6, 8, 0, 1}},
		},
		{
			param{[]int{1}, []int{9, 9, 9, 9, 9}},
			answer{[]int{0, 0, 0, 0, 0, 1}},
		},
		{
			param{[]int{9, 9, 9, 9, 9}, []int{1}},
			answer{[]int{0, 0, 0, 0, 0, 1}},
		},
		{
			param{[]int{2, 4, 3}, []int{5, 6, 4}},
			answer{[]int{7, 0, 8}},
		},
		{
			param{[]int{1, 8, 3}, []int{7, 1}},
			answer{[]int{8, 9, 3}},
		},
		// If you need multiple tests, you can copy the above elements.
	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")

	for _, q := range qs {
		_, p := q.answer, q.param
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(addTwoNumbers(structures.Ints2List(p.one), structures.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
