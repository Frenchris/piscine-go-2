package piscine


func IterativePower(nb int, power int) int{
	helper := nb
	if power == 0{
		nb = 1
	} else {
		for i := 1; i < power; i++ {
			nb *= helper
		}
	}
	return nb
}