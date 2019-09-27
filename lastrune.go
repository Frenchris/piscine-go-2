package piscine

func LastRune(s string) rune {

	stringByRunes := []rune(s)

	return stringByRunes[StrLen(s)-1]

}
