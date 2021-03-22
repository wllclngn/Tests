// BIG END, LITTLE END BYTE CONVERSION
// "Work in Progress"

package main

import (
	"fmt"
)

func bigLil() {

}

func lilBig(x []byte) []byte {

	var byteR []byte

	if (len(x) % 4) != 0 {
		fmt.Println("ERROR:", x)
		panic("Bit's length is not modulus of four.")
	}

	return byteR
}

func main() {
	/*
		nibble := make([]byte, 2, 2)
		nibble = []byte{0, 1}
		lilBig(nibble)
	*/

	bite := make([]byte, 8, 8)
	bite = []byte{0, 0, 0, 1, 0, 0, 1, 1}
	lilBig(bite)

	bigBite := make([]byte, 8, 8)
	bigBite = []byte{0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1}
	lilBig(bigBite)

	/*

		testLilBig2 := lilBig(bite)
		testLilBig3 := lilBig(bigBite)

		testBigLil := bigLil(testLilBig2)
	*/
}
