package stack

// [Stack]
//
// [Stack]: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29
type Stack[t any] struct {
	length int
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Length() int {
	return 0
}

func (s *Stack[T]) Push(item T) {}

func (s *Stack[T]) Pop() (value T, ok bool) {
	return
}

func (s *Stack[T]) Peek() (value T, ok bool) {
	return
}
