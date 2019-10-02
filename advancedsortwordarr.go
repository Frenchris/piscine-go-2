package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < ArrayStrLength(array)-1; i++ {
		for j := i + 1; j < ArrayStrLength(array); j++ {
			wordI := array[i]
			wordJ := array[j]

			k := 0
			l := 0

			for k < StrLen(wordI) && l < StrLen(wordJ) {
				if f(wordI, wordJ) == 1 {
					array[i], array[j] = array[j], array[i]
					k = StrLen(wordI)
					l = StrLen(wordJ)
				} else if f(wordI, wordJ) == -1 {
					k = StrLen(wordI)
					l = StrLen(wordJ)
				} else if f(wordI, wordJ) == 0 {
					k++
					l++
				}

			}
		}
	}
}
