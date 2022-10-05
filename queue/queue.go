package queue

import (
	"log"

	"github.com/undervars/playground/linkedList/doubleLinkedList"
)

type Queue[T any] struct {
	l *doubleLinkedList.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		l: &doubleLinkedList.LinkedList[T]{},
	}
}

func (s *Queue[T]) Count() int {
	return s.l.Count()
}
func (s *Queue[T]) Push(value T) {
	s.l.PushBack(value)
}

func (s *Queue[T]) Pop() T {
	popped := s.l.PopFront()
	if popped == nil {
		log.Fatal("Nothing to pop")
	}
	return popped.Value
}
