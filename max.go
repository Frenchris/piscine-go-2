package piscine

func Max(arr []int) int {
	big := -2147483648

	for i := 0; i < ArrayIntLength(arr); i++ {
		if arr[i] > big {
			big = arr[i]
		}
	}
	return big
}
