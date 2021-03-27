// BINARY SEARCH, ALPHABETIC

package main

import "fmt"
import "strings"

func binSearch(x string, y []string) int {
    low := 0
    high := len(y) - 1

    for low <= high {
        point := low + (high-low)/2

        switch {
            case strings.EqualFold(x, y[point]):
                return point
            case strings.Compare(x, y[point]) == -1:
                high = point - 1
            default:
                low = point + 1
        }
    }
    return -1
}

func main() {
    var searched []string = []string{
        "aardvark",
        "anteater",
        "chimpanzee",
        "hippo",
        "marsupial",
        "orangutan",
        "rhino",
    }
    fmt.Println("Sorted slice:", searched)
    sought := "aardvark"
    i := binSearch(sought, searched)
    switch {
        case i < 0:
            fmt.Println("The word", sought, "could not be found!")
        default:
            fmt.Println("The word", sought, "was found at index:", i, searched[i])
    }
}
