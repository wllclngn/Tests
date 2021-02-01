package main

import "fmt"
import "strconv"
import "strings"

// PREPEND EXAMPLE
// value2 = append([]string{"0x"}, value2...)

// CONVERT STRING TO INT W/ ONE RETURN VALUE
// y, _ := strconv.Atoi(x)

func dectohex(x int) {
    fmt.Println(x, "is the original decimal.")
    var value []int
    value2 := []string {"0x"}
    if x > 9 {
        for i := 0; i < 9223372036854775807; i++ {
            if x >= 16 {
                y := x % 16
                value = append(value, y)
                x = x / 16
            } else {
                value = append(value, x)
                break
            }
        }
        for j := len(value)-1; j >= 0; j-- {
            if value[j] == 10 {
                value2 = append(value2, "A")
            } else if value[j] == 11 {
                value2 = append(value2, "B")
            } else if value[j] == 12 {
                value2 = append(value2, "C")
            } else if value[j] == 13 {
                value2 = append(value2, "D")
            } else if value[j] == 14 {
                value2 = append(value2, "E")
            } else if value[j] == 15 {
                value2 = append(value2, "F")
            } else {
                valueStr := strconv.Itoa(value[j])
                value2 = append(value2, valueStr)
            }
        }
    } else {
        z := strconv.Itoa(x)
        value2 = append(value2, z)
    }
    value3 := strings.Join(value2, "")
    fmt.Println(value3, "is the HEX string.")
    return
}

func main() {
    dectohex(117)
    dectohex(352)
    dectohex(23000)
    dectohex(12384800984930)
}
