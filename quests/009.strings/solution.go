package strings

type TextStats struct {
	ByteLength int
	RuneCount  int
}

func AnalyzeText(s string) TextStats {
	return TextStats{
		ByteLength: len(s),
		RuneCount:  len([]rune(s)),
	}
}

func RuneFrequencies(s string) map[rune]int {
	rF := make(map[rune]int)
	for _, r := range []rune(s) {
		rF[r] = rF[r] + 1
	}
	return rF
}

func FirstRunePosition(s string, target rune) int {
	for i, r := range s {
		if r == target {
			return i
		}
	}
	return -1
}

func ExtractRunes(s string) []rune {
	return []rune(s)
}

func HasOnlyASCII(s string) bool {
	for _, r := range []rune(s) {
		if r > 127 {
			return false
		}
	}
	return true
}
