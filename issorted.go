package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	for i := 0; i < ArrayIntLength(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			return false
		}
	}
	return true
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
