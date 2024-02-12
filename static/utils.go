package static

func SpeedToMovement(increment float64, speed float64) float64 {
	return increment * 5 * speed
}

func LetterCount(offset float64, scale float64) int {
	return int((ResolutionWidth - offset) / ((IconWidth - IconOverlapInRow) * scale))
}
