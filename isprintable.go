package piscine

func IsPrintable(str string) bool {

	for i := 0; i < StrLen(str); i++ {
		if !(str[i] >= 0 && str[i] <= 31) {
			return false
		}
	}
	return true
}
