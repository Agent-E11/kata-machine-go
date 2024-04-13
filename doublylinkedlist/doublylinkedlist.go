package doublylinkedlist

type DoublyLinkedList[T comparable] struct {
	length int
}

func New[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{} // TODO: Remove this
}

func (dll DoublyLinkedList[T]) Length() int {
	return 0
}

func (dll *DoublyLinkedList[T]) Prepend(item T) {}

func (dll *DoublyLinkedList[T]) InsertAt(item T, idx int) {}

func (dll *DoublyLinkedList[T]) Append(item T) {}

func (dll *DoublyLinkedList[T]) Remove(item T) (value T, ok bool) {
	return
}

func (dll *DoublyLinkedList[T]) Get(idx int) (value T, ok bool) {
	return
}

func (dll *DoublyLinkedList[T]) RemoveAt(idx int) (value T, ok bool) {
	return
}
