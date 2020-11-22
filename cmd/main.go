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
	if len(str) < 1 {
		fmt.Println("Please, enter a string to encode.")
		return
	}
	fmt.Println(counter.FreqCount(str))
	tree := tree.BuildCodeTree()
	//time.Sleep(time.Second * 3)
	(*tree).PrintTree()
}
