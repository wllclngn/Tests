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

func insertionSort(arr []int, x int, y int) {
	for i := x; i <= y; i++ {
		for j := i; j > x && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func mergeSort(arr []int, x int, y int, z int) {
	len1, len2 := y-x+1, z-y
	var arr2, arr3 []int
	for i := 0; i < len1; i++ {
		arr2 = append(arr2, arr[x+i])
	}
	for j := 0; j < len2; j++ {
		arr3 = append(arr3, arr[y+1+j])
	}

	a, b, c := 0, 0, x

	for a < len1 && b < len2 {
		if arr2[a] <= arr3[b] {
			arr[c] = arr2[a]
			a++
		} else {
			arr[c] = arr3[b]
			b++
		}
		c++
	}

	for a < len1 {
		arr[c] = arr2[a]
		c++
		a++
	}

	for b < len2 {
		arr[c] = arr3[b]
		c++
		b++
	}
}

func timSort(arr []int) {
	n := len(arr)
	minRun := calcMinRun(n)

	for i := 0; i < n; i += minRun {
		end := Min(i+minRun-1, n-1)
		insertionSort(arr, i, end)
	}

	for j := minRun; j < n; j <<= 1 {
		for left := 0; left < n; left += (j << 1) {
			mid := Min(n-1, left+j-1)
			right := Min((left + (j << 1) - 1), (n - 1))

			mergeSort(arr, left, mid, right)
		}
	}
}

func main() {
	slice := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
		-2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
		88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}
	fmt.Println(slice)
	start := time.Now()
	timSort(slice)
	// start2 := time.Now()
	// Resolves so fast time.Since() doesn't work
	elapsed := time.Since(start)
	fmt.Println(slice)
	fmt.Println("Start:", start)
	fmt.Println("Elapsed:", elapsed)
}
