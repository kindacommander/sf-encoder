package tree

import (
	"sync"
)

type node struct {
	data      string
	leftNode  *node
	rightNode *node
}

func newNode(data string) *node {
	return &node{data, nil, nil}
}

func (n *node) isLeaf() bool {
	return n.leftNode == nil && n.rightNode == nil
}

func (n *node) insert(wg *sync.WaitGroup, data string) {
	defer wg.Done()
	if n == nil {
		return
	}
	if len(data) == 1 {
		return
	}
	halfLen := findHalfLen(data)
	n.leftNode = newNode(data[:halfLen])
	n.rightNode = newNode(data[halfLen:])

	var newWg sync.WaitGroup
	newWg.Add(1)
	go n.leftNode.insert(&newWg, data[:halfLen])
	newWg.Add(1)
	go n.rightNode.insert(&newWg, data[halfLen:])
	newWg.Wait()
}

func findHalfLen(str string) int {
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
		for _, d := range currFreqsBuf {
			if string(r) == d.Str {
				freq += d.TotalFreq
			}
		}
	}
	return freq
}
