package piscine

func ActiveBits(n int) int {

	binary := ConvertFromDecimal(n, "01")

	result := CountInstances(binary, "1")

	return result
}
