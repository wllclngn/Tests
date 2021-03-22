// LITTLE END BITS TO HEX CONVERSION

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func litEndToHex(b []byte) string {

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

	intNum := 0

	for i := 0; i < len(b); i++ {
		intByte := int(b[(len(b) - (i + 1))])
		intNum = intNum + (intByte * int(math.Pow(2, float64(i))))
		fmt.Println(intNum)
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

func bigLilConv(x []byte) []byte {

	// var bitter []byte

	if (len(x) % 4) != 0 {
		fmt.Println("ERROR:", x)
		panic("Bit's length is not modulus of four.")
	}

	for i := 0; i < len(x)/2; i++ {
		if (i % 4) == 0 {
			x[i], x[i+1], x[i+2], x[i+3], x[len(x)-(4+i)], x[len(x)-(3+i)], x[len(x)-(2+i)], x[len(x)-(1+i)] = x[len(x)-(4+i)], x[len(x)-(3+i)], x[len(x)-(2+i)], x[len(x)-(1+i)], x[i], x[i+1], x[i+2], x[i+3]
		}
	}

	return x
}

func main() {

	/*
		nibble := make([]byte, 1, 1)
		nibble = []byte{0, 1}
		lilBig(nibble)
	*/

	bite := make([]byte, 1, 1)
	bite = []byte{0, 0, 0, 1, 0, 0, 1, 1}
	fmt.Println(bite)
	test1 := bigLilConv(bite)
	fmt.Println(test1)

	bigBite := make([]byte, 2, 2)
	bigBite = []byte{0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(bigBite)
	test2 := bigLilConv(bigBite)
	fmt.Println(test2)
	test3 := bigLilConv(test2)
	fmt.Println(test3)

	litEndToHex(test3)

}
