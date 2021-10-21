// GO VERSION OF https://just-the-punctuation.glitch.me/ //
// "For the love of Go..."

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func punctuation(x string) []string {
	var justPunct []string
	for i := 0; i < len(x); i++ {
		switch {
		case x[i] == 33:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 34:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 35:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 36:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 37:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 38:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 39:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 40:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 41:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 42:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 43:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 44:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 45:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 46:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 37:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 42:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 46:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 47:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 58:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 59:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 61:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 63:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 94:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 96:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 126:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 128:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 133:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 139:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 150:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 151:
			justPunct = append(justPunct, string("—"))
		case x[i] == 155:
			justPunct = append(justPunct, string(x[i]))
		case x[i] == 182:
			justPunct = append(justPunct, string(x[i]))
		}
	}
	return justPunct
}

func main() {

	data, err := ioutil.ReadFile("[PATH TO FILE]")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	data_str := string(data[:])
	data_str = strings.Replace(data_str, "\n", "", -1)
	//fmt.Println("INPUT DATA:", data_str)
	fmt.Println(punctuation(data_str))
}
