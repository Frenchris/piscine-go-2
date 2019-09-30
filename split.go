package piscine

func Split(str, charset string) []string {

	striArray := []rune(str)
	helper := ""
	var result []string

	for i := 0; i < CountInstances(str, charset); i++ {
		helper = string(striArray[0:Index(string(striArray), charset)])
		println(helper)
		result = append(result, helper)

		striArray = striArray[Index(string(striArray), charset)+StrLen(charset) : StrLen(string(striArray))]
	}
	result = append(result, string(striArray[0:StrLen(string(striArray))]))

	return result
}

// result := make([]string, count+1)
// stri := ""
// j := 0
// for i := 0; i < StrLen(string(sArray)); i++ {
// 	if string(sArray[i]) == charset {

// 		stri = string(sArray[0:i])
// 		if string(sArray[i+1]) == " " {
// 			sArray = sArray[i+2 : StrLen(string(sArray))]

// 		} else {
// 			sArray = sArray[i+1 : StrLen(string(sArray))]
// 		}
// 		result[j] = stri
// 		i = -1
// 		j++
// 		//fmt.Println(string(sArray))
// 	}
// 	if j == count {
// 		result[j] = string(sArray[0:StrLen(string(sArray))])
// 	}
// }
// return result
