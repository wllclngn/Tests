package main

import "fmt"

func ack(x int, y int) int {
	// fmt.Println(x, y, "are the current values.")
	if x == 0 {
		return y + 1
	}
	if y == 0 {
		return ack(x-1, 1)
	}
	return ack(x-1, ack(x, y-1))
}

func main() {
	ack(6, 6)
}
