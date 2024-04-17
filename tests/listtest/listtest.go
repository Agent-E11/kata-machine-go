package listtest

import (
	"testing"
)

type list[T comparable] interface {
	Length() int

	Prepend(item T)

	InsertAt(item T, idx int)

	Append(item T)

	Remove(item T) (value T, ok bool)

	Get(idx int) (value T, ok bool)

	RemoveAt(idx int) (value T, ok bool)
}

func TestList(t *testing.T, l list[int]) {
	l.Append(5)
	l.Append(7)
	l.Append(9)

	v, ok := l.Get(2)
	if v != 9 || !ok {
		t.Fatalf("expected l.Get(2) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = l.RemoveAt(1)
	if v != 7 || !ok {
		t.Fatalf("expected l.RemoveAt(1) to be 7, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	if l.Length() != 2 {
		t.Fatalf("expected l.Length to be 2. Got length: %v", l.Length())
	}

	l.Append(11)
	v, ok = l.RemoveAt(1)
	if v != 9 || !ok {
		t.Fatalf("expected l.RemoveAt(1) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = l.Remove(9)
	if ok {
		t.Fatalf("expected l.Remove(9), ok to be false. Got value: %v, ok: %t", v, ok)
	}
	v, ok = l.RemoveAt(0)
	if v != 5 || !ok {
		t.Fatalf("expected l.RemoveAt(0) to be 5, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = l.RemoveAt(0)
	if v != 11 || !ok {
		t.Fatalf("expected l.RemoveAt(0) to be 11, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	if l.Length() != 0 {
		t.Fatalf("expected l.Length to be 0. Got length: %v", l.Length())
	}

	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(9)

	v, ok = l.Get(2)
	if v != 5 || !ok {
		t.Fatalf("expected l.Get(2) to be 5, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = l.Get(0)
	if v != 5 || !ok {
		t.Fatalf("expected l.Get(0) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	v, ok = l.Remove(9)
	if v != 9 || !ok {
		t.Fatalf("expected l.Remove(9) to be 9, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
	if l.Length() != 2 {
		t.Fatalf("expected l.Length to be 2. Got length: %v", l.Length())
	}
	v, ok = l.Get(0)
	if v != 7 || !ok {
		t.Fatalf("expected l.Get(0) to be 7, and ok to be true. Got value: %v, ok: %t", v, ok)
	}
}
