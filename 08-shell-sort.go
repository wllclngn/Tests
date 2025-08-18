package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shellSort(x []int) []int {
	n, k, gaps := len(x), 1, []int{1}

	for {
		a, b, c := k, 2, 1
		for a > 0 {
			if a&1 != 0 {
				c *= b
			}
			a >>= 1
			b *= b
		}
		gap := c + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			for j := i; j > 0; j -= gap {
				if x[j-gap] > x[j] {
					x[j-gap], x[j] = x[j], x[j-gap]
				}
			}
		}
	}
	return x
}

func main() {

	slice := make([]int, 100, 100)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(50) - rand.Intn(50)
	}

	start := time.Now()
	shellSort(slice)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
	//fmt.Println(slicer)

}
