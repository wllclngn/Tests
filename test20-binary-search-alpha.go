// BINARY SEARCH, ALPHABETIC, IRREGARDLESS OF CASING

package main

import "fmt"
import "strings"

func binSearch(x string, y []string) int {
    high := len(y) - 1
    low := 0

    for high >= low {
        point := ((high-low)/2) + low

        switch {
            case strings.EqualFold(x, y[point]):
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
    var searched []string = []string{
        "aardvark",
        "anteater",
        "chimpanzee",
        "hippo",
        "marsupial",
        "orangutan",
        "rhino",
    }
    fmt.Println("SORTED SLICE LIBRARY:", searched)
    sought := "HIPPO"
    i := binSearch(sought, searched)
    if i > 0 {
        fmt.Println("SEARCH:", sought, "\nINDEX RESULT:", i, "\nSLICE LIBRARY MATCH:", searched[i])
    } else {
        fmt.Println(sought, "was not found in the slice's library!")
    }
}
