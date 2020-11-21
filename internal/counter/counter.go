package counter

import "sort"

type Char struct {
	Symbol string
	Freq   int
}

var Freqs []Char

func FreqCount(str string) []Char {
	for _, el := range str {
		if i, ok := exists(el); ok {
			Freqs[i].Freq++
		} else {
			Freqs = append(Freqs, Char{string(el), 1})
		}
	}
	sort.SliceStable(Freqs, func(i, j int) bool {
		return Freqs[i].Freq > Freqs[j].Freq
	})

	return Freqs
}

func exists(Symbol rune) (int, bool) {
	for i, el := range Freqs {
		if el.Symbol == string(Symbol) {
			return i, true
		}
	}
	return -1, false
}
