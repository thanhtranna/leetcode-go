package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

func Test_Problem0016(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{[]int{-1, 2, 1, -4}, 1},
			ans{2},
		},
		question{
			para{[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, 1},
			ans{0},
		},
		question{
			para{[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, 0},
			ans{0},
		},
		// If you need multiple tests, you can copy the above elements.
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, threeSumClosest(p.one, p.two), "input:%v", p)
	}
}
