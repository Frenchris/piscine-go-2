package piscine

func NRune(s string, n int) rune {

	stringByRunes := []rune(s)

	for i := 1; i <= ArrayStrLength(s); i++ {
		if i == n {
			return stringByRunes[i-1]
		}
	}
	return '0'
}
