package piscine

func ActiveBits(n int) int {

	binary := ConvertBase(Itoa(n), "0123456789", "01")

	return CountInstances(binary, "1")
}
