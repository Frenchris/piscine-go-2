package piscine

func IsNumeric(str string) bool {

	for i := 0; i < StrLen(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			return false
		}
	}
	return true
}
