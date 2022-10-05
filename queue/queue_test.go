package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/undervars/playground/queue"
)

func TestPushPop(t *testing.T) {
	q := queue.New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	// s.Pop()
	p1 := q.Pop()
	assert.Equal(t, 1, p1)
	p2 := q.Pop()
	assert.Equal(t, 2, p2)
	p3 := q.Pop()
	assert.Equal(t, 3, p3)
	assert.Equal(t, 1, q.Count())
	p4 := q.Pop()
	assert.Equal(t, 0, q.Count())
	assert.Equal(t, 4, p4)
}

func BenchmarkLinkedListQueue(b *testing.B) {
	s := queue.New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	s := queue.NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
