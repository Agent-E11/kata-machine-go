package arraylist_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/arraylist"
)

func TestArrayList(t *testing.T) {
	list := arraylist.New[int](3)

	list.Append(5)
	list.Append(7)
	list.Append(9)

	v, ok := list.Get(2)
	if v != 9 || !ok {
		t.Fatalf("expected list.Get(2) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = list.RemoveAt(1)
	if v != 7 || !ok {
		t.Fatalf("expected list.RemoveAt(1) to be 7, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	if list.Length != 2 {
		t.Fatalf("expected list.Length to be 2. Got length: %v", list.Length)
	}

	list.Append(11)
	v, ok = list.RemoveAt(1)
	if v != 9 || !ok {
		t.Fatalf("expected list.RemoveAt(1) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = list.Remove(9)
	if ok {
		t.Fatalf("expected list.Remove(9), ok to be false. Got value: %v, ok: %t", v, ok)
	}
	v, ok = list.RemoveAt(0)
	if v != 5 || !ok {
		t.Fatalf("expected list.RemoveAt(0) to be 5, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = list.RemoveAt(0)
	if v != 11 || !ok {
		t.Fatalf("expected list.RemoveAt(0) to be 11, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	if list.Length != 0 {
		t.Fatalf("expected list.Length to be 0. Got length: %v", list.Length)
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	v, ok = list.Get(2)
	if v != 5 || !ok {
		t.Fatalf("expected list.Get(2) to be 5, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = list.Get(0)
	if v != 5 || !ok {
		t.Fatalf("expected list.Get(0) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = list.Remove(9)
	if v != 9 || !ok {
		t.Fatalf("expected list.Remove(9) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	if list.Length != 2 {
		t.Fatalf("expected list.Length to be 2. Got length: %v", list.Length)
	}
	v, ok = list.Get(0)
	if v != 7 || !ok {
		t.Fatalf("expected list.Get(0) to be 7, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
}
