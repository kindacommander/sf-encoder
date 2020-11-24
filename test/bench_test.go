package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kindacommander/sf-encoder/internal/tree"
)

func Benchmark(b *testing.B) {
	start := time.Now()
	for i := 0; i < b.N; i++ {
		str := "abbads,mnmxcvkxl;ckvoipoit][p[]p[]p[p]qweqwezxchkhkhghlgfghfghklggnhxcvmn,/./,.,//.,g"
		tree.BuildCodeTree(str)
		//t := tree.BuildCodeTree(str)
		// t.PrintTree()
		// t.PrintCodeTable()
		// fmt.Println("Encoded string: ")
		// fmt.Println(t.Encode())

		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
		fmt.Printf("Average timing: %v\n", time.Since(start)/time.Duration(b.N))
	}
}
