// INSERTION SORT
package main

import "fmt"

func insertionSort(n int, arr []int) {

	for i := 0; i < n; i++ {
		j := i
		for j >= 0 && arr[j] > arr[j+1] {
			s := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = s
			j--
		}
		fmt.Println(arr)
	}

}

func main() {
	numb := 5
	puzzle1 := []int{5, 6, 2, 4, 3, 1}
	insertionSort(numb, puzzle1)
}
