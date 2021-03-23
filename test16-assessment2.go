package main

import (
	"fmt"
	"strings"
)
func dig(x string) string {

	newString := []string{}

	for i := 0; i < len(x); i++ {
		if string(x[i]) == "<" {
			if i == len(x)-1 {
				newString = append(newString, "<")
				newString = append(newString, ">")
			} else {
				if string(x[i+1]) == ">" {
					newString = append(newString, "<")
					newString = append(newString, ">")
					i++
				} else {
					newString = append(newString, "<")
					newString = append(newString, ">")
				}
			}
		} else {
			newString = append(newString, "<")
			newString = append(newString, ">")
		}
	}
	newString2 := strings.Join(newString, "")
	newlen := len(newString2) / 2
	fmt.Println(newlen, " ", newString2)
	return newString2
}

func main() {
	/*
		this := "><><<>>>>>"
		fmt.Println(this)
		dig(this)
		this2 := ">>>>>>>>"
		fmt.Println(this2)
		dig(this2)
		this3 := "<<<<<<<<"
		fmt.Println(this3)
		dig(this3)
	*/
	this4 := "<<<<<<<<<<<<<<<<<>><><>>><>><<<<<<<<>>><<<><>>><<"
	lenthis4 := len(this4)
	fmt.Println(this4, " ", lenthis4)
	dig(this4)
}
