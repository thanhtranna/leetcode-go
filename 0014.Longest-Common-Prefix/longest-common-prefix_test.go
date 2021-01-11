package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para is the parameter
// one represents the first parameter
type para struct {
	one []string
}

// ans is the answer
// one represents the first answer
type ans struct {
	one string
}

func Test_Problem0014(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[]string{"abcdd", "abcde", "ab"},
			},
			ans{"ab"},
		},
		question{
			para{
				[]string{"abcdd", "abcde"},
			},
			ans{"abcd"},
		},
		question{
			para{
				[]string{"flower", "flow", "flight"},
			},
			ans{"fl"},
		},
		question{
			para{
				[]string{},
			},
			ans{""},
		},
		// If you need multiple tests, you can copy the above element.
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, longestCommonPrefix(p.one), "input:%v", p)
	}
}
