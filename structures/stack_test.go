package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	ast := assert.New(t)

	s := NewStack()
	ast.True(s.IsEmpty(), "Check whether the newly created s is empty")

	start, end := 0, 100

	for i := start; i < end; i++ {
		s.Push(i)
		ast.Equal(i-start+1, s.Len(), "Check the length of q after Push.")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, s.Pop(), "Pop out the number from s.")
	}

	ast.True(s.IsEmpty(), "Check whether s after Pop is empty")
}
