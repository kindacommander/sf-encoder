package tree

import (
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
