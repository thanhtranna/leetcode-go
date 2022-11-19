package leetcode

import (
	"fmt"
	"testing"
)

type question3 struct {
	param
	answer
}

// para Is the parameter
// one Represents the first parameter
type param struct {
	s string
}

// ans Is the answer
// one Represents the first answer
type answer struct {
	one int
}

func Test_Problem3(t *testing.T) {

	qs := []question3{

		{
			param{"abcabcbb"},
			answer{3},
		},
		{
			param{"bbbbb"},
			answer{1},
		},
		{
			param{"pwwkew"},
			answer{3},
		},
		{
			param{""},
			answer{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3------------------------\n")

	for _, q := range qs {
		_, p := q.answer, q.param
		fmt.Printf("【input】:%v       【output】:%v\n", p, lengthOfLongestSubstring(p.s))
		// fmt.Printf("【input】:%v       【output】:%v\n", p, lengthOfLongestSubstringV2(p.s))
	}
	fmt.Printf("\n\n\n")
}
