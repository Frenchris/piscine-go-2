package piscine

func Unmatch(arr []int) int {

	nbrs := ""
	for i := 0; i < ArrayIntLength(arr); i++ {
		nbrs += Itoa(arr[i])
	}
	for j := 0; j < ArrayIntLength(arr); j++ {

		//	fmt.Println(nbrs, Itoa(j))
		//	fmt.Println(CountInstances(nbrs, Itoa(j)))

		if CountInstances(nbrs, Itoa(j)) == 1 {

			return j
		}
	}
	return -1
}
