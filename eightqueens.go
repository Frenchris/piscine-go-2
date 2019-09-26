package piscine

import (
	"github.com/01-edu/z01"
)

//size of the board of chees
const size = 8

//basically the board, inicializado a false
var board [size][size]int

func EightQueens() {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			board[x][y] = y + 1
		}
	}

	seeColumn(0, 0)

}

func seeColumn(row int, col int) {
	for col < size {
		if seeQueen(row, col) {
			save := board[row][col]
			board[row][col] = 0

			if row == size-1 {
				for x := 0; x < size; x++ {
					for y := 0; y < size; y++ {
						if board[x][y] == 0 {
							helperPrint := y + 1
							PrintNbr(helperPrint)
						}
					}
				}
				z01.PrintRune('\n')
			} else {
				seeColumn(row+1, 0)

			}
			board[row][col] = save

		}
		col++
	}
}

func seeQueen(row int, col int) bool {

	//check for rows
	for i := 0; i < size; i++ {
		if board[row][i] == 0 {
			return false
		}
	}
	//check for collumns
	for i := 0; i < size; i++ {
		if board[i][col] == 0 {
			return false
		}
	}
	i := row
	j := col
	//check the diagonals
	for i >= 0 && j >= 0 {
		if board[i][j] == 0 {
			return false
		}
		i--
		j--
	}

	i = row
	j = col
	//check diagonals
	for i >= 0 && j < size {
		if board[i][j] == 0 {
			return false
		}

		i--
		j++
	}
	return true
}
