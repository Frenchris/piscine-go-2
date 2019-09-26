package piscine

func ToLower(s string) string {
	str := []rune(s)

	for i := 0; i < StrLen(s); i++ {
		if str[i] >= 65 && str[i] <= 90 {
			str[i] += 32
		}
	}

	return string(str)

}
