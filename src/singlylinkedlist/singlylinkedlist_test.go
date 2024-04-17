package singlylinkedlist_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/singlylinkedlist"
	"github.com/agent-e11/kata-machine-go/listtest"
)

func TestDoublyLinkedList(t *testing.T) {
	l := singlylinkedlist.New[int]()

	listtest.TestList(t, &l)
}
