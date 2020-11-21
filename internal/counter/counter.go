package counter

import "sort"

type Char struct {
	symbol string
	freq   int
}

func FreqCount(str string) []Char {
	var freqs []Char
	for _, el := range str {
		if i, ok := exists(freqs, el); ok {
			freqs[i].freq++
		} else {
			freqs = append(freqs, Char{string(el), 1})
		}
	}
	sort.SliceStable(freqs, func(i, j int) bool {
		return freqs[i].freq > freqs[j].freq
	})

	return freqs
}

func exists(freqs []Char, symbol rune) (int, bool) {
	for i, el := range freqs {
		if el.symbol == string(symbol) {
			return i, true
		}
	}
	return -1, false
}
