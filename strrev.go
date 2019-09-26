package piscine

func StrRev(s string) string {
	strLe := StrLen(s)

	stri := []byte(s)

	for i := 0; i < strLe/2; i++ {
		stri[i], stri[strLe-1-i] = stri[strLe-1-i], stri[i]
	}

	reverse := string(stri)

	return reverse
}
