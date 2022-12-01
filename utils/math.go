package utils

import (
	"sort"
	"strconv"
)

// ConvertToInt converts a string to integer
func ConvertToInt(s string) int {
	value, err := strconv.Atoi(s)
	Check(err)
	return value
}

// SumSlice sums the contents of the slice
func SumSlice(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func MulSlice(numbers []int) int {
	result := 1
	for _, n := range numbers {
		result *= n
	}
	return result
}

// Return the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Median(numbers []int) int {
	sort.Ints(numbers)
	if len(numbers)%2 == 0 {
		return (numbers[len(numbers)/2] + (numbers[len(numbers)/2] + 1)) / 2
	} else {
		return numbers[(len(numbers)+1)/2]
	}
}

func Mean(numbers []int) int {
	sum := SumSlice(numbers)
	return (sum + 1) / len(numbers)
}
