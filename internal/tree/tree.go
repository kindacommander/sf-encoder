package tree

import (
	"fmt"
	"strings"
	"sync"

	"github.com/kindacommander/sf-encoder/internal/counter"
)

type Tree struct {
	root *Node
}

func (t Tree) PrintTree() {
	printNode(t.root, 0, 'R')
}

func printNode(node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("%c:%v\n", ch, node.data)
	printNode(node.leftNode, ns+2, 'L')
	printNode(node.rightNode, ns+2, 'R')
}

func BuildCodeTree() Tree {
	var decSlice []string
	for _, data := range counter.Freqs {
		decSlice = append(decSlice, data.Str)
	}
	tree := Tree{NewNode(strings.Join(decSlice, ""))}

	var wg sync.WaitGroup
	wg.Add(1)
	tree.root.insert(&wg, strings.Join(decSlice, ""))
	wg.Wait()
	return tree
}
