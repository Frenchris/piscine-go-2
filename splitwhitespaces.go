package piscine

func SplitWhiteSpaces(str string) []string {
	count := 0
	sArray := []rune(str)
	for i := 0; i < StrLen(str); i++ {
		if i != StrLen(str)-1 {
			if sArray[i] == '\n' || sArray[i] == '\t' || sArray[i] == ' ' {
				if !(sArray[i+1] == '\n' || sArray[i+1] == '\t' || sArray[i+1] == ' ') {
					count++
				}
			}
		}
	}
	result := make([]string, count+1)
	stri := ""
	j := 0
	for i := 0; i < StrLen(string(sArray)); i++ {
		if sArray[i] == '\n' || sArray[i] == '\t' || sArray[i] == ' ' || sArray[i] == '\v' {

			stri = string(sArray[0:i])
			if string(sArray[i+1]) == " " {
				sArray = sArray[i+2 : StrLen(string(sArray))]

			} else {
				sArray = sArray[i+1 : StrLen(string(sArray))]
			}
			result[j] = stri
			i = -1
			j++
			//fmt.Println(string(sArray))
		}
		if j == count {
			result[j] = string(sArray[0:StrLen(string(sArray))])
		}
	}
	return result
}
