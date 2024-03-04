package betonicSort

import (
	"errors"
)

// isPowerOfTwo returns true if n is a power of 2, false otherwise.
func isPowerOfTwo(n int) bool {
	return (n != 0) && (n&(n-1) == 0)
}

// BetonicSort sorts the string in ascending order. Returns sorted array and error.
func BetonicSort(arr string) (string, error) {
	n := len(arr) // length of array

	// check if array is empty
	if n == 0 {
		return "", errors.New("empty array")
	}

	// check if n is a power of 2
	if !isPowerOfTwo(n) {
		return "", errors.New("array length must be a power of 2")
	}

	// convert string to array of runes
	arrRune := []rune(arr)

	// betonic sort
	for k := 2; k <= n; k *= 2 {
		for j := k / 2; j > 0; j /= 2 {
			for i := 0; i < len(arrRune); i++ {
				l := i ^ j
				if l > i && l < len(arrRune) {
					if (i&k == 0) && (arrRune[i] > arrRune[l]) || (i&k != 0) && (arrRune[i] < arrRune[l]) {
						arrRune[i], arrRune[l] = arrRune[l], arrRune[i]
					}
				}
			}
		}
	}

	// convert array of runes to string
	return string(arrRune), nil
}
