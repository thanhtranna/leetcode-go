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
	one string
}

// ans is the answer
// one represents the first answer
type ans struct {
	one int
}

func Test_Problem0013(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{"XXXIX"},
			ans{39},
		},
		question{
			para{"MDCCCLXXXVIII"},
			ans{1888},
		},
		question{
			para{"MCMLXXVI"},
			ans{1976},
		},
		question{
			para{"MMMCMXCIX"},
			ans{3999},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, romanToInt(p.one), "input: %v", p)
	}
}
