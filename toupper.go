package piscine

func ToUpper(s string) string {

	str := []rune(s)

	for i := 0; i < StrLen(s); i++ {
		if str[i] >= 97 && str[i] <= 122 {
			str[i] -= 32
		}
	}

	return string(str)

}
