package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {

	helper := AtoiBase(nbr, baseFrom)
	return ConvertFromDecimal(helper, baseTo)
}
