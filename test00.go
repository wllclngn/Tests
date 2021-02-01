package main

import "fmt"
import "strconv"

func main() {
    var value []string
    // value = append(value, "0x")
    fmt.Println(value, "is the original string.")
    // FIRST PASS
    for i := 0; i < 16; i++ {
        y := strconv.Itoa(i) + "BOOTY"
        value = append(value, y)
    }
    fmt.Println(value, "is the first pass' string.")
    for j := 0; j < len(value); j++ {
        fmt.Println(value[j], "is index", j, "'s value on the first pass' string.")
    }
    // SECOND PASS
    for k := 0; k < 16; k++ {
        value[k] = strconv.Itoa(k) + "CLAP"
    }
    for l := 0; l < len(value); l++ {
        fmt.Println(value[l], "is index", l, "'s value on the second pass' string.")
    }
}
