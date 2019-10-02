package piscine

func Unmatch(arr []int) int {

	nbrs := ""
	for i := 0; i < ArrayIntLength(arr); i++ {
		nbrs += Itoa(arr[i])
	}
	for j := 0; j < ArrayIntLength(arr); j++ {
		if CountInstances(nbrs, Itoa(j)) != -1 {
			if !(CountInstances(nbrs, Itoa(j))%2 == 0) {

				return j
			}
		}
	}
	return -1
}
