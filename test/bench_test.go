package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kindacommander/sf-encoder/internal/tree"
)

func Benchmark(b *testing.B) {
	var start time.Time
	var timing time.Duration

	for i := 0; i < b.N; i++ {

		start = time.Now()

		str := "abbads,mnmxcvkxl;ckvoipoit][p[]p[]p[p]qweqwezxchkhkhghlgfghfghklggnhxcv192839095654-098-0213-==-==++_+++___~~~````asdkjaskxcn,mbmn,/./,.,//.,g"
		tree.BuildCodeTree(str)
		//t := tree.BuildCodeTree(str)
		// t.PrintTree()
		// t.PrintCodeTable()
		// fmt.Println("Encoded string: ")
		// fmt.Println(t.Encode())

		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}

		timing += time.Since(start)

	}
	fmt.Printf("Average timing: %v\n", timing/time.Duration(b.N))
}
