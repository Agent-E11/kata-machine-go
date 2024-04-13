package queue

type Queue[T any] struct {
	length int
}

func New[T any]() Queue[T] {
	return Queue[T]{} // TODO: Remove this
}

func (q *Queue[T]) Length() int {
	return 0 // TODO: Remove this
}

func (q *Queue[T]) Enqueue(item T) {}

func (q *Queue[T]) Dequeue() (value T, ok bool) {
	return // TODO: Remove this
}

func (q *Queue[T]) Peek() (value T, ok bool) {
	return // TODO: Remove this
}
