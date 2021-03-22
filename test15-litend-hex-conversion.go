// JANKY LITTLE END BITS TO HEX CONVERSION

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

/*

	uintByte := binary.BigEndian.Uint16()
	uintByte := binary.BigEndian.Uint32()
	uintByte := binary.BigEndian.Uint64()
	uintByte := binary.LittleEndian.Uint64()
	strInt := strconv.Itoa(intByte)

	ui, err := strconv.ParseUint(strInt, 2, 16)
	if err != nil {
		fmt.Println(err)
		return "Pretty much."
	}

	return fmt.Sprintf("%x", ui)
*/

func lilEndToHex(b []byte) string {

	intNum := 0

	for i := 0; i < len(b); i++ {
		intByte := int(b[(len(b) - (i + 1))])
		intNum = intNum + (intByte * int(math.Pow(2, float64(i))))
	}

	fmt.Println(intNum, "is the original decimal.")
	var value []int
	value2 := []string{"0x"}
	if intNum > 9 {
		for i := 0; i < 9223372036854775807; i++ {
			if intNum >= 16 {
				y := intNum % 16
				value = append(value, y)
				intNum = intNum / 16
			} else {
				value = append(value, intNum)
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
		z := strconv.Itoa(intNum)
		value2 = append(value2, z)
	}
	value3 := strings.Join(value2, "")
	fmt.Println(value3, "is the HEX string.")
	return value3
}

func main() {

	/*
		nibble := make([]byte, 1, 1)
		nibble = []byte{0, 1}
		lilBig(nibble)
	*/

	start := time.Now()
	litBite := make([]byte, 1, 1)
	litBite = []byte{0, 0, 0, 1, 0, 0, 1, 1}
	fmt.Println(litBite, "is the original Little End Binary string.")

	lilEndToHex(litBite)
	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)

	start = time.Now()
	litBite2 := make([]byte, 2, 2)
	litBite2 = []byte{0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(litBite2, "is the original Little End Binary string.")

	lilEndToHex(litBite2)
	elapsed = time.Since(start)
	fmt.Println("Elapsed:", elapsed)

}
