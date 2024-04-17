package arraylist_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/arraylist"
	"github.com/agent-e11/kata-machine-go/listtest"
)

func TestArrayList(t *testing.T) {
	l := arraylist.New[int](3)
	listtest.TestList(t, &l)
}
