package tree

import (
	"strings"
)

func BuildCodeTreeUnparallel(str string) Tree {
	freqs := freqCount(str)
	var decSlice []string
	for _, data := range freqs {
		decSlice = append(decSlice, data.Str)
	}

	codes = make(chan encodedChar)
	done = make(chan struct{})

	codeTable := make([]encodedChar, len(currFreqsBuf))

	go codeListener(codes, &codeTable)

	tree := Tree{str, newNode(strings.Join(decSlice, "")), freqs, codeTable}

	tree.root.insertUnparallel(strings.Join(decSlice, ""), "", codes)
	close(codes)
	<-done

	return tree
}

func (n *node) insertUnparallel(data string, currCode string, codes chan encodedChar) {
	if n == nil {
		return
	}
	if len(data) == 1 {
		codes <- encodedChar{data, currCode}
		return
	}
	halfLen := findHalfLen(data)
	n.leftNode = newNode(data[:halfLen])
	n.rightNode = newNode(data[halfLen:])

	leftStr := currCode + "1"
	rightStr := currCode + "0"
	n.leftNode.insertUnparallel(data[:halfLen], leftStr, codes)
	n.rightNode.insertUnparallel(data[halfLen:], rightStr, codes)
}
