package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {

	helper := AtoiBase(nbr, baseFrom)
	return ConvertFromDecimal(helper, baseTo)
}

func ConvertFromDecimal(nbr int, base string) string {
	result := ""

	for nbr >= StrLen(base) {
		result += string(base[nbr%StrLen(base)])
		nbr /= StrLen(base)
	}
	result += string(base[nbr])

	return StrRev(result)
}
