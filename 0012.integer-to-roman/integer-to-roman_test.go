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
	one int
}

// ans is the answer
// one represents the first answer
type ans struct {
	one string
}

func Test_Problem0012(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{39},
			ans{"XXXIX"},
		},
		question{
			para{1888},
			ans{"MDCCCLXXXVIII"},
		},
		question{
			para{1976},
			ans{"MCMLXXVI"},
		},
		question{
			para{3999},
			ans{"MMMCMXCIX"},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, intToRoman(p.one), "input: %v", p)
	}
}
