package queue_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/queue"
)

func TestQueue(t *testing.T) {
	q := queue.New[int]()

    q.Enqueue(5);
    q.Enqueue(7);
    q.Enqueue(9);

	v, ok := q.Dequeue()
	if v != 5 || !ok {
		t.Fatalf("expected q.Dequeue() value to be 5, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	l := q.Length()
	if l != 2 {
		t.Fatalf("expected q.Length() to be 2, got %v", l)
	}

    q.Enqueue(11);
    v, ok = q.Dequeue()
	if v != 7 || !ok {
		t.Fatalf("expected q.Dequeue() value to be 7, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = q.Dequeue()
	if v != 9 || !ok {
		t.Fatalf("expected q.Dequeue() value to be 9, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = q.Peek()
	if v != 11 || !ok {
		t.Fatalf("expected q.Peek() value to be 11, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = q.Dequeue()
	if v != 11 || !ok {
		t.Fatalf("expected q.Dequeue() value to be 11, and ok to be true, got value: %v, ok: %t", v, ok)
	}
    v, ok = q.Dequeue()
	if ok {
		t.Fatalf("expected q.Dequeue() ok to be false, got value: %v, ok: %t", v, ok)
	}
	l = q.Length()
	if l != 0 {
		t.Fatalf("expected q.Length() to be 0, got %v", l)
	}

    q.Enqueue(69);
    v, ok = q.Peek()
	if v != 69 || !ok {
		t.Fatalf("expected q.Peek() value to be 69, and ok to be true, got value: %v, ok: %t", v, ok)
	}
	l = q.Length()
	if l != 1 || !ok {
		t.Fatalf("expected q.Length() to be 1, got %v", l)
	}
}
