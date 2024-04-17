package doublylinkedlist_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/doublylinkedlist"
	"github.com/agent-e11/kata-machine-go/listtest"
)

func TestDoublyLinkedList(t *testing.T) {
	l := doublylinkedlist.New[int]()

	listtest.TestList(t, &l)
}
