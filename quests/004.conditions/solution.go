package conditions

func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string {
	if age <= 0 || balance < 0 {
		return "INVALID"
	} else if age < 18 || !hasID {
		return "REJECTED"
	} else if isVIP && balance >= 10000 {
		return "VIP_ACCESS"
	} else if balance >= 1000 {
		return "STANDARD_ACCESS"
	}

	return "LIMITED_ACCESS"
}

func EvaluateGrade(score int) string {
	var grade string
	switch {
	case (score < 0 || score > 100):
		grade = "INVALID"
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	case score < 60:
		grade = "F"
	}

	return grade
}
