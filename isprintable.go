package piscine

func IsPrintable(str string) bool {

	s := []rune(str)

	for i := 0; i < StrLen(str); i++ {
		if s[i] >= 0 && s[i] <= 31 {
			return false
		}
	}
	return true
}
