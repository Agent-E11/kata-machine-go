package minheap

// [Heap]
//
// [Heap]: https://en.wikipedia.org/wiki/Heap_%28data_structure%29
type MinHeap struct {
	length int
}

func New() MinHeap {
	return MinHeap{} // TODO: Remove this
}

func (mh *MinHeap) Length() int {
	return 0
}

func (mh *MinHeap) Insert(value int) {}

func (mh *MinHeap) Delete() (value int, ok bool) {
	return
}
