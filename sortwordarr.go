package piscine

func SortWordArr(array []string) {

	for i := 0; i < ArrayStrLength(array)-1; i++ {
		for j := i + 1; j < ArrayStrLength(array); j++ {
			wordI := array[i]
			wordJ := array[j]
			if wordI[0] > wordJ[0] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

}
