package minheap_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/minheap"
)

func TestMinHeap(t *testing.T) {
	heap := minheap.New()

	if heap.Length() != 0 {
		t.Fatal("expected length of new heap to be 0")
	}

	v, ok := heap.Delete()
	if ok {
		t.Fatalf("expected heap.Delete() on empty heap ok to be false, got value: %v, ok: %t", v, ok)
	}

    heap.Insert(5)
    heap.Insert(3)
    heap.Insert(69)
    heap.Insert(420)
    heap.Insert(4)
    heap.Insert(1)
    heap.Insert(8)
    heap.Insert(7)

	if heap.Length() != 8 {
		t.Fatal("expected length of heap to be 8 after inserts")
	}

	v, ok = heap.Delete()
	if v != 1 || !ok {
		t.Fatalf("expected heap.Delete() value to be 1, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = heap.Delete()
	if v != 3 || !ok {
		t.Fatalf("expected heap.Delete() value to be 3, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = heap.Delete()
	if v != 4 || !ok {
		t.Fatalf("expected heap.Delete() value to be 4, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = heap.Delete()
	if v != 5 || !ok {
		t.Fatalf("expected heap.Delete() value to be 5, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	if heap.Length() != 4 {
		t.Fatal("expected length of heap to be 4")
	}
    v, ok = heap.Delete()
	if v != 7 || !ok {
		t.Fatalf("expected heap.Delete() value to be 7, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = heap.Delete()
	if v != 8 || !ok {
		t.Fatalf("expected heap.Delete() value to be 8, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = heap.Delete()
	if v != 69 || !ok {
		t.Fatalf("expected heap.Delete() value to be 69, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = heap.Delete()
	if v != 420 || !ok {
		t.Fatalf("expected heap.Delete() value to be 420, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	if heap.Length() != 0 {
		t.Fatal("expected length of heap to be 0")
	}
}
