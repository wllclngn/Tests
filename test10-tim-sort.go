// TIM SORT FOR GOLANG
// THE INSERT-MERGE SORT BY TIM PETERS

package main

import (
	"fmt"
	"time"
)

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func calcMinRun(n int) int {
	MIN_MERGE := 32

	r := 0
	for n >= MIN_MERGE {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func insertionSort(arr []int, left int, right int) {
	for i := left; i <= right; i++ {
		j := i
		for j > left && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}

func merge(arr []int, l int, m int, r int) {
	len1, len2 := m-l+1, r-m
	var left []int
	var right []int
	for i := 0; i < len1; i++ {
		left = append(left, arr[l+i])
	}
	for i := 0; i < len2; i++ {
		right = append(right, arr[m+1+i])
	}

	i, j, k := 0, 0, l

	for i < len1 && j < len2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len1 {
		arr[k] = left[i]
		k++
		i++
	}

	for j < len2 {
		arr[k] = right[j]
		k++
		j++
	}
}

func timSort(arr []int) {
	n := len(arr)
	minRun := calcMinRun(n)

	// RUN
	for start := 0; start < n; start += minRun {
		end := Min(start+minRun-1, n-1)
		insertionSort(arr, start, end)
	}

	size := minRun
	for size < n {
		for left := 0; left < n; left += (2 * size) {

			mid := Min(n-1, left+size-1)
			right := Min((left + (2 * size) - 1), (n - 1))

			merge(arr, left, mid, right)
		}
		size = 2 * size
	}

}

func main() {
	puzzle1 := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
		-2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
		88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}
	fmt.Println(puzzle1)
	start := time.Now()
	timSort(puzzle1)
	start2 := time.Now()
	// Resolves so fast time.Since() doesn't work
	// elapsed := time.Since(start)
	fmt.Println(puzzle1)
	fmt.Println("Start:", start)
	fmt.Println("Elapsed:", start2)
}
