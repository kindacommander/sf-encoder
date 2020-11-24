package test

import (
	"fmt"
	"testing"

	"github.com/kindacommander/sf-encoder/internal/tree"
)

func TestDataRace(t *testing.T) {
	str := "abbads,mnmxcvkxl;ckvoipoit][p[]p[]p[p]qweqwezxchkhkhghlgfghfghklggnhxcvmn,/./,.,//.,g"
	if len(str) < 1 {
		fmt.Println("Please, enter a string to encode.")
		return
	}
	tr := tree.BuildCodeTree(str)
	tr.PrintTree()
	tr.PrintCodeTable()
}
