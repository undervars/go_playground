package stack

import (
	"log"

	"github.com/undervars/playground/linkedList/doubleLinkedList"
)

type Stack[T any] struct {
	l *doubleLinkedList.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		l: &doubleLinkedList.LinkedList[T]{},
	}
}

func (s *Stack[T]) Count() int {
	return s.l.Count()
}
func (s *Stack[T]) Push(value T) {
	s.l.PushBack(value)
}

func (s *Stack[T]) Pop() T {
	popped := s.l.PopBack()
	if popped == nil {
		log.Fatal("Nothing to pop")
	}
	return popped.Value
}
