package bioutil

import "testing"

func TestNumberToSymbol(t *testing.T) {
	symbols := map[int]string{0: "A", 1: "C", 2: "G", 3: "T"}
	for index, symbol := range symbols {
		v := NumberToSymbol(index)
		if v != symbol {
			t.Errorf("Expected %s, got %v", symbol, v)
		}
	}
}

func TestSymbolToNumber(t *testing.T) {
	bases := map[string]int{"A": 0, "C": 1, "G": 2, "T": 3}
	for symbol, value := range bases {
		v := SymbolToNumber(symbol)
		if v != value {
			t.Errorf("Expected %d, got %v", value, v)
		}
	}
}

func TestPatternToNumber(t *testing.T) {
	pattern := "CGTAAAACTGCTTGAGA"
	value := 7248256904
	v := PatternToNumber(pattern)
	if v != value {
		t.Errorf("Expected %d, got %v", value, v)
	}
}

func TestNumberToPattern(t *testing.T) {
	pattern := "ACGGAGGT"
	k := 8
	index := 6699
	v := NumberToPattern(index, k)
	if v != pattern {
		t.Errorf("Expected %s, got %v", pattern, v)
	}
}

func TestReversePattern(t *testing.T) {
	pattern := "ACGGAGGT"
	reverse := "TGGAGGCA"
	v := ReversePattern(pattern)
	if v != reverse {
		t.Errorf("Expected %s, got %v", reverse, v)
	}
}
