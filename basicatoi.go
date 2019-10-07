package piscine

func BasicAtoi(s string) int {
	stri := []rune(s)

	res := 0

	for i := 0; i < StrLen(s); i++ {
		exp := StrLen(s) - 1 - i
		if stri[i] == '0' {
			res += 0 * Power(10, exp)
		} else if stri[i] == '1' {
			res += 1 * Power(10, exp)
		} else if stri[i] == '2' {
			res += 2 * Power(10, exp)
		} else if stri[i] == '3' {
			res += 3 * Power(10, exp)
		} else if stri[i] == '4' {
			res += 4 * Power(10, exp)
		} else if stri[i] == '5' {
			res += 5 * Power(10, exp)
		} else if stri[i] == '6' {
			res += 6 * Power(10, exp)
		} else if stri[i] == '7' {
			res += 7 * Power(10, exp)
		} else if stri[i] == '8' {
			res += 8 * Power(10, exp)
		} else if stri[i] == '9' {
			res += 9 * Power(10, exp)
		}
	}

	return res
}
