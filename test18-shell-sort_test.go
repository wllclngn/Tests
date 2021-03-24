// SHELL SORT

package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	slice := make([]int, 100, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}

	shellSort(slice)

	//slice[1] = 1000

	for i := 0; i < len(slice); i++ {
		if i < len(slice)-1 {
			if slice[i] > slice[i+1] {
				t.Log("ERROR: Solution is incorrect. Please check shellSort()'s accuracy.")
				t.FailNow()
			}
		}
	}

}

func shellSort(x []int) {
	n, gaps, k := len(x), []int{1}, 1

	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if x[j-gap] > x[j] {
					x[j-gap], x[j] = x[j], x[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}
