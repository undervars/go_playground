package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/undervars/playground/stack"
)

func TestSliceStack(t *testing.T) {

	s := stack.SliceStack[int]{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	p1 := s.Pop()
	assert.Equal(t, 4, p1)
	p2 := s.Pop()
	assert.Equal(t, 3, p2)
	p3 := s.Pop()
	assert.Equal(t, 2, p3)
	assert.Equal(t, 1, s.Count())
	p4 := s.Pop()
	assert.Equal(t, 0, s.Count())
	assert.Equal(t, 1, p4)
}
