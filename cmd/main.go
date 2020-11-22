package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kindacommander/sf-encoder/internal/counter"
	"github.com/kindacommander/sf-encoder/internal/tree"
)

func main() {
	str := strings.Join(os.Args[1:], "")
	fmt.Println(counter.FreqCount(str))
	tree := tree.BuildCodeTree()
	(*tree).PrintTree()
}
