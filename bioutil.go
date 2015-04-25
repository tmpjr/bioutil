package bioutil

func NumberToSymbol(index int) string {
	var symbol string
	symbols := map[int]string{0: "A", 1: "C", 2: "G", 3: "T"}
	if _, ok := symbols[index]; ok {
		symbol = symbols[index]
	} else {
		symbol = ""
	}

	return symbol
}

func SymbolToNumber(symbol string) int {
	bases := map[string]int{"A": 0, "C": 1, "G": 2, "T": 3}
	return bases[symbol]
}

func PatternToNumber(pattern string) int {
	if pattern == "" {
		return 0
	}
	// get last char of pattern
	symbol := pattern[len(pattern)-1:]
	// remove last char of pattern
	p := pattern[0 : len(pattern)-1]
	// compute base4 value of pattern
	value := 4*PatternToNumber(p) + SymbolToNumber(symbol)

	return value
}

func NumberToPattern(index int, k int) string {
	if k == 1 {
		return NumberToSymbol(index)
	}

	q := index / 4
	r := index % 4
	pattern := NumberToPattern(q, k-1)

	symbol := NumberToSymbol(r)

	return pattern + symbol
}

// Accept a string and reverse it rune-wise left to right
func ReversePattern(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
