// PRODUCT ARRAY/SLICE EXCEPT SELF

package main

import "fmt"

func prodArr(x []int) []int {

	res := make([]int, len(x))

	product := 1
	for i := 0; i < len(x); i++ {
		res[i] = product
		product *= x[i]
	}

	product = 1
	for i := len(x) - 1; i >= 0; i-- {
		res[i] *= product
		product *= x[i]
	}

	return res
}

func main() {
	slice := []int{2, 4, 6, 8, 10}
	slice2 := prodArr(slice)
	fmt.Println(slice2)
}
