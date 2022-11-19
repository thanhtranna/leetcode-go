package generateparentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	parameter
	answer
}

type parameter struct {
	one int
}

type answer struct {
	one []string
}

func Test_Problem0022(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			parameter{2},
			answer{[]string{
				"(())",
				"()()",
			}},
		},
		{
			parameter{3},
			answer{[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			}},
		},
		{
			parameter{4},
			answer{[]string{
				"(((())))",
				"((()()))",
				"((())())",
				"((()))()",
				"(()(()))",
				"(()()())",
				"(()())()",
				"(())(())",
				"(())()()",
				"()((()))",
				"()(()())",
				"()(())()",
				"()()(())",
				"()()()()",
			}},
		},
	}

	for _, q := range qs {
		a, p := q.answer, q.parameter
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, generateParenthesis(p.one), "【output】::%v", p)
	}
}

func Benchmark_generateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(10)
	}
}
