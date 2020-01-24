package commons

import (
	"testing"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "", want: false},
		{input: "he77", want: false},
		{input: "hello", want: false},
		{input: "5667\n", want: false},
		{input: "1234", want: true},
	}

	for _, tc := range tests {
		got := IsNumeric(tc.input)
		if got != tc.want {
			t.Fatalf("%q expected: %v, got: %v", tc.input, tc.want, got)
		}
	}
}

func TestLeftPad(t *testing.T) {
	tests := []struct {
		input  string
		str    string
		length int
		want   string
	}{
		{input: "", str: "0", length: 3, want: "000"},
		{input: "45", str: "0", length: 9, want: "000000045"},
		{input: "765s", str: "0", length: 0, want: "765s"},
	}

	for _, tc := range tests {
		got := LeftPad(tc.input, tc.str, tc.length)
		if got != tc.want {
			t.Fatalf("%q expected: %q, got: %v", tc.input, tc.want, got)
		}
	}
}
