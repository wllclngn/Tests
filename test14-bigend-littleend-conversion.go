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
	//var nibble []byte
	nibble := make([]byte, 2, 2)
	nibble = []byte{0, 1}
	lilBig(nibble)

	/*
		var bite []byte
		bite = make([]byte, 8, 8)
		bite == []byte{0, 0, 0, 1, 0, 0, 1, 1}
		lilBig(bite)

		var bite []byte
		bite = make([]byte, 8, 8)
		bite == []byte{0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1}
		lilBig(bite)

		testLilBig2 := lilBig(bite)
		testLilBig3 := lilBig(bigBite)

		testBigLil := bigLil(testLilBig2)
	*/
}
