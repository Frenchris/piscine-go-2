package piscine

func SortIntegerTable(table []int) {

	for i := 0; i < ArrayIntLength(table)-1; i++ {
		for j := i + 1; j < ArrayIntLength(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}

}
