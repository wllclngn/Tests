// TRANSPORT ZEROES

package main

import (
	"fmt"
	"math/rand"
)

func zeroMover(board []int) {

	for i := 0; i < len(board)-1; i++ {
		if board[i] == 0 && board[i+1] != 0 {
			board[i], board[i+1] = board[i+1], board[i]
			continue
		} else if i < len(board)-2 && board[i] == 0 && board[i+1] == 0 {
			j := zerored(board, i+2)
			board[i], board[j] = board[j], board[i]
		}
	}
}

func zerored(board []int, x int) int {
	for ; x < len(board)-1; x++ {
		if board[x] != 0 {
			break
		}
	}
	return x
}

func main() {
	board := make([]int, 100, 100)
	for i := 0; i < len(board); i++ {
		board[i] = rand.Intn(2)
	}
	fmt.Println(board)
	zeroMover(board)
	fmt.Println(board)
}
