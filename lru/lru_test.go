package lru_test

import (
	"testing"

	lrucache "github.com/agent-e11/kata-machine-go/lru"
)

func TestLRU(t *testing.T) {
	lru := lrucache.New[string, int](3)

	v, ok := lru.Get("foo")
	if ok {
		t.Fatalf("expected foo to not be present, got value: %v, ok: %t", v, ok)
	}
	lru.Update("foo", 10)
	v, ok = lru.Get("foo")
	if v != 10 || !ok {
		t.Fatalf("expected foo to be 10, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	lru.Update("bar", 20)
	v, ok = lru.Get("bar")
	if v != 20 || !ok {
		t.Fatalf("expected bar to be 20, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	lru.Update("baz", 30)
	v, ok = lru.Get("baz")
	if v != 30 || !ok {
		t.Fatalf("expected baz to be 30, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	lru.Update("ball", 100)
	v, ok = lru.Get("ball")
	if v != 100 || !ok {
		t.Fatalf("expected ball to be 100, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	v, ok = lru.Get("foo")
	if ok {
		t.Fatalf("expected foo to not be present, got value: %v, ok: %t", v, ok)
	}
	v, ok = lru.Get("bar")
	if v != 20 || !ok {
		t.Fatalf("expected ball to be 100, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	lru.Update("foo", 10)
	v, ok = lru.Get("foo")
	if v != 10 || !ok {
		t.Fatalf("expected foo to be 10, and ok to be true, got value: %v, ok: %t", v, ok)
	}
}
