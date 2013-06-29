package editd

func EditDistance(s1, s2 string) int {
	return editd(s1, s2)
}

func editd(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2) * insCost(' ')
	}
	if len(s2) == 0 {
		return len(s1) * delCost(' ')
	}

	sub := editd(trim(s1), trim(s2)) + subCost(last(s1), last(s2))
	ins := editd(s1, trim(s2)) + insCost(last(s2))
	del := editd(trim(s1), s2) + delCost(last(s1))

	return min(sub, min(ins, del))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trim(s string) string {
	return s[:len(s)-1]
}

func last(s string) rune {
	return rune(s[len(s)-1])
}

func insCost(r rune) int {
	return 1
}

func delCost(r rune) int {
	return 1
}

func subCost(r1, r2 rune) int {
	if r1 == r2 {
		return 0
	}

	return 1
}
