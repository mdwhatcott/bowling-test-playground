package bowling

func Score(rolls []int) int {
	return score(rolls, 1, 0)
}
func score(rolls []int, frame, total int) int {
	if len(rolls) == 0 {
		return total
	}
	if frame > frameCount {
		return total
	}
	if rolls[0] == allPins { // strike
		return score(rolls[1:], frame+1, total+allPins+rolls[1]+rolls[2])
	}
	if rolls[0]+rolls[1] == allPins { // spare
		return score(rolls[2:], frame+1, total+allPins+rolls[2])
	}
	return score(rolls[2:], frame+1, total+rolls[0]+rolls[1])
}

const (
	allPins    = 10
	frameCount = 10
)
