package stringconcatenation

import (
	"testing"
)

type Tests struct {
	name     string
	input    int
	expected string
}

var testsValues = []Tests{
	{
		name:     "five iterations",
		input:    5,
		expected: "aaaaa",
	},
	{
		name:     "ten iterations",
		input:    10,
		expected: "aaaaaaaaaa",
	},
	{
		name:     "negative input",
		input:    -3,
		expected: "",
	},
}

func TestStringFromAssignment(t *testing.T) {

	for _, tt := range testsValues {
		t.Run(tt.name, func(t *testing.T) {
			got := StringFromAssignment(tt.input)
			if got != tt.expected {
				t.Errorf("StringFromAssignment(%d) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestStringFromAppendJoin(t *testing.T) {
	for _, tt := range testsValues {
		t.Run(tt.name, func(t *testing.T) {
			got := StringFromAppendJoin(tt.input)
			if got != tt.expected {
				t.Errorf("StringFromAppendJoin(%d) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
func TestStringFromBuffer(t *testing.T) {
	for _, tt := range testsValues {
		t.Run(tt.name, func(t *testing.T) {
			got := StringFromBuffer(tt.input)
			if got != tt.expected {
				t.Errorf("StringFromAppendBuffer(%d) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkStringFromAssignment(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAssignment(10)
	}
}

func BenchmarkStringFromAppendJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAppendJoin(10)
	}
}

func BenchmarkStringFromBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromBuffer(10)
	}
}
