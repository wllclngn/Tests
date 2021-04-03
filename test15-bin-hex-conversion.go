// BIG END, LITTLE END BIT CONVERSION

package main

import (
	"fmt"
)

func bigLilConv(x []byte) []byte {

	if (len(x) % 4) != 0 {
		fmt.Println("ERROR:", x)
		panic("Bits' length is not modulus of four.")
	} else if (len(x) % 8) != 0 {
		fmt.Println("ERROR:", x)
		panic("Bits' length is not 8-bit based.")
	}

	for i := 0; i < (len(x) >> 1); i += 4 {
		if (i % 4) == 0 {

			// x[i+0], x[i+1], x[i+2], x[i+3], x[len(x)-(i+4)], x[len(x)-(i+3)], x[len(x)-(i+2)], x[len(x)-(i+1)] = x[len(x)-(i+4)], x[len(x)-(i+3)], x[len(x)-(i+2)], x[len(x)-(i+1)], x[i+0], x[i+1], x[i+2], x[i+3]

			x[i+0], x[len(x)-(i+4)] = x[len(x)-(i+4)], x[i+0]
			x[i+1], x[len(x)-(i+3)] = x[len(x)-(i+3)], x[i+1]
			x[i+2], x[len(x)-(i+2)] = x[len(x)-(i+2)], x[i+2]
			x[i+3], x[len(x)-(i+1)] = x[len(x)-(i+1)], x[i+3]

		}
	}

	return x
}

func main() {

	/*
		var nibble []byte
		nibble := make([]byte, 1, 1)
		nibble = []byte{0, 1}
		bigLilConv(nibble)
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

	biggerBite := make([]byte, 2, 2)
	biggerBite = []byte{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(biggerBite)
	test4 := bigLilConv(biggerBite)
	fmt.Println(test4)
	test5 := bigLilConv(test4)
	fmt.Println(test5)

}
