// BIG END, LITTLE END BYTE CONVERSION
// "Work in Progress"

package main

import (
	"fmt"
	"strconv"
)

func bigLil() {

}

func lilBig(x int) int {

	intConv := strconv.Itoa(x)
	if (len(intConv) % 4) != 0 {
		fmt.Println("ERROR: " + intConv)
		panic("Please check the length of your bits.")
	}

	var num int

	return num
}

func main() {
	//nibble := make([]byte, 2)
	nibble := 0x1101
	lilBig(nibble)

	/*
		bite := 00010011
		bigBite := 0001001101111111


		testLilBig2 := lilBig(bite)
		testLilBig3 := lilBig(bigBite)

		testBigLil := bigLil(testLilBig2)
	*/

}
