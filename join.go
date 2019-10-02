package piscine

func Join(strs []string, sep string) string {

	result := ""
	for i := 0; i < ArrayStrLength(strs)-1; i++ {
		result += strs[i] + sep
	}
	result += strs[ArrayStrLength(strs)-1]

	return result
}
