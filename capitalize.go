package piscine

func Capitalize(s string) string {
	str := []rune(s)
	if str[0] >= 97 && str[0] <= 122 {
		str[0] -= 32
	}
	for i := 0; i < StrLen(s); i++ {
		if !((str[i] >= '0' && str[i] <= '9') ||
			(str[i] >= 'a' && str[i] <= 'z') ||
			(str[i] >= 'A' && str[i] <= 'Z')) {
			if i != StrLen(s)-1 {
				if str[i+1] >= 'a' && str[i+1] <= 'z' {
					str[i+1] -= 32
				}
			}
		}
	}
	return string(str)
}
