package static

func NormalSpeed(increment float64) float64 {
	return increment * 4
}

func LetterCount(offset float64, scale float64) int {
	return int((ResolutionWidth - offset) / ((IconWidth - IconOverlap) * scale))
}
