/*
!!! ONLY SINGLE RETURN CURRENTLY !!!
IMPORT, MANIPULATE AND SEARCH DATA
EXPONENTIAL, EXHAUSTIVE, MULTI-RETURN SEARCH UTILIZING BINARY SEARCH
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func readFile(x string) []string {
	data, err := ioutil.ReadFile(x)
	if err != nil {
		fmt.Println("File input ERROR:", err)
		ERROR := []string{"ERROR."}
		return ERROR
	}
	// CONVERT DATA TO []string, TRIM "\n" CHARACTERS
	var searched []string
	setup := ""
	for j := 0; j < len(data); j++ {
		switch {
		case data[j] == 10:
			searched = append(searched, setup)
			setup = ""
		case j == (len(data) - 1):
			setup = setup + string(data[j])
			searched = append(searched, setup)
			setup = ""
		default:
			setup = setup + string(data[j])
		}
	}
	return searched
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

func binSearch(x string, y []string, a int, z int) int {
	point := ((z - a) >> 1) + a

	for z >= a {

		switch {
		case strings.EqualFold(x, y[point]):
			return point
		case strings.Compare(x, y[point]) == -1:
			return binSearch(x, y, a, (point - 1))
		case strings.Compare(x, y[point]) == 1:
			return binSearch(x, y, (point + 1), z)
		}
	}
	return -1
}

func expoSearch(x string, y []string) int {
	if len(y) == 0 {
		return 0
	}

	i := 1
	for i < len(y) && y[i] <= x {
		i <<= 1
	}

	high := Min(i, len(y))
	low := (i >> 1)
	return binSearch(x, y, low, high)
}

func main() {
	searched := readFile("F://[4] PROGRAMMING//[02] GO/test23-tester.txt")
	//data_str := readFile("/run/media/EXTHD/DOCUMENTS [EXTHD]/tester.txt")
	searchedSort := shellSort(searched)
	sought := "newData"
	intSl := expoSearch(sought, searchedSort)
	if intSl != -1 {
		fmt.Printf("SEARCH: \"%v\"\nINDEX: %d\nSLICE LIBRARY MATCH: \"%v\"\n", sought, intSl, searched[intSl])
	} else {
		fmt.Printf("\"%v\" has no match the slice's library! %v is the result.", sought, intSl)
	}
}
