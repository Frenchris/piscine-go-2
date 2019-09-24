package piscine

func StrLen(str string) int {

	length := 0

	for count,_ := range str {
		length += 1
	}

	return length
}