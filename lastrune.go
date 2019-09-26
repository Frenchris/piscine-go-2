package piscine

func LastRune(s string) rune {

	stringByRunes := []rune(s)

	return stringByRunes[ArrayStrLength(s)-1]

}
