// MERGE SORT
// returning to in future, return arr unnecessary due to memory address assignment

package main

import "fmt"

func mergeSort(x []int) []int {

	if len(x) < 2 {
		return x
	}
	mid := len(x) / 2
	return Merge(mergeSort(x[:mid]), mergeSort(x[mid:]))
}

func Merge(x, y []int) []int {

	girth, i, j := len(x)+len(y), 0, 0
	slice := make([]int, girth, girth)
	k := 0

	for i < len(x) && j < len(y) {
		if x[i] <= y[j] {
			slice[k] = x[i]
			k, i = k+1, i+1
		} else {
			slice[k] = y[j]
			k, j = k+1, j+1
		}
	}
	for i < len(x) {
		slice[k] = x[i]
		k, i = k+1, i+1
	}
	for j < len(y) {
		slice[k] = y[j]
		k, j = k+1, j+1
	}

	return slice
}

func main() {
	slice := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
		-2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
		88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}
	fmt.Println(slice)
	slice2 := mergeSort(slice)
	fmt.Println(slice2)
}
