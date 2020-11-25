package tree

import (
	"sort"
)

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
