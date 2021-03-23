// BINARY TO HEX CONVERSION

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func binToHex(b []byte) string {

	binInt := 0
	power := 0

	for i := len(b) - 1; i >= 0; i-- {
		binInt += int(b[i]) * int(math.Pow(2, float64(power)))
		power++
	}

	fmt.Println("DEC:", binInt)
	var value []int
	value2 := []string{"0x"}
	if binInt > 9 {
		for i := 0; i < 9223372036854775807; i++ {
			if binInt >= 16 {
				y := binInt % 16
				value = append(value, y)
				binInt = binInt / 16
			} else {
				value = append(value, binInt)
				break
			}
		}
		for j := len(value) - 1; j >= 0; j-- {
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
		z := strconv.Itoa(binInt)
		value2 = append(value2, z)
	}
	value3 := strings.Join(value2, "")
	fmt.Println("HEX:", value3)
	return value3
}

func main() {

	start := time.Now()
	nibble := make([]byte, 1, 1)
	nibble = []byte{0, 1}
	fmt.Println("BIN:", nibble)

	binToHex(nibble)
	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed, "\n")

	start = time.Now()
	litBite := make([]byte, 1, 1)
	litBite = []byte{0, 0, 0, 1, 0, 0, 1, 1}
	fmt.Println("BIN:", litBite)

	binToHex(litBite)
	elapsed = time.Since(start)
	fmt.Println("Elapsed:", elapsed, "\n")

	start = time.Now()
	litBite2 := make([]byte, 2, 2)
	litBite2 = []byte{0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println("BIN:", litBite2)

	binToHex(litBite2)
	elapsed = time.Since(start)
	fmt.Println("Elapsed:", elapsed, "\n")

	start = time.Now()
	litBite3 := make([]byte, 2, 2)
	litBite3 = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println("BIN:", litBite3)

	binToHex(litBite3)
	elapsed = time.Since(start)
	fmt.Println("Elapsed:", elapsed, "\n")

}
