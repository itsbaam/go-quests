package strings

import "unicode/utf8"

type TextStats struct {
	ByteLength int
	RuneCount  int
}

func AnalyzeText(s string) TextStats {
	return TextStats{
		ByteLength: len(s),
		// RuneCount:  len([]rune(s)),
		RuneCount: utf8.RuneCountInString(s),
	}
}

func RuneFrequencies(s string) map[rune]int {
	rF := make(map[rune]int)

	for _, r := range s {
		rF[r] = rF[r] + 1
	}
	// for len(s) > 0 {
	// 	r, size := utf8.DecodeRuneInString(s)
	// 	rF[r] = rF[r] + 1
	// 	s = s[size:]
	// }

	return rF
}

func FirstRunePosition(s string, target rune) int {
	for i, r := range s {
		if r == target {
			return i
		}
	}
	// for i := 0; i < len(s); {
	// 	r, size := utf8.DecodeRuneInString(s[i:])
	// 	if r == target {
	// 		return i
	// 	}
	// 	i += size
	// }
	return -1
}

func ExtractRunes(s string) []rune {
	// return []rune(s)

	runes := make([]rune, 0, utf8.RuneCountInString(s))
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		runes = append(runes, r)
		s = s[size:]
	}
	return runes
}

func HasOnlyASCII(s string) bool {
	// for _, r := range s {
	// 	if r > 127 {
	// 		return false
	// 	}
	// }
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r > 127 {
			return false
		}
		s = s[size:]
	}
	return true
}
