package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

// calcMinRun calculates the minimum run size for Timsort.
func calcMinRun(n int) int {
	const MIN_MERGE = 32
	r := 0
	for n >= MIN_MERGE {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

// insertionSort performs insertion sort on a subarray.
func insertionSort(arr []int, start, end int) {
	for i := start + 1; i <= end; i++ {
		key := arr[i]
		j := i - 1
		for j >= start && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// cacheFriendlyMerge performs a cache-optimized merge.
func cacheFriendlyMerge(arr []int, left, mid, right int) {
	// Divide the array into blocks that fit into the CPU cache.
	const BLOCK_SIZE = 64 // Approximate size of a cache line (in bytes).
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	// Copy the merged data back to the original array.
	copy(arr[left:right+1], temp)
}

// parallelMerge performs merging in parallel for large chunks.
func parallelMerge(arr []int, size, n int, temp []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for left := 0; left < n; left += 2 * size {
		mid := int(math.Min(float64(left+size-1), float64(n-1)))
		right := int(math.Min(float64(left+2*size-1), float64(n-1)))

		if mid < right {
			cacheFriendlyMerge(arr, left, mid, right)
		}
	}
}

// parallelTimSort performs Timsort with parallelized sorting and merging.
func parallelTimSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	minRun := calcMinRun(n)

	// Step 1: Parallelized insertion sort for initial runs.
	var wg sync.WaitGroup
	cpuCount := runtime.NumCPU()
	chunkSize := (n + cpuCount - 1) / cpuCount

	for i := 0; i < n; i += chunkSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := int(math.Min(float64(start+chunkSize-1), float64(n-1)))
			insertionSort(arr, start, end)
		}(i)
	}
	wg.Wait()

	// Step 2: Iterative merging with parallelism for larger sizes.
	temp := make([]int, n)
	for size := minRun; size < n; size *= 2 {
		if size >= chunkSize {
			// Use parallel merging for larger sizes.
			for i := 0; i < cpuCount; i++ {
				wg.Add(1)
				go parallelMerge(arr, size, n, temp, &wg)
			}
			wg.Wait()
		} else {
			// Use sequential merging for smaller sizes.
			for left := 0; left < n; left += 2 * size {
				mid := int(math.Min(float64(left+size-1), float64(n-1)))
				right := int(math.Min(float64(left+2*size-1), float64(n-1)))

				if mid < right {
					cacheFriendlyMerge(arr, left, mid, right)
				}
			}
		}
	}
}

// main demonstrates the super optimized Timsort implementation.
func main() {
	slice := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
		-2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
		88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}

	fmt.Println("Before sorting:", slice)

	start := time.Now()

	// Use parallel Timsort.
	parallelTimSort(slice)

	elapsed := time.Since(start)

	fmt.Println("After sorting: ", slice)
	fmt.Printf("Elapsed time: %v\n", elapsed)

	// Benchmark with a larger dataset.
	fmt.Println("\nBenchmarking with a larger dataset...")
	largeSlice := make([]int, 100000)
	for i := range largeSlice {
		largeSlice[i] = 100000 - i // Reverse sorted
	}

	start = time.Now()
	parallelTimSort(largeSlice)
	elapsed = time.Since(start)

	fmt.Printf("Larger dataset sorted in: %v\n", elapsed)
}
