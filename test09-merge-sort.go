// MERGE SORT
// return arr unnecessary due to memory address assignment

package main

import "fmt"

func mergeSort(arr []int) {
	if len(arr) > 1 {

		mid := len(arr) / 2

		L := arr[:mid]
		R := arr[mid:]

		mergeSort(L)
		mergeSort(R)

		for i := 0; i < len(L); i++ {
			if L[i] > R[i] {
				L[i], R[i] = R[i], L[i]
			}
		}
	}
}

func main() {
	puzzle1 := []int{12, 11, 13, 5, 7, 6}
	fmt.Println(puzzle1)
	mergeSort(puzzle1)
	fmt.Println(puzzle1)
}
