package tree

import (
	"fmt"
	"strings"
	"sync"
)

var currFreqsBuf []data // Current freqs buffer

type Tree struct {
	root      *node
	FreqsBuf  []data // Individual buffer for each tree
	codeTable []encodedChar
}

func (t Tree) PrintTree() {
	printNode(t.root, 0, 'M')
}

func (t Tree) CodeTable() []encodedChar {
	return t.codeTable
}

func (t Tree) PrintCodeTable() {
	for _, el := range t.codeTable {
		fmt.Println("{" + el.char + "}" + " = " + el.code)
	}
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

var (
	mu    sync.Mutex
	codes chan encodedChar
	done  chan struct{}
)

func BuildCodeTree(str string) Tree {
	freqs := freqCount(str)
	var decSlice []string
	for _, data := range freqs {
		decSlice = append(decSlice, data.Str)
	}

	codes = make(chan encodedChar)
	done = make(chan struct{})

	codeTable := make([]encodedChar, len(currFreqsBuf))

	go codeListener(codes, &codeTable)

	tree := Tree{newNode(strings.Join(decSlice, "")), freqs, codeTable}

	var wg sync.WaitGroup
	wg.Add(1)
	tree.root.insert(&wg, strings.Join(decSlice, ""), "", codes)
	wg.Wait()
	close(codes)
	<-done

	return tree
}

func codeListener(codes chan encodedChar, codeTable *[]encodedChar) {
	i := 0
	var wg sync.WaitGroup
	for code := range codes {
		wg.Add(1)
		go func(i int, code encodedChar, wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			(*codeTable)[i] = code
			mu.Unlock()
		}(i, code, &wg)
		i++
	}
	wg.Wait()
	done <- struct{}{}
}
