package betonicSortStack_test

import (
	"algosITMO/lab3/betonicSortStack"
	"errors"
	"strconv"
	"testing"
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
		{input: "12a4", want: "", err: errors.New("array must include only numbers")},
		{input: "1232332456789", want: "", err: errors.New("array length must be a power of 2")},
		{input: "", want: "", err: errors.New("empty array")},
	}
	for _, test := range tests {
		got, err := betonicSortStack.BitonicSort(test.input)
		if (err == nil && test.err != nil) || (err != nil && test.err == nil) || (err != nil && test.err != nil && err.Error() != test.err.Error()) {
			t.Errorf("betonicSort(%q) = %q, %v; want %q, %v", test.input, got, err, test.want, test.err)
		}
		if got != test.want {
			t.Errorf("betonicSort(%q) = %q, want %q, error: %v", test.input, got, test.want, test.err)
		}
	}
}

// BenchmarkBitonicSort benchmarks the BitonicSort function
func BenchmarkBitonicSort(b *testing.B) {
	// Create a test case with a large input array
	input := "12345678901234567890"

	// Run the BitonicSort function b.N times
	for i := 0; i < b.N; i++ {
		_, _ = betonicSortStack.BitonicSort(input)
	}
}

// Helper function to generate a random string of digits
func generateRandomString(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += strconv.Itoa(i % 10) // Appending digits 0-9 cyclically
	}
	return str
}

// BenchmarkBitonicSortLargeInput benchmarks the BitonicSort function with a large input
func BenchmarkBitonicSortLargeInput(b *testing.B) {
	// Generate a large input array
	input := generateRandomString(1048576) // Change the length as needed

	// Run the BitonicSort function b.N times
	for i := 0; i < b.N; i++ {
		_, _ = betonicSortStack.BitonicSort(input)
	}
}
