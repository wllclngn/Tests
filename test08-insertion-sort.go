// INSERTION SORT
// return arr unnecessary due to memory address assignment

package main

import "fmt"

func insertionSort(x int, arr []int) {
	for i := 0; i < x; i++ {
		j := i
		for j >= 0 && arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
}

func main() {
	puzzle1 := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
		-2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
		88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}
	fmt.Println(puzzle1)
	insertionSort(len(puzzle1)-1, puzzle1)
	fmt.Println(puzzle1)
}
