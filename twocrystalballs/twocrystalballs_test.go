package twocrystalballs_test

import (
	"math/rand"
	"testing"

	"github.com/agent-e11/kata-machine-go/twocrystalballs"
)

func TestTwoCrystalBalls(t *testing.T) {
	expect := rand.Intn(1000)
	data := make([]bool, 1000)
	for i := expect; i < 1000; i++ {
		data[i] = true
	}

	idx := twocrystalballs.TwoCrystalBalls(data)
	if idx != expect {
		t.Fatalf("expected idx index to be %v, got %v", expect, idx)
	}
	idx = twocrystalballs.TwoCrystalBalls(make([]bool, expect))
	if idx != -1 {
		t.Fatalf("expected idx index to be -1, got %v", idx)
	}
}
