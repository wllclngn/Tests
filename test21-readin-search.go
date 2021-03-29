// IMPORT, MANIPULATE AND SEARCH DATA
// BROKEN FOR NOW

package main

import (
	"fmt"
	"io/ioutil"
)

// GLOBAL CHANNEL
var intSl = make(chan int)

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func shellSort(x []string) []string {
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

func binSearch(x string, y []string, l int, r int) chan int {

	if r >= l {
		mid := l + (r-l)/2

		if y[mid] == x {
			return mid
		}

		if y[mid] > x {
			intSl <- binSearch(x, y, l, mid-1)
			fmt.Printf("%v is the found int", intSl)
			return intSl
		}

		intSl <- binSearch(x, y, mid+1, r)
		fmt.Printf("%v is the found int", intSl)
		return intSl
	}

	return -1
}

func expoSearch(x string, y []string) int {

	n := len(y) / len(y[0])
	if y[0] == x {
		return 0
	}

	i := 1
	for i < n && y[i] <= x {
		i = i * 2
	}

	intSl <- binSearch(x, y, (i / 2), Min(i, len(y)))
	fmt.Printf("%v is the found int", intSl)
	return -1
}

func main() {
	data, err := ioutil.ReadFile("d://DOCUMENTS [EXTHD]/tester.txt")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	// CONVERT DATA TO []string, TRIM "\n" CHARACTERS
	var data_str []string
	setup := ""
	for j := 0; j < len(data); j++ {
		if data[j] == 10 {
			data_str = append(data_str, setup)
			setup = ""
		} else {
			setup = setup + string(data[j])
		}
	}
	searched := shellSort(data_str)
	sought := "marsupial"
	m := expoSearch(sought, searched)
	if m != -1 {
		fmt.Printf("SEARCH: \"%v\"\nINDEX: %d\nSLICE LIBRARY MATCH: \"%v\"\n", sought, intSl, searched[m])
	} else {
		fmt.Printf("\"%v\" has no match the slice's library!", sought)
	}
}
