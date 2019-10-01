package piscine

func ForEach(f func(int), arr []int) {
	for i := 0; i < ArrayIntLength(arr); i++ {
		f(arr[i])
	}

}
