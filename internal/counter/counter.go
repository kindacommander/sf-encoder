package counter

import "sort"

type Data struct {
	Str       string
	TotalFreq int
}

var Freqs []Data

func FreqCount(str string) []Data {
	for _, el := range str {
		if i, ok := exists(el); ok {
			Freqs[i].TotalFreq++
		} else {
			Freqs = append(Freqs, Data{string(el), 1})
		}
	}
	sort.SliceStable(Freqs, func(i, j int) bool {
		return Freqs[i].TotalFreq > Freqs[j].TotalFreq
	})

	return Freqs
}

func exists(Str rune) (int, bool) {
	for i, el := range Freqs {
		if el.Str == string(Str) {
			return i, true
		}
	}
	return -1, false
}
