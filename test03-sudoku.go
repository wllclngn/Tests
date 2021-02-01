package main

import (
	"fmt"
	"time"
)

// Thx, Josh, for the "k < 10" part, not "k < 9".

// Initiate matrix search
func neo(a [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if a[i][j] == 0 {
				for k := 0; k < 10; k++ {
					if gnosis(i, j, k, a) {
						a[i][j] = k
						neo(a)
						a[i][j] = 0
					}
				}
				return
			}
		}
	}
	fmt.Println("Solved Matrix:")
	for q := 0; q < len(a); q++ {
		fmt.Println(q, a[q])
	}
	return
}

// Trek over matrix to deduce whether to place number
func gnosis(x int, y int, z int, a [][]int) bool {
	for l := 0; l < 9; l++ {
		if a[x][l] == z {
			return false
		}
	}
	for m := 0; m < 9; m++ {
		if a[m][y] == z {
			return false
		}
	}
	xSquare := ((x / 3) * 3)
	ySquare := ((y / 3) * 3)
	for n := 0; n < 3; n++ {
		for o := 0; o < 3; o++ {
			if a[xSquare+n][ySquare+o] == z {
				return false
			}
		}
	}
	return true
}

func main() {
	/// #1 ///
	puzzle1 := [][]int{{0, 0, 0, 0, 1, 2, 3, 0, 0}, {0, 1, 0, 0, 4, 5, 0, 0, 0}, {6, 0, 0, 0, 7, 0, 0, 0, 0},
		{7, 4, 0, 0, 0, 0, 8, 9, 2}, {0, 0, 3, 0, 0, 0, 6, 0, 0}, {5, 8, 9, 0, 0, 0, 0, 1, 3},
		{0, 0, 0, 0, 5, 0, 0, 0, 7}, {0, 0, 0, 1, 8, 0, 0, 4, 0}, {0, 0, 2, 9, 6, 0, 0, 0, 0}}
	start := time.Now()
	neo(puzzle1)
	elapsed := time.Since(start)
	fmt.Println("Start:", start)
	fmt.Println("Elapsed:", elapsed)
	/// #2 ///
	puzzle2 := [][]int{{7, 0, 0, 0, 0, 0, 4, 0, 0}, {0, 2, 0, 0, 7, 0, 0, 8, 0}, {0, 0, 3, 0, 0, 8, 0, 0, 9},
		{0, 0, 0, 5, 0, 0, 3, 0, 0}, {0, 6, 0, 0, 2, 0, 0, 9, 0}, {0, 0, 1, 0, 0, 7, 0, 0, 6},
		{0, 0, 0, 3, 0, 0, 9, 0, 0}, {0, 3, 0, 0, 4, 0, 0, 6, 0}, {0, 0, 9, 0, 0, 1, 0, 0, 5}}
	start2 := time.Now()
	neo(puzzle2)
	elapsed2 := time.Since(start2)
	fmt.Println("Start:", start2)
	fmt.Println("Elapsed:", elapsed2)
	/// #3 ///
	// "The World's Hardest Sudoku Puzzle"
	puzzle3 := [][]int{{8, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 3, 6, 0, 0, 0, 0, 0}, {0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0}, {0, 0, 0, 0, 4, 5, 7, 0, 0}, {0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8}, {0, 0, 8, 5, 0, 0, 0, 1, 0}, {0, 9, 0, 0, 0, 0, 4, 0, 0}}
	start3 := time.Now()
	neo(puzzle3)
	elapsed3 := time.Since(start3)
	fmt.Println("Start:", start3)
	fmt.Println("Elapsed:", elapsed3)
}
