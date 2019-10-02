package piscine

func ActiveBits(n int) uint {

	binary := ConvertFromDecimal(n, "01")

	result := CountInstances(binary, "1")

	return uint(result)
}
