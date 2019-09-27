package piscine

func IsPrintable(str string) bool {

	s := []rune(str)

	for i := 0; i < StrLen(str); i++ {
		if !(s[i] >= 33 && s[i] <= 126) {
			return false
		}
	}
	return true
}
