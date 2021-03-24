// INSERTION SORT

package main

import "fmt"
import "time"

func insertionSort(x int, arr []int) {
    for i := 0; i < x; i++ {
        for j := i; j >= 0 && arr[j] > arr[j+1]; j-- {
            arr[j], arr[j+1] = arr[j+1], arr[j]
        }
    }
}

func main() {
    slice := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
        -2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
        88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}
    fmt.Println(slice)
    start := time.Now()
    insertionSort((len(slice) - 1), slice)
    elapsed := time.Since(start)
    fmt.Println(slice)
    fmt.Println("Start:", start)
    fmt.Println("Elapsed:", elapsed)
}
