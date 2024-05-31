package main

import (
	"fmt"
	"math/rand"
)

func createField(width int, height int, mines int) [][]string {
	field := make([][]string, height+2)
	// an array that contains the coordinates of the mines
	minesCoordinates := make([][]int, mines)
	// init field
	for i := 0; i < height+2; i++ {
		field[i] = make([]string, width+2)
		for j := 0; j < width+2; j++ {
			field[i][j] = "."
		}
	}
	// init minesCoordinates
	for i := 0; i < mines; i++ {
		minesCoordinates[i] = make([]int, 2)
	}
	// fill field with mines
	for i := 0; i < mines; i++ {
		for {
			x := rand.Intn(width) + 1
			y := rand.Intn(height) + 1
			place := false
			// check if there is already a mine at this position
			if minesCoordinates[x][0] == x && minesCoordinates[y][1] == y {
				continue
			}
			// place mine
			minesCoordinates[i][0] = x
			minesCoordinates[i][1] = y

			field[y][x] = "*"

			place = true
			if place {
				break
			}
		}

	}
	// fill field with numbers
	for i := 1; i < height+1; i++ {
		for j := 1; j < width+1; j++ {
			if field[i][j] == "." {
				count := 0
				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if field[k][l] == "*" {
							count++
						}
					}
				}
				if count > 0 {
					field[i][j] = fmt.Sprintf("%d", count)
				}
			}
		}
	}
	return field
}
func printField(field [][]string) {
	for i := 1; i < len(field)-1; i++ {
		for j := 1; j < len(field[i])-1; j++ {
			fmt.Printf("%s ", field[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := createField(20, 25, 99)
	printField(board)
	// fmt.Println(board)
}
