package piscine

/*==================================================================
  ===================						  ======================
  ===========			 	   INTS  				 ===============
  ===================					      ======================
  ==================================================================*/

//ConvertFromDecimal returns nbr in base indicated by base
//ex: 12,"01" -> 1100
func ConvertFromDecimal(nbr int, base string) string {
	result := ""

	for nbr >= StrLen(base) {
		result += string(base[nbr%StrLen(base)])
		nbr /= StrLen(base)
	}
	result += string(base[nbr])

	return StrRev(result)
}

//PrintNumberAtoiBase returns nbr in base indicated by base, in decimal base
//ex: "101","01" -> 5
func PrintNumberAtoiBase(nbr string, base string) int {

	arrayBase := []rune(base)
	arrayNbr := []rune(nbr)
	result := 0

	for i := 0; i < StrLen(nbr); i++ {
		for j := 0; j < StrLen(base); j++ {

			if arrayNbr[i] == arrayBase[j] {

				result += j * (Power(StrLen(base), StrLen(nbr)-i-1))
			}
		}
	}

	return result

}

//Power returns the result of n^m
func Power(n int, m int) int {
	helper := n
	if m == 0 {
		n = 1
	} else {
		for i := 1; i < m; i++ {
			n *= helper
		}
	}
	return n
}

var resultItoa string

//Itoa returns nbr in string type
func Itoa(nbr int) string {
	t := 1

	if nbr < 0 {
		resultItoa += "-"
		t = -1
	}
	if nbr != 0 {
		q := (nbr / 10) * t
		if q != 0 {
			Itoa(q)
		}
		d := ((nbr % 10) * t) + '0'
		resultItoa += string(rune(d))
	} else {
		resultItoa += "0"
	}
	result := resultItoa
	resultItoa = ""
	return result
}

/*==================================================================
  ===================						  ======================
  ===========				  STRINGS				 ===============
  ===================					      ======================
  ==================================================================*/

//CommonLetters returns a string with the letters in common
//of str1 and str2 in order of str1
func CommonLetters(str1, str2 string) string {

	first := []rune(str1)
	second := []rune(str2)

	result := ""
	i := 0

	for i < ArrayRuneLength(first) {
		for j := 0; j < ArrayRuneLength(second); j++ {
			if i >= ArrayRuneLength(first) {
				break
			}
			if first[i] == second[j] {
				result += string(first[i])
				i++
				j = -1
			}
			if StrLen(result) == ArrayRuneLength(first) {
				break
			}
		}
		i++
	}
	return result
}

//RemoveDuplicates returns s wit no duplicate letters
func RemoveDuplicates(s string) string {

	arr := []rune(s)

	equals := false

	if CountInstances(s, string(arr[0])) == ArrayRuneLength(arr) {
		arr = arr[0:1]
	} else {
		for i := 0; i < ArrayRuneLength(arr)-1; i++ {
			for j := i + 1; j < ArrayRuneLength(arr); j++ {

				if equals {
					i = 0
					j = i + 1
				}

				if arr[i] == arr[j] {
					helper := arr
					helper = nil

					for k := 0; k < j; k++ {
						helper = append(helper, arr[k])
					}

					for k := j + 1; k < ArrayRuneLength(arr); k++ {
						helper = append(helper, arr[k])
					}
					arr = helper
					equals = true
				} else {
					equals = false
				}
			}
		}
	}

	return string(arr)

}

//CountInstances returns the number of times that toFind appears in s
func CountInstances(s string, toFind string) int {
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

/*==================================================================
  ===================						  ======================
  ===========					ARRAYS				 ===============
  ===================					      ======================
  ==================================================================*/

//ArrayStrLength returns the length of the array str []string
func ArrayStrLength(str []string) int {

	length := 0

	for range str {
		length++
	}

	return length
}

//ArrayRuneLength returns the length of the array str []rune
func ArrayRuneLength(str []rune) int {

	length := 0

	for range str {
		length++
	}

	return length
}

//ArrayIntLength returns the length of the array array []int

func ArrayIntLength(array []int) int {

	length := 0
	for range array {
		length++
	}
	return length

}
