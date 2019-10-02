package piscine

import "fmt"

func SortWordArr(array []string) {

	for i := 0; i < ArrayStrLength(array)-1; i++ {
		for j := i + 1; j < ArrayStrLength(array); j++ {
			wordI := array[i]
			wordJ := array[j]

			k := 0
			l := 0

			for k < StrLen(wordI) && l < StrLen(wordJ) {
				if rune(wordI[k]) > rune(wordJ[l]) {
					array[i], array[j] = array[j], array[i]
					fmt.Println("asdasdasda")
					k = StrLen(wordI)
					l = StrLen(wordJ)

				} else if rune(wordI[k]) < rune(wordJ[l]) {
					k = StrLen(wordI)
					l = StrLen(wordJ)

				} else {
					k++
					l++
				}

			}

		}
	}

}
