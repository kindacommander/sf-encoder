package test

import (
	"fmt"
	"testing"

	"github.com/kindacommander/sf-encoder/internal/tree"
)

func Benchmark(b *testing.B) {
	//str := strings.Join(os.Args[1:], "")
	for i := 0; i < b.N; i++ {
		str := "abbads,mnmxcvkxl;ckvoipoit][p[]p[]p[p]qweqwezxchkhkhghlgfghfghklggnhxcvmn,/./,.,//.,g"
		if len(str) < 1 {
			fmt.Println("Please, enter a string to encode.")
			return
		}
		tree.BuildCodeTree(str)
		// tree := tree.BuildCodeTree(str)
		// tree.PrintTree()
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}
