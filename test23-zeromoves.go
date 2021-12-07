// TRANSPORT ZEROES

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func stretch(x []int, y int, z int) {
	for ; z < len(x); z++ {
		if x[z] != 0 {
			x[z], x[y] = x[y], x[z]
			y++
		}
	}
}

func zeroMover(board []int) {
everything:
	for i := 0; i < len(board)-1; i++ {
		if board[i] == 0 && board[i+1] != 0 {
			board[i], board[i+1] = board[i+1], board[i]
		} else if i < len(board)-2 && board[i] == 0 && board[i+1] == 0 {
			j := i + 2
			stretch(board, i, j)
			break everything
		}
	}
}

func main() {
	board := make([]int, 10000000, 10000000)
	for i := 0; i < len(board); i++ {
		board[i] = rand.Intn(2)
	}
	//fmt.Println(board)
	start := time.Now()
	zeroMover(board)
	elapsed := time.Since(start)
	//fmt.Println(board)
	fmt.Println(elapsed)
}
