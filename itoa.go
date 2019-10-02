package piscine

var result1 string

func Itoa(nbr int) string {
	t := 1

	if nbr < 0 {
		result1 = result1 + "-"
		t = -1
	}
	if nbr != 0 {
		q := (nbr / 10) * t
		if q != 0 {
			Itoa(q)
		}
		d := ((nbr % 10) * t) + '0'
		result1 = result1 + string(rune(d))
	} else {
		result1 = result1 + "0"
	}
	result := result1
	result1 = ""
	return result
}
