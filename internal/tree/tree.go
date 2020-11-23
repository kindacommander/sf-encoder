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

	CodeTable = make([]encodedChar, len(currFreqsBuf))
	codes = make(chan encodedChar)
	done = make(chan struct{})

	go codeListener(codes)

	tree := Tree{newNode(strings.Join(decSlice, "")), freqs}

	var wg sync.WaitGroup
	wg.Add(1)
	tree.root.insert(&wg, strings.Join(decSlice, ""), "", codes)
	wg.Wait()
	close(codes)
	<-done
	return tree
}

var (
	mu        sync.Mutex
	CodeTable []encodedChar
	codes     chan encodedChar
	done      chan struct{}
)

func codeListener(codes chan encodedChar) {
	i := 0
	var wg sync.WaitGroup
	for code := range codes {
		wg.Add(1)
		go func(i int, code encodedChar, wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			CodeTable[i] = code
			mu.Unlock()
		}(i, code, &wg)
		i++
	}
	wg.Wait()
	done <- struct{}{}
	// for {
	// 	select {
	// 	case code := <-codes:
	// 		mu.Lock()
	// 		CodeTable[i] = code
	// 		i++
	// 		mu.Unlock()
	// 	}
	// 	if i == len(CodeTable) {
	// 		close(codes)
	// 		break
	// 	}
	// }
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
