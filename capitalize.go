package piscine

func Capitalize(s string) string {
	str := []rune(s)
	if str[0] >= 'a' && str[0] <= 'z' {
		str[0] -= 32
	}

	for i := 1; i < StrLen(s); i++ {
		if VerCharacter(str, i) {
			if str[i] >= 'A' && str[i] <= 'Z' {
				if VerCharacter(str, i-1) {
					str[i] += 32
				}
			}
		} else {
			if i != StrLen(s)-1 {
				if str[i+1] >= 97 && str[i+1] <= 122 {
					str[i+1] -= 32
				}
			}
		}
	}
	return string(str)
}

func VerCharacter(str []rune, i int) bool {
	if (str[i] >= '0' && str[i] <= '9') ||
		(str[i] >= 'a' && str[i] <= 'z') ||
		(str[i] >= 'A' && str[i] <= 'Z') {
		return true
	}

	return false
}
