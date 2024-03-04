package betonicSort_test

import (
	"errors"
	"testing"

	"github.com/17HIERARCH70/betonicSort"
)

func TestBetonicSort(t *testing.T) {
	tests := []struct {
		input string
		want  string
		err   error
	}{
		{input: "1234", want: "1234", err: nil},
		{input: "4321", want: "1234", err: nil},
		{input: "1243", want: "1234", err: nil},
		{input: "3214", want: "1234", err: nil},
		{input: "12345678", want: "12345678", err: nil},
		{input: "87654321", want: "12345678", err: nil},
		{input: "12435678", want: "12345678", err: nil},
		{input: "87543212", want: "12234578", err: nil},
		{input: "1232332456789", want: "", err: errors.New("array length must be a power of 2")},
		{input: "", want: "", err: errors.New("empty array")},
	}
	for _, test := range tests {
		got, err := betonicSort.BetonicSort(test.input)
		if (err == nil && test.err != nil) || (err != nil && test.err == nil) || (err != nil && test.err != nil && err.Error() != test.err.Error()) {
			t.Errorf("betonicSort(%q) = %q, %v; want %q, %v", test.input, got, err, test.want, test.err)
		}
		if got != test.want {
			t.Errorf("betonicSort(%q) = %q, want %q, error: %v", test.input, got, test.want, test.err)
		}
	}
}
