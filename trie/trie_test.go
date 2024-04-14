package trie_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/trie"
)

func TestTrie(t *testing.T) {
	tr := trie.New()

    tr.Insert("foo");
    tr.Insert("fool");
    tr.Insert("foolish");
    tr.Insert("bar");

	words := tr.Find("fo")
	slices.Sort(words)

	expect := []string{
        "foo",
        "fool",
        "foolish",
	}

	if !slices.Equal(words, expect) {
		t.Fatalf("expected words to be %v, got %v", expect, words)
	}

    tr.Delete("fool");

	words = tr.Find("fo")
	slices.Sort(words)


	expect = []string{
        "foo",
        "foolish",
	}

	if !slices.Equal(words, expect) {
		t.Fatalf("expected words to be %v, got %v", expect, words)
	}
}
