package singlylinkedlist

// [Linked List]
//
// [Linked List]: https://en.wikipedia.org/wiki/Linked_list
type SinglyLinkedList[T comparable] struct {
	length int
}

func New[T comparable]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{}
}

func (dll SinglyLinkedList[T]) Length() int {
	return 0
}

func (dll *SinglyLinkedList[T]) Prepend(item T) {}

func (dll *SinglyLinkedList[T]) InsertAt(item T, idx int) {}

func (dll *SinglyLinkedList[T]) Append(item T) {}

func (dll *SinglyLinkedList[T]) Remove(item T) (value T, ok bool) {
	return
}

func (dll *SinglyLinkedList[T]) Get(idx int) (value T, ok bool) {
	return
}

func (dll *SinglyLinkedList[T]) RemoveAt(idx int) (value T, ok bool) {
	return
}
