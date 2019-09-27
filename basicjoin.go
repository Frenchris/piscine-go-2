package piscine

func BasicJoin(strs []string) string {

	result := ""

	for i := 0; i < ArrayStrLength(strs); i++ {
		result += strs[i]
	}
	return result
}
