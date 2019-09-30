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
