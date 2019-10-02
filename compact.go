package piscine

func Compact(ptr *[]string) int {

	slice := *ptr
	count := 0

	for i := 0; i < ArrayStrLength(slice); i++ {
		if slice[i] != "" {
			count++
		}
	}
	var array []string
	for i := 0; i < ArrayStrLength(slice); i++ {
		if slice[i] != "" {
			array = append(array, slice[i])
		}
	}

	*ptr = array

	return count
}
