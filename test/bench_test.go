package test

import (
	"fmt"
	"testing"

	"github.com/kindacommander/sf-encoder/internal/tree"

	"github.com/kindacommander/sf-encoder/internal/counter"
)

func Benchmark(b *testing.B) {
	//str := strings.Join(os.Args[1:], "")
	for i := 0; i < b.N; i++ {
		str := "abbads,mnmxcvkxl;ckvoipoit][p[]p[]p[p]qweqwezxchkhkhghlgfghfghklggnhxcvmn,/./,.,//.,g"
		if len(str) < 1 {
			fmt.Println("Please, enter a string to encode.")
			return
		}
		fmt.Println(counter.FreqCount(str))
		tree := tree.BuildCodeTree()
		(*tree).PrintTree()
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}
