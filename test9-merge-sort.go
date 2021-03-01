// MERGE SORT
// return arr unnecessary due to memory address assignment

package main

import "fmt"

func mergeSort(arr []int) {
	if len(arr) > 1 {

		mid := len(arr) / 2

		L := arr[mid:]
		R := arr[:mid]

		mergeSort(L)
		mergeSort(R)

		// i, j, k := 0, 0, 0

		for i := 0; i < len(R); i++ {
			if L[i] < R[i] {
				temp := L[i]
				L[i] = R[i]
				R[i] = temp
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
