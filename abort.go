package piscine

func Abort(a, b, c, d, e int) int {
	var array []int
	array = append(array, a)
	array = append(array, b)
	array = append(array, c)
	array = append(array, d)
	array = append(array, e)

	SortIntegerTable(array)
	return array[2]
}
