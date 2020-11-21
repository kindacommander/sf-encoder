package counter

func FreqCount(str string) map[string]int {
	freqs := make(map[string]int)
	for _, el := range str {
		if _, exists := freqs[string(el)]; exists == false {
			freqs[string(el)] = 1
		} else {
			freqs[string(el)]++
		}
	}
	return freqs
}
