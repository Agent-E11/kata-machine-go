package hashmap_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/hashmap"
)

func TestHashMap(t *testing.T) {
	m := hashmap.New[string, int]()

	s := m.Size()
	if s != 0 {
		t.Fatalf("expected m.Size() to be 0, got %v", s)
	}
    m.Set("foo", 55);
    s = m.Size()
	if s != 1 {
		t.Fatalf("expected m.Size() to be 1, got %v", s)
	}
    m.Set("fool", 75);
    s = m.Size()
	if s != 2 {
		t.Fatalf("expected m.Size() to be 2, got %v", s)
	}
    m.Set("foolish", 105);
    s = m.Size()
	if s != 3 {
		t.Fatalf("expected m.Size() to be 3, got %v", s)
	}
    m.Set("bar", 69);
    s = m.Size()
	if s != 4 {
		t.Fatalf("expected m.Size() to be 4, got %v", s)
	}

	v, ok := m.Get("bar")
	if v != 69 || !ok {
		t.Fatalf("expected m.Get() value to be 69 and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = m.Get("blaz")
	if ok {
		t.Fatalf("expected m.Get() ok to be false, got value: %v, ok: %t", v, ok)
	}

	v, ok = m.Delete("barblabr")
	if ok {
		t.Fatalf("expected m.Delete() ok to be false, got value: %v, ok: %t", v, ok)
	}
    s = m.Size()
	if s != 4 {
		t.Fatalf("expected m.Size() to be 4, got %v", s)
	}

    v, ok = m.Delete("bar")
	if v != 69 || !ok {
		t.Fatalf("expected m.Get() value to be 69 and ok to be true, got value: %v, ok: %t", v, ok)
	}
    s = m.Size()
	if s != 3 {
		t.Fatalf("expected m.Size() to be 3, got %v", s)
	}
    v, ok = m.Get("bar")
	if ok {
		t.Fatalf("expected m.Get() ok to be false, got value: %v, ok: %t", v, ok)
	}
}
