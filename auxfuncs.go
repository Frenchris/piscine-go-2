package piscine

/*==================================================================
  ===================						  ======================
  ===========			 	   INTS  				 ===============
  ===================					      ======================
  ==================================================================*/

//PrimeNumbers returns an array of prime ints, from 2 to nbr
func PrimeNumbers(nbr int) []int {
	var result []int
	for i := 2; i <= nbr; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

//ConvertFromDecimal returns nbr in decimal base converted to a number with a base indicated by base
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

var result string

//Itoa returns nbr in string type
func Itoa(nbr int) string {
	result = ""
	t := 1

	if nbr < 0 {
		result += "-"
		t = -1
	}
	if nbr != 0 {
		q := (nbr / 10) * t
		if q != 0 {
			Itoa(q)
		}
		d := ((nbr % 10) * t) + '0'
		result += string(rune(d))
	} else {
		result += "0"
	}

	return result
}

//IsOperator returns true if str is an operator, false otherwise
func IsOperator(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "/" || str == "%" {
		return true
	}
	return false
}

//IsNumber returns true if str is a number, false otherwise
func IsNumber(str string) bool {

	if str[0] == '-' && StrLen(str) != 1 {
		str = str[1:StrLen(str)]
	}

	for i := 0; i < StrLen(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			return false
		}
	}
	return true
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

func LowerCase(s string) string {
	str := []rune(s)
	if str[0] >= 'A' && str[0] <= 'Z' {
		str[0] += 32
	}
	for i := 1; i < StrLen(s); i++ {
		if VerCharacter(str, i) {
			if str[i] >= 'A' && str[i] <= 'Z' {
				if VerCharacter(str, i-1) {
					str[i] += 32
				}
			}
		} else {
			if i != StrLen(s)-1 {
				if str[i+1] >= 97 && str[i+1] <= 122 {
					str[i+1] -= 32
				}
			}
		}
	}
	return string(str)
}

//IsStrictlyString returns true if str countains only letters
func IsStrictlyString(str string) bool {
	s := []rune(str)
	for i := 0; i < StrLen(str); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z' ||
			s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}
	return true
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

//RemoveFromStrArray removes from an array every instance given by charset
func RemoveFromStrArray(array []string, charset string) []string {
	var result []string
	for i := 0; i < ArrayStrLength(array); i++ {
		if array[i] != charset {
			result = append(result, array[i])
		}
	}

	return result
}

//CheckStrings returns false if in an array of "ints" exists any "strings"
//ex: {"2", "6", "hello","4"} -> false
//	  {"1", "7", "9", "1"} -> true
func CheckStrings(arr []string) bool {
	for i := 0; i < ArrayStrLength(arr); i++ {
		if IsStrictlyString(arr[i]) {
			return false
		}
	}
	return true
}
