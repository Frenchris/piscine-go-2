package piscine

func Index(s string, toFind string) int {

	sArr := []rune(s)
	toFindArr := []rune(toFind)
	count := 0
	i := 0
	pos := 0
	firstLetter := 0

	if toFind == "" {
		return 0
	} else {
		for j := 0; j < StrLen(s); j++ {
			if sArr[j] == toFindArr[i] {
				count++
				if firstLetter == 0 {
					pos = j
					firstLetter++
				}
				i++

				if count == StrLen(toFind) {
					i = 0
					return pos
				}
			} else if count > 0 {
				count = 0
				i = 0
				firstLetter = 0
				j--
			}
		}

		if count == 0 {
			pos = -1
		}
	}

	return pos

}
