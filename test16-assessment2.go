package main

import (
	"fmt"
	"strings"
)

func dig(x string) string {

	newString := []string{}
	completions := 0

	for i := 0; i < len(x); i++ {
		if i == len(x)-1 {
			newString = append(newString, "<>")
		} else {
			if string(x[i]) == "<" && string(x[i+1]) == ">" {
				newString = append(newString, "<>")
				completions++
				i++
			} else {
				newString = append(newString, "<>")
			}
		}
	}
	newString2 := strings.Join(newString, "")
	newLen := len(newString2)
	fmt.Println("COMPLETE #s:", completions, "LEN:", newLen, "\nNEW:"+newString2+"\n")
	return newString2
}

func main() {
	this := "><><<>>>>>"
	thisLen := len(this)
	fmt.Println("LEN:", thisLen, "\nORIGINAL:", this)
	dig(this)
	this2 := ">>>>>>>>"
	thisLen2 := len(this2)
	fmt.Println("LEN:", thisLen2, "\nORIGINAL:", this2)
	dig(this2)
	this3 := "<<<<<<<<"
	thisLen3 := len(this3)
	fmt.Println("LEN:", thisLen3, "\nORIGINAL:", this3)
	dig(this3)
	this4 := "<<<<<<<<<<<<<<<<<>><><>>><>><<<<<<<<>>><<<><>>><<"
	thisLen4 := len(this4)
	fmt.Println("LEN:", thisLen4, "\nORIGINAL:", this4)
	dig(this4)
}
