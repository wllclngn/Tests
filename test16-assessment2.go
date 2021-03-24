package main

import (
	"fmt"
	"strings"
)

func dig(x string) string {

	thisLen := len(x)
	fmt.Println("LEN:", thisLen, "\nORIGINAL:", x)
	newString := []string{}
	completions := 0

	for i := 0; i < len(x); i++ {
		if i == len(x)-1 {
			newString = append(newString, "<>")
			break
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
	if (newLen / (thisLen - completions)) == 2 {
		fmt.Println("All non-complete <>s have been filled out.")
	} else {
		panic("Please check dig's logic. All non-<>s must be completed.")
	}
	fmt.Println("COMPLETE #s:", completions, "LEN:", newLen, "\nNEW:"+newString2+"\n")
	return newString2
}

func main() {
	dig(">>>>>>>>")
	dig("<<<<<<<<")
	dig("><><<>>>>>")
	dig("<<<<<<<<<<<<<<<<<>><><>>><>><<<<<<<<>>><<<><>>><<")
}
