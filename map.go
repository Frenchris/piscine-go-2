package piscine

func Map(f func(int) bool, arr []int) []bool {
	var result []bool

	for i := 0; i < ArrayIntLength(arr); i++ {
		result = append(result, f(arr[i]))
	}

	return result
}
