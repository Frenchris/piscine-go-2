package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := 0; i < ArrayStrLength(args); i++ {
		result += args[i]
		if i != ArrayStrLength(args)-1 {
			result += "\n"
		}
	}
	return result
}
