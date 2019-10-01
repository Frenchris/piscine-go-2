package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	ascOrder := 1
	decOrder := 1

	for i := 0; i < ArrayIntLength(tab)-1; i++ {
		if f(tab[i], tab[i+1]) >= 0 {
			ascOrder++
		} else if f(tab[i], tab[i+1]) <= 0 {
			decOrder++
		}
	}

	if ascOrder == ArrayIntLength(tab) || decOrder == ArrayIntLength(tab) {
		return true
	}
	return false

}

func F(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
