package letter

func FrequencyParallel(text string, c chan map[rune]int) {
	c <- Frequency(text)
}

func Frequency(str string) map[rune]int {
	m := make(map[rune]int)
	for _, symbol := range str {
		m[symbol]++
	}
	return m
}

func ConcurrentFrequency(text []string) map[rune]int {
	c := make(chan map[rune]int)
	new_m := make(map[rune]int)
	for i := range text {
		go FrequencyParallel(text[i], c)
	}
	for range text {
		for symbol, count := range <-c {
			new_m[symbol] += count
		}
	}
	return new_m
}
