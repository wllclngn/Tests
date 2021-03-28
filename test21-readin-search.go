// Import file, sort it and search over file contents...

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func binSearch(x string, y []string) int {
	high := len(y) - 1
	low := 0

	for high >= low {
		point := ((high - low) / 2) + low

		switch {
		case strings.EqualFold(x, y[point]):
			fmt.Println(point)
			return point
		case strings.Compare(x, y[point]) == -1:
			high = point - 1
		case strings.Compare(x, y[point]) == 1:
			low = point + 1
		}
	}
	return -1
}

func main() {
	data, err := ioutil.ReadFile("[FILE]")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	data_str := make([]string, len(data))
	k := 0
	for j := 0; j < len(data_str); j++ {
		if data[j] == 10 {
			k++
		} else {
			data_str[k] = data_str[k] + string(data[j])
		}
	}
	var data_str2 []string
	for l := 0; l < len(data_str); l++ {
		if data_str[l] != "" {
			data_str2 = append(data_str2, data_str[l])
		}
	}
	fmt.Println("INPUT DATA:", data_str2)
	searched := shellSort(data_str2)
	fmt.Println("SORTED SLICE LIBRARY:", searched)
	sought := "maRSupial"
	m := binSearch(sought, searched)
	if m >= 0 {
		fmt.Println("SEARCH:", sought, "\nINDEX RESULT:", m, "\nSLICE LIBRARY MATCH:", searched[m])
	} else {
		fmt.Println(sought, "was not found in the slice's library!")
	}
}
