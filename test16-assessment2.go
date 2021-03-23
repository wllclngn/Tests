package main

import (
	"fmt"
	"strings"
)

func dig(x string) string {

	newString := []string{}

	for i := 0; i < len(x); i++ {
		if string(x[i]) == "<" {
			if i < len(x)-1 {
				if string(x[i+1]) == "<" {
					newString = append(newString, "<")
					newString = append(newString, ">")
				} else {
					newString = append(newString, "<")
					newString = append(newString, ">")
					i++
				}
			}
		} else {
			if i < len(x)-1 {
				if string(x[i+1]) == ">" {
					newString = append(newString, "<")
					newString = append(newString, ">")
				} else {
					newString = append(newString, "<")
					newString = append(newString, ">")
				}
			}
		}
	}
	value3 := strings.Join(newString, "")
	fmt.Println(value3)
	return value3
}

func main() {
	this := "><><<>>>>>"
	fmt.Println(this)
	dig(this)
	this2 := ">>>>>>>>"
	fmt.Println(this2)
	dig(this2)
	this3 := "<<<<<<<<"
	fmt.Println(this3)
	dig(this3)
}
