package test

import (
	"testing"

	"github.com/kindacommander/sf-encoder/internal/tree"

	"github.com/kindacommander/sf-encoder/internal/counter"
)

func TestHalfLen(t *testing.T) {
	counter.Freqs = []counter.Data{
		counter.Data{"a", 5},
		counter.Data{"d", 4},
		counter.Data{"b", 2},
		counter.Data{"s", 1},
	}
	if hl := tree.FindHalfLen("adbs"); hl != 1 {
		t.Error("Expected 1, got ", hl)
	}
	counter.Freqs = []counter.Data{
		counter.Data{"a", 6},
		counter.Data{"b", 3},
		counter.Data{"c", 2},
		counter.Data{"d", 2},
		counter.Data{"e", 1},
	}
	if hl := tree.FindHalfLen("abcde"); hl != 1 {
		t.Error("Expected 1, got ", hl)
	}
	counter.Freqs = []counter.Data{
		counter.Data{"b", 39},
		counter.Data{"c", 18},
	}
	if hl := tree.FindHalfLen("bc"); hl != 1 {
		t.Error("Expected 1, got ", hl)
	}
}
