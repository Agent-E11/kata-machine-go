package ringbuffer

// [Circular Buffer]
//
// [Circular Buffer]: https://en.wikipedia.org/wiki/Circular_buffer
type RingBuffer[T any] struct {
	length int
}

func New[T any]() RingBuffer[T] {
	return RingBuffer[T]{}
}

func (rb *RingBuffer[T]) Length() int {
	return 0
}

func (rb *RingBuffer[T]) Push(item T) {}

func (rb *RingBuffer[T]) Get(idx int) (value T, ok bool) {
	return
}

func (rb *RingBuffer[T]) Pop() (value T, ok bool) {
	return
}
