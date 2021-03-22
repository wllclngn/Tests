// BIG END, LITTLE END BYTE CONVERSION
// "Work in Progress"

package main

import (
	"fmt"
)

func bigLil() {

}

func lilBig(x []byte) int {

	if (len(x) % 4) != 0 {
		fmt.Println("ERROR: ", x)
		panic("Please check the length of your bits.")
	}

	var num int

	return num
}

func main() {
	var nibble []byte
	nibble = make([]byte, 2, 2)
	// s == []byte{0, 0, 0, 0, 0}
	//nibble := make([]byte, 2)
	lilBig(nibble)

	/*
		bite := 00010011
		bigBite := 0001001101111111


		testLilBig2 := lilBig(bite)
		testLilBig3 := lilBig(bigBite)

		testBigLil := bigLil(testLilBig2)
	*/

}
