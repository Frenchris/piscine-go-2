package piscine

func Index(s string, toFind string) int {

	sArr := []rune(s)
	toFindArr := []rune(toFind)
	count := 0
	i := 0
	count2 := 0

	for j := 0; j < StrLen(s); j++ {
		if sArr[j] == toFindArr[i] {
			count++
			i++

			if count == StrLen(toFind) {
				i = 0
				count = 0
				count2++

			}
		}

	}

	if count2 == 0 {
		count2--
	}

	return count2
}
