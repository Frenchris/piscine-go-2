package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {

	negative := false
	if baseFrom == baseTo {
		return nbr
	} else if baseTo == "0123456789" {
		return Itoa(AtoiBase(nbr, baseFrom))
	} else if baseFrom == "0123456789" {
		SeeNumber(Atoi(nbr), baseTo, negative)
	} else {
		helper := string(AtoiBase(nbr, "0123456789"))
		SeeNumber(Atoi(helper), baseTo, negative)
	}
	return ""
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
