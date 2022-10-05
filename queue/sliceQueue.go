package queue

type SliceQueue[T any] struct {
	arr []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{
		arr: make([]T, 0),
	}
}

func (s *SliceQueue[T]) Count() int {
	return len(s.arr)
}
func (s *SliceQueue[T]) Push(value T) {
	s.arr = append(s.arr, value)
}

func (s *SliceQueue[T]) Pop() T {
	firstElement := s.arr[0]
	s.arr = s.arr[1:]

	return firstElement
}
