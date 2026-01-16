package functions

import "errors"

// Divide returns a / b or an error if b == 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("invalid division by 0")
	}
	return a / b, nil
}

// SumAll returns the sum of all provided integers
func SumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// MaxMin returns the max and min of all provided integers
// Returns an error if no numbers are provided
func MaxMin(nums ...int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("no nums provided")
	}
	max, min := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min, nil
}

// ConcatAll joins all strings using the provided separator
func ConcatAll(sep string, strs ...string) string {
	res := ""
	for i, strg := range strs {
		res += strg
		if i < (len(strs) - 1) {
			res += sep
		}
	}
	return res
}
