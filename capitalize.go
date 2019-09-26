package piscine

func Capitalize(s string) string {
	str := []rune(s)
	if str[0] >= 97 && str[0] <= 122 {
		str[0] -= 32
	}
	for i := 0; i < StrLen(s); i++ {
		if (str[i] >= 48 && str[i] <= 57) ||
			(str[i] >= 97 && str[i] <= 122) ||
			(str[i] >= 65 && str[i] <= 90) {

		} else {
			if i != StrLen(s)-1 {
				if str[i+1] >= 97 && str[i+1] <= 122 {
					str[i+1] -= 32
				}
			}
			if i != StrLen(s)-1 {
				if str[i+1] >= 65 && str[i+1] <= 90 {
					str[i+1] += 32
				}
			}
		}
	}
	return string(str)
}
