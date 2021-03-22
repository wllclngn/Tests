// BIG END, LITTLE END BYTE CONVERSION

package main

import (
	"fmt"
)

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
		nibble := make([]byte, 2, 2)
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

}
