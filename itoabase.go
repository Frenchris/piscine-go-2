package piscine

func ItoaBase(value, base int) string {
	baseS := "0123456789ABCDEF"
	baseS = baseS[0:base]
	result := ConvertFromDecimal(value, baseS)
	return result
}
