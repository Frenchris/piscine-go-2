package piscine

func MakeRange(min, max int) []int {

	result := make([]int, max-min)

	if min > max {
		return nil
	}

	for i := 0; i < max-min; i++ {
		result[i] = i + min
	}

	return result

}
