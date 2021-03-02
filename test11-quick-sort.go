// QUICK SORT

package main

import "fmt"

func partition(arr []int, x int, y int) int {
	i := (x - 1)
	pivot := arr[y]
	for j := x; j < y; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[y] = arr[y], arr[i+1]
	return (i + 1)
}

func quickSort(arr []int, x int, y int) []int {
	if len(arr) == 1 {
		return arr
	}
	if x < y {
		partInd := partition(arr, x, y)
		quickSort(arr, x, partInd-1)
		quickSort(arr, partInd+1, y)
	}
	return arr
}

func main() {
	puzzle1 := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
		-2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
		88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}
	fmt.Println(puzzle1)
	quickSort(puzzle1, 0, (len(puzzle1) - 1))
	fmt.Println(puzzle1)
}
