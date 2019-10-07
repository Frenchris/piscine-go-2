package piscine

func AtoiBase(s string, base string) int {
	noValid := ""

	if base == "0123456789" {

		return Atoi(s)
	}

	if StrLen(base) < 2 {
		noValid = "NV"
	} else {
		for i := 0; i < StrLen(base)-1; i++ {
			for j := i + 1; j < StrLen(base); j++ {
				if base[i] == base[j] ||
					base[i] == '+' || base[i] == '-' ||
					base[j] == '+' || base[j] == '-' {
					noValid = "NV"
				}
			}
		}
	}

	if noValid == "NV" {
		return 0
	}

	return PrintNumberAtoiBase(s, base)

}
