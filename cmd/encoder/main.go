package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kindacommander/sf-encoder/internal/tree"
)

func main() {

	printTree := flag.Bool("t", false, "If true - prints the view of the code tree.")
	printCodeTable := flag.Bool("c", false, "If true - prints the code table.")

	flag.Parse()

	str := strings.Join(os.Args[1:], "")
	if len(str) < 1 {
		fmt.Println("Please, enter a string to encode.")
		return
	}
	t := tree.BuildCodeTree(str)

	if *printTree {
		fmt.Println("Code tree:")
		t.PrintTree()
	}

	if *printCodeTable {
		fmt.Println("Code table:")
		t.PrintCodeTable()
	}

	fmt.Println("Encoded string:")
	fmt.Println(t.Encode())
}
