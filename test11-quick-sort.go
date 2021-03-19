// QUICK SORT

package main

import "fmt"
import "time"

func partition(arr []int, x int, y int) int {
    i := (x - 1)
    fulcrum := arr[y]
    for j := x; j < y; j++ {
        if arr[j] <= fulcrum {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[y] = arr[y], arr[i+1]
    return (i + 1)
}

func quickSort(arr []int, x int, y int) {
    if len(arr) <= 1 {
        panic("The input slice is empty, or has a size of one.")
    } else if x < y {
        ledger := partition(arr, x, y)
        quickSort(arr, x, ledger-1)
        quickSort(arr, ledger+1, y)
    }
}

func main() {
    slice := []int{-14, -14, -13, -7, -4, -2, 0, 0, 5, 7, 7, 8, 12, 15, 15,
        -2, 7, 15, -14, 0, 15, 0, 7, -7, -4, -13, 5, 8, -14, 12, 49, 6, 78, 99,
        88, 48, 38, 29, 30, 133, 34, 52, 526, 664, 267, 377}
    fmt.Println(slice)
    start := time.Now()
    quickSort(slice, 0, (len(slice) - 1))
    elapsed := time.Since(start)
    fmt.Println(slice)
    fmt.Println("Start:", start)
    fmt.Println("Elapsed:", elapsed)
}
