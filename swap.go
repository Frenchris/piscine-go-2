package piscine

func Swap(a *int, b *int){
	c := *b
	d := *a
	
	*a = c
	*b = d
}