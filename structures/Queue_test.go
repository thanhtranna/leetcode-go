package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()
	ast.True(q.IsEmpty(), "Check whether the newly created queue is empty")

	start, end := 0, 100

	for i := start; i < end; i++ {
		q.Push(i)
		ast.Equal(i-start+1, q.Len(), "Check the length of queue after Push.")
	}

	for i := start; i < end; i++ {
		ast.Equal(i, q.Pop(), "Pop out from queue")
	}

	ast.True(q.IsEmpty(), "Check whether the queue after Pop is empty")
}
