package stack

type SliceStack[T any] struct {
	arr []T
}

func (s *SliceStack[T]) Count() int {
	return len(s.arr)
}

func (s *SliceStack[T]) Push(val T) {
	s.arr = append(s.arr, val)
}

func (s *SliceStack[T]) Pop() T {
	lastIndex := len(s.arr) - 1
	popped := s.arr[lastIndex]
	s.arr = s.arr[:lastIndex]

	return popped
}
