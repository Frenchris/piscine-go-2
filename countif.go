package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	for i := 0; i < ArrayStrLength(tab); i++ {
		if f(tab[i]) {
			count++
		}
	}

	return count
}
