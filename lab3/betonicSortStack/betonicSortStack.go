package betonicSortStack

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/17HIERARCH70/stackGo"
)

// Check if string is an integer
func isInteger(arr string) bool {
	match, _ := regexp.MatchString(`^\d+$`, arr)
	return match
}

// Check if number is a power of 2
func isPowerOfTwo(n int) bool {
	return (n != 0) && (n&(n-1) == 0)
}

// Betonic sort for integers
func BitonicSort(arr string) (string, error) {
	n := len(arr)

	if n == 0 {
		return "", errors.New("empty array")
	}

	if !isPowerOfTwo(n) {
		return "", errors.New("array length must be a power of 2")
	}

	if !isInteger(arr) {
		return "", errors.New("array must include only numbers")
	}

	// Create a stack struct
	stack := stackGo.NewStack()

	// Push array to stack
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(string(arr[i]))
		if err != nil {
			return "", err
		}
		stack.Push(num)
	}

	// Betonic sort for stack
	for k := 2; k <= n; k *= 2 {
		for j := k / 2; j > 0; j /= 2 {
			for i := 0; i < stack.Len(); i++ {
				l := i ^ j
				if l > i && l < stack.Len() {
					result1, err := stack.Compare(i, l)
					if err != nil {
						return "", err
					}
					result2, err := stack.Compare(i, l)
					if err != nil {
						return "", err
					}
					if (i&k == 0) && (result1 == 1) || (i&k != 0) && (result2 == -1) {
						err := stack.Swap(i, l)
						if err != nil {
							return "", errors.New("error with swapping")
						}
					}
				}
			}
		}
	}
	res := ""

	// Pop elements from the stack to construct the result
	for stack.Len() > 0 {
		num, _ := stack.Pop()
		res = strconv.Itoa(num) + res
	}

	// Return the sorted result as a string
	return res, nil
}
