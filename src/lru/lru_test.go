package lru_test

import (
	"testing"

	// TODO: Change this alias
	"github.com/agent-e11/kata-machine-go/src/lru"
)

func TestLRU(t *testing.T) {
	cache := lru.New[string, int](3)

	v, ok := cache.Get("foo")
	if ok {
		t.Fatalf("expected foo to not be present, got value: %v, ok: %t", v, ok)
	}
	cache.Update("foo", 10)
	v, ok = cache.Get("foo")
	if v != 10 || !ok {
		t.Fatalf("expected foo to be 10, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	cache.Update("bar", 20)
	v, ok = cache.Get("bar")
	if v != 20 || !ok {
		t.Fatalf("expected bar to be 20, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	cache.Update("baz", 30)
	v, ok = cache.Get("baz")
	if v != 30 || !ok {
		t.Fatalf("expected baz to be 30, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	cache.Update("ball", 100)
	v, ok = cache.Get("ball")
	if v != 100 || !ok {
		t.Fatalf("expected ball to be 100, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	v, ok = cache.Get("foo")
	if ok {
		t.Fatalf("expected foo to not be present, got value: %v, ok: %t", v, ok)
	}
	v, ok = cache.Get("bar")
	if v != 20 || !ok {
		t.Fatalf("expected ball to be 100, and ok to be true, got value: %v, ok: %t", v, ok)
	}

	cache.Update("foo", 10)
	v, ok = cache.Get("foo")
	if v != 10 || !ok {
		t.Fatalf("expected foo to be 10, and ok to be true, got value: %v, ok: %t", v, ok)
	}
}
