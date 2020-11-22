package tree

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

var currFreqsBuf []data // Current freqs buffer

type Tree struct {
	root     *node
	FreqsBuf []data // Individual buffer for each tree
}

func (t Tree) PrintTree() {
	printNode(t.root, 0, 'M')
}

func printNode(node *node, ns int, ch rune) {
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

func BuildCodeTree(str string) Tree {
	freqs := freqCount(str)
	var decSlice []string
	for _, data := range freqs {
		decSlice = append(decSlice, data.Str)
	}
	tree := Tree{newNode(strings.Join(decSlice, "")), freqs}

	var wg sync.WaitGroup
	wg.Add(1)
	tree.root.insert(&wg, strings.Join(decSlice, ""))
	wg.Wait()
	return tree
}

type data struct {
	Str       string
	TotalFreq int
}

func freqCount(str string) []data {
	for _, el := range str {
		if i, ok := exists(el); ok {
			currFreqsBuf[i].TotalFreq++
		} else {
			currFreqsBuf = append(currFreqsBuf, data{string(el), 1})
		}
	}
	sort.SliceStable(currFreqsBuf, func(i, j int) bool {
		return currFreqsBuf[i].TotalFreq > currFreqsBuf[j].TotalFreq
	})
	freqs := make([]data, len(currFreqsBuf))
	copy(freqs, currFreqsBuf)

	//  ONLY FOR DEBUG
	fmt.Println(freqs)
	//  ONLY FOR DEBUG

	return freqs
}

func exists(str rune) (int, bool) {
	for i, el := range currFreqsBuf {
		if el.Str == string(str) {
			return i, true
		}
	}
	return -1, false
}
