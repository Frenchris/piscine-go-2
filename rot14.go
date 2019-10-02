package piscine

func Rot14(str string) string {

	s := []rune(str)

	for i := 0; i < StrLen(str); i++ {
		if (s[i] >= 'A' && s[i] <= 'L') ||
			(s[i] >= 'a' && s[i] <= 'l') {
			s[i] += 14
		} else if (s[i] > 'L' && s[i] <= 'Z') ||
			(s[i] > 'l' && s[i] <= 'z') {
			s[i] -= 12
		}
	}

	return string(s)

}
