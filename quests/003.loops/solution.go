package loops

import (
	"slices"
	"strings"
)

func SumEvenNumbers(n int) int {
	if n <= 0 {
		return 0
	}
	var sum int
	for i := range n + 1 {
		if i%2 == 0 {
			sum = sum + i
		}
	}
	return sum
}

var (
	vowels = []string{"a", "e", "i", "o", "u"}
)

func KeepOnlyConsonants(strs []string) []string {
	res := []string{}
	for _, strg := range strs {
		var chrs string
		for _, c := range strg {
			if !slices.Contains(vowels, strings.ToLower(string(c))) {
				chrs += string(c)
			}
		}
		if chrs != "" {
			res = append(res, chrs)
		}
	}
	return res
}
