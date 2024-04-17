package ringbuffer_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/ringbuffer"
)

func TestRingBuffer(t *testing.T) {
	buffer := ringbuffer.New[int]()

	v, ok := buffer.Pop()
	if ok {
		t.Fatalf("expected buffer.Pop() ok to be false, got value: %v, ok: %t", v, ok)
	}
	l := buffer.Length()
	if l != 0 {
		t.Fatalf("expected buffer.Length() to be 0, got %v", l)
	}

    buffer.Push(5);
	v, ok = buffer.Pop()
	if v != 5 || !ok {
		t.Fatalf("expected buffer.Pop() value to be 5, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = buffer.Pop()
	if ok {
		t.Fatalf("expected buffer.Pop() ok to be false, got value: %v, ok: %t", v, ok)
	}

    buffer.Push(42);
    buffer.Push(9);
    v, ok = buffer.Pop()
	if v != 42 || !ok {
		t.Fatalf("expected buffer.Pop() value to be 42, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	l = buffer.Length()
	if l != 2 {
		t.Fatalf("expected buffer.Length() to be 2, got %v", l)
	}
    v, ok = buffer.Pop()
	if v != 9 || !ok {
		t.Fatalf("expected buffer.Pop() value to be 9, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = buffer.Pop()
	if ok {
		t.Fatalf("expected buffer.Pop() ok to be false, got value: %v, ok: %t", v, ok)
	}

    buffer.Push(42);
    buffer.Push(9);
    buffer.Push(12);
	l = buffer.Length()
	if l != 3 {
		t.Fatalf("expected buffer.Length() to be 3, got %v", l)
	}
    v, ok = buffer.Get(2)
	if v != 12 || !ok {
		t.Fatalf("expected buffer.Get() value to be 12, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = buffer.Get(1)
	if v != 9 || !ok {
		t.Fatalf("expected buffer.Get() value to be 9, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = buffer.Get(0)
	if v != 42 || !ok {
		t.Fatalf("expected buffer.Get() value to be 42, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	l = buffer.Length()
	if l != 3 {
		t.Fatalf("expected buffer.Length() to be 3, got %v", l)
	}
}
