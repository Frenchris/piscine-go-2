package piscine

func IsAlpha(str string) bool {

	s := []rune(str)
	count := 0

	for i := 0; i < StrLen(str); i++ {
		if VerCharacter(s, i) {
			count++
		} else {
			return false
		}
	}

	if count == StrLen(str) {
		return true
	}

	return false
}
