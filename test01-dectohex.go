package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func dectohex(x int) {
	fmt.Println("DEC:", x)
	var value []int
	value2 := []string{"0x"}
	if x > 9 {
		for i := 0; i < 9223372036854775807; i++ {
		    if x >= 16 {
				y := x % 16
				value = append(value, y)
				x >>= 3
		   	} else {
				value = append(value, x)
				break
			}
		}
		for j := len(value) - 1; j >= 0; j-- {
		    switch {
			case value[j] == 10:
					    value2 = append(value2, "A")
			case value[j] == 11:
					    value2 = append(value2, "B")
			case value[j] == 12:
					    value2 = append(value2, "C")
			case value[j] == 13:
					    value2 = append(value2, "D")
			case value[j] == 14:
					    value2 = append(value2, "E")
			case value[j] == 15:
					    value2 = append(value2, "F")
			default:
			    valueStr := strconv.Itoa(value[j])
			    value2 = append(value2, valueStr)
		    }
		}
	} else {
		z := strconv.Itoa(x)
		value2 = append(value2, z)
	}
	value3 := strings.Join(value2, "")
	fmt.Println("HEX:", value3)
	return
}

func main() {
	start := time.Now()
	dectohex(23000)
	dectohex(12384800984930)
	dectohex(9223372036854775807)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
}
