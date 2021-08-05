package anagram

import (
	"fmt"
	"sort"
	"testing"
)

func equal(a, b [][]string) bool {
	if len(b) != len(a) {
		return false
	}

	flag := true
	for i, aElem := range a {
		bElem := b[i]
		sort.Strings(aElem)
		sort.Strings(bElem)
		if len(bElem) <= 0 {
			fmt.Println("Index out of range")
			flag = false
			break
		}

		if fmt.Sprintf("%v", aElem) != fmt.Sprintf("%v", bElem) {
			fmt.Printf("Expected: %s, Actual: %s", aElem, bElem)
			flag = false
			break
		}
	}

	return flag
}

func TestDetectAnagrams(t *testing.T) {
	for _, tt := range testCases {
		actual := Detect(tt.candidates)
		if !equal(tt.expected, actual) {
			msg := `FAIL: %s
				Candidates %q
				Expected %q
				Got %q
				`
			t.Fatalf(msg, tt.description, tt.candidates, tt.expected, actual)
		} else {
			t.Logf("PASS: %s", tt.description)
		}
	}
}

func BenchmarkDetectAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Detect(tt.candidates)
		}
	}
}
