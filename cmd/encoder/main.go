package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kindacommander/sf-encoder/internal/tree"
)

func main() {
	str := strings.Join(os.Args[1:], "")
	if len(str) < 1 {
		fmt.Println("Please, enter a string to encode.")
		return
	}
	t := tree.BuildCodeTree(str)

	fmt.Println("Tree:")
	t.PrintTree()

	fmt.Println("Code table:")
	t.PrintCodeTable()

	fmt.Println("Encoded string:")
	fmt.Println(t.Encode())
}
