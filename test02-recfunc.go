package main

import "fmt"
import "time"

func recFunc(x int) int {
	if (x >= 0) {
		// fmt.Println(x)
		return recFunc(x - 1)
	}
	return x
}

func main() {
	start := time.Now()
	// x := recFunc(100)
	recFunc(100)
	elapsed := time.Since(start)
	// fmt.Println(x, ": Main function call.")
	fmt.Println(elapsed, ": is the execution time for recursive function.")
	start2 := time.Now()
	for i := 100; i >= 0; i-- {
		//fmt.Println(i)
	}
	elapsed2 := time.Since(start2)
	fmt.Println(elapsed2, ": is the execution time for iterative.")
}
