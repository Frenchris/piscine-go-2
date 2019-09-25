package piscine

import (
	"github.com/01-edu/z01"
)

func Raid1d(x,y int){
	
	cantoEsqSup := 'A'
	cantoDirSup := 'C'
	cantoEsqInf := 'A'
	cantoDirInf := 'C'
	ladosHor := 'B'
	ladosVer := 'B'


	if x==1 && y==1{
		
		z01.PrintRune(cantoEsqSup)
		z01.PrintRune('\n')
		
	}else{

		z01.PrintRune(cantoEsqSup)
		
		if x!=1{
			for i := 0; i < x - 2; i++ {
				z01.PrintRune(ladosHor)
			}

			z01.PrintRune(cantoDirSup)
		}
		z01.PrintRune('\n')

		for i := 0; i < y - 2; i++ {
			
			z01.PrintRune(ladosVer)
			
			if x!=1 {
				for j := 0; j < x - 2; j++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune(ladosVer)
			}
			
				z01.PrintRune('\n')
		}
	if y!=1 {
			
			z01.PrintRune(cantoEsqInf)
			if x!=1{
				for i := 0; i < x - 2; i++ {
					z01.PrintRune(ladosHor)
				}

				z01.PrintRune(cantoDirInf)
			}
			z01.PrintRune('\n')		
		}	
	}
}