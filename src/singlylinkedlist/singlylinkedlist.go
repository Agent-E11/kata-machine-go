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

func (sll SinglyLinkedList[T]) Length() int {
	return 0
}

func (sll *SinglyLinkedList[T]) Prepend(item T) {}

func (sll *SinglyLinkedList[T]) InsertAt(item T, idx int) {}

func (sll *SinglyLinkedList[T]) Append(item T) {}

func (sll *SinglyLinkedList[T]) Remove(item T) (value T, ok bool) {
	return
}

func (sll *SinglyLinkedList[T]) Get(idx int) (value T, ok bool) {
	return
}

func (sll *SinglyLinkedList[T]) RemoveAt(idx int) (value T, ok bool) {
	return
}
