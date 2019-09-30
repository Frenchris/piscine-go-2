package piscine

func SplitWhiteSpaces(str string) []string {
	count := 0
	sArray := []rune(str)
	for i := 0; i < StrLen(str); i++ {
		if sArray[i] == '\n' || sArray[i] == '\t' || sArray[i] == ' ' {
			count++
		}
	}
	result := make([]string, count+1)
	stri := ""
	j := 0
	for i := 0; i < StrLen(string(sArray)); i++ {
		if sArray[i] == '\n' || sArray[i] == '\t' || sArray[i] == ' ' {
			stri = string(sArray[0:i])
			sArray = sArray[i+1 : StrLen(string(sArray))]
			result[j] = stri
			i = 0
			j++
			//fmt.Println(string(sArray))
		}
		if j == count {
			result[j] = string(sArray[0:StrLen(string(sArray))])
		}
	}

	//stri = string([]byte{str[0], str[1]})

	return result
}
