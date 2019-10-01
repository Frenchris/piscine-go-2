package piscine

func Any(f func(string) bool, arr []string) bool {

	for i := 0; i < ArrayStrLength(arr); i++ {
		if f(arr[i]) {
			return true
		}
	}
	return false
}
