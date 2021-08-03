package main

import (
	"testing"
)

func TestFindFirstStringInBracket(t *testing.T) {
	for _, tt := range testCases {
		actual := findFirstStringInBracket(tt.input)
		if tt.expected != actual {
			msg := `FAIL: %s
				Input %q
				Expected %q
				Got %q
				`
			t.Fatalf(msg, tt.description, tt.input, tt.expected, actual)
		} else {
			t.Logf("PASS: %s", tt.description)
		}
	}
}

func BenchmarkFindFirstStringInBracket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			findFirstStringInBracket(tt.input)
		}
	}
}
