package stack_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/stack"
)

func TestStack(t *testing.T) {
	list := stack.New[int]()

    list.Push(5);
    list.Push(7);
    list.Push(9);

	v, ok := list.Pop()
	if v != 9 || !ok {
		t.Fatalf("expected list.Pop() value to be 9, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	l := list.Length()
	if l != 2 || !ok {
		t.Fatalf("expected list.Length() to be 2, got %v", l)
	}

    list.Push(11);
    v, ok = list.Pop()
	if v != 11 || !ok {
		t.Fatalf("expected list.Pop() value to be 11, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = list.Pop()
	if v != 7 || !ok {
		t.Fatalf("expected list.Pop() value to be 7, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = list.Peek()
	if v != 5 || !ok {
		t.Fatalf("expected list.Pop() value to be 5, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = list.Pop()
	if v != 5 || !ok {
		t.Fatalf("expected list.Pop() value to be 5, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = list.Pop()
	if ok {
		t.Fatalf("expected list.Pop() ok to be false, got value: %v, ok: %t", v, ok)
	}

    list.Push(69);
    v, ok = list.Peek()
	if v != 69 || !ok {
		t.Fatalf("expected list.Pop() value to be 69, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    l = list.Length()
	if l != 1 || !ok {
		t.Fatalf("expected list.Length() to be 1, got %v", l)
	}
}
