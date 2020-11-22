package tree

import (
	"fmt"
	"strings"
	"sync"

	"github.com/kindacommander/sf-encoder/internal/counter"
)

type Data counter.Data

type Node struct {
	data      string
	leftNode  *Node
	rightNode *Node
}

func NewNode(data string) *Node {
	return &Node{data, nil, nil}
}

func (n *Node) isLeaf() bool {
	return n.leftNode == nil && n.rightNode == nil
}

func (n *Node) insert(wg *sync.WaitGroup, data string) {
	defer wg.Done()
	if n == nil {
		return
	}
	if len(data) == 1 {
		return
	}
	halfLen := FindHalfLen(data)
	n.leftNode = NewNode(data[:halfLen])
	n.rightNode = NewNode(data[halfLen:])

	var newWg sync.WaitGroup
	newWg.Add(1)
	go n.leftNode.insert(&newWg, data[:halfLen])
	newWg.Add(1)
	go n.rightNode.insert(&newWg, data[halfLen:])
	newWg.Wait()
}

func FindHalfLen(str string) int {
	halfFreqLen := len(str) / 2
	for {
		firstHalfFreq := evaluateFreq(str[:halfFreqLen])
		secondHalfFreq := evaluateFreq(str[halfFreqLen:])
		if abs(firstHalfFreq-secondHalfFreq) > evaluateFreq(string(str[halfFreqLen-1])) {
			halfFreqLen--
		} else {
			break
		}
	}
	return halfFreqLen
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func evaluateFreq(str string) int {
	freq := 0
	for _, r := range str {
		for _, d := range counter.Freqs {
			if string(r) == d.Str {
				freq += d.TotalFreq
			}
		}
	}
	return freq
}

type Tree struct {
	root *Node
}

// Sholud be done somehow better
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
