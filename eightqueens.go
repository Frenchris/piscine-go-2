package piscine

import (
	"github.com/01-edu/z01"
)

var board [8][8]int

func EightQueens() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = j + 1
		}
	}

	seeColumn(0, 0)

}

func seeColumn(row int, col int) {
	for col < 8 {
		if seeQueen(row, col) {
			previous := board[row][col]
			board[row][col] = 0

			if row == 8-1 {
				for i := 0; i < 8; i++ {
					for j := 0; j < 8; j++ {
						if board[i][j] == 0 {
							helperPrint := j + 1
							PrintNbr(helperPrint)
						}
					}
				}
				z01.PrintRune('\n')
			} else {
				seeColumn(row+1, 0)

			}
			board[row][col] = previous

		}
		col++
	}
}

func seeQueen(row int, col int) bool {

	//check for rows
	for i := 0; i < 8; i++ {
		if board[row][i] == 0 {
			return false
		}
	}
	//check for collumns
	for i := 0; i < 8; i++ {
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
	for i >= 0 && j < 8 {
		if board[i][j] == 0 {
			return false
		}

		i--
		j++
	}
	return true
}
