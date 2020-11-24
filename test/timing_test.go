package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kindacommander/sf-encoder/internal/tree"
)

func TestTiming(t *testing.T) {
	defer elapsed("Building tree")()

	str := "abbads,mnmxcvkxl;ckvoipoit][p[]p[]p[p]qweqwezxchkhkhghlgfghfghklggnhxcvmn,/./,.,//.,g"
	tree.BuildCodeTree(str)
}

func elapsed(fn string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", fn, time.Since(start))
	}
}
