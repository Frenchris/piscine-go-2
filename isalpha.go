package piscine

func IsAlpha(str string) bool {

	s := []rune(str)
	//count := 0

	for i := 0; i < StrLen(str); i++ {
		if !VerCharacter(s, i) {
			return false
		}
	}
	return true
}
