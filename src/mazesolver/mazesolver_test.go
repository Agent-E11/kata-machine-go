package mazesolver_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go"
	"github.com/agent-e11/kata-machine-go/src/mazesolver"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
        "xxxxxxxxxx x",
        "x        x x",
        "x        x x",
        "x xxxxxxxx x",
        "x          x",
        "x xxxxxxxxxx",
	}

	expect := []kata.Point{
        { X: 1, Y: 5 },
        { X: 1, Y: 4 },
        { X: 2, Y: 4 },
        { X: 3, Y: 4 },
        { X: 4, Y: 4 },
        { X: 5, Y: 4 },
        { X: 6, Y: 4 },
        { X: 7, Y: 4 },
        { X: 8, Y: 4 },
        { X: 9, Y: 4 },
        { X: 10, Y: 4 },
        { X: 10, Y: 3 },
        { X: 10, Y: 2 },
        { X: 10, Y: 1 },
        { X: 10, Y: 0 },
	}

    // there is only one path through
	path := mazesolver.Solve(
		maze,
		'x',
		kata.Point{ X: 1, Y: 5 },
		kata.Point{ X: 10, Y: 0 },
	)

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}
