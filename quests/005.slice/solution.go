package slice

func ProcessScores(scores []int) []int {
	if len(scores) == 0 {
		return []int{}
	}
	var scoresRes []int
	for _, score := range scores {
		if score < 0 || score > 100 {
			continue
		}
		if score < 40 {
			score = 40
		}
		scoresRes = append(scoresRes, score)
	}

	if len(scoresRes) > 5 {
		for i := 0; i < len(scoresRes); i++ {
			scoresRes[i] += 5
			if scoresRes[i] > 100 {
				scoresRes[i] = 100
			}
		}
	}
	return scoresRes
}
