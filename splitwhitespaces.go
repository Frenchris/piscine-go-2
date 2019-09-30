package piscine

import "fmt"

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
		if sArray[i] == '\n' || sArray[i] == '\t' || sArray[i] == ' ' || sArray[i] == '\v' {

			fmt.Println(sArray)
			fmt.Println(string(sArray))
			fmt.Println()
			stri = string(sArray[0:i])
			if string(sArray[i+1]) == " " {
				sArray = sArray[i+2 : StrLen(string(sArray))]
			} else {
				sArray = sArray[i+1 : StrLen(string(sArray))]
			}
			fmt.Println(sArray)
			fmt.Println(string(sArray))
			result[j] = stri
			i = -1

			j++
			//fmt.Println(string(sArray))
		}
		if j == count-1 {
			result[j] = string(sArray[0:StrLen(string(sArray))])
		}
	}

	fmt.Println([]rune(result[0]))
	fmt.Println()
	return result
}
