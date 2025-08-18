// PRODUCT ARRAY/SLICE

package main

import "fmt"

func subArr(x []int) int {
	cur, neg, max := 1, 1, x[0]

	for i := 0; i < len(x); i++ {

		switch {
		case x[i] > 0:
			cur, neg = x[i]*cur, x[i]*neg
		case x[i] < 0:
			cur, neg = x[i]*neg, x[i]*cur
		default:
			cur, neg = 0, 1
		}

		if max < cur {
			max = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return max
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := subArr(slice)
	fmt.Println(slice2)
}
