// TRANSPORT ZEROES

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func stretch(x []int, y int, z int) {
	fulcrum := z
	for i := y; i < len(x); i++ {
		if x[i] != 0 {
			x[i], x[fulcrum] = x[fulcrum], x[i]
			fulcrum++
		}
	}
}

func zeroMover(board []int) {
everything:
	for i := 0; i < len(board)-1; i++ {
		if board[i] == 0 && board[i+1] != 0 {
			board[i], board[i+1] = board[i+1], board[i]
		} else if i < len(board)-2 && board[i] == 0 && board[i+1] == 0 {

			for j := i; j < len(board)-1; j++ {
				if board[j] != 0 {
					board[i], board[j] = board[j], board[i]
				}
				stretch(board, i, j)
				break everything
			}
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
