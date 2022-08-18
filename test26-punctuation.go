// GO VERSION OF https://just-the-punctuation.glitch.me/
// "For the love of Go—"

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func punctuation(x string) []string {

	x = strings.Replace(x, "\r", "", -1)
	x = strings.Replace(x, "\n", "", -1)
	x = strings.Replace(x, " ", "", -1)

	var justPunct []string
	testing := ""
	noQuote := ""

	for i := 0; i < len(x); i++ {

		testing = strconv.QuoteRuneToASCII(rune(x[i]))

		if len(testing) == 8 {
			noQuote = testing[2 : len(testing)-1]
			switch {
			case noQuote == "u0094":
				justPunct = append(justPunct, "—")
			case noQuote == "u0099":
				justPunct = append(justPunct, "’")
			case noQuote == "u009c":
				justPunct = append(justPunct, "“")
			case noQuote == "u009d":
				justPunct = append(justPunct, "”")
			}
		}

		switch {
		case x[i] == 33:
			justPunct = append(justPunct, "!")
		case x[i] == 34:
			justPunct = append(justPunct, "\"")
		case x[i] == 35:
			justPunct = append(justPunct, "#")
		case x[i] == 36:
			justPunct = append(justPunct, "$")
		case x[i] == 37:
			justPunct = append(justPunct, "%")
		case x[i] == 38:
			justPunct = append(justPunct, "&")
		case x[i] == 39:
			justPunct = append(justPunct, "'")
		case x[i] == 40:
			justPunct = append(justPunct, "(")
		case x[i] == 41:
			justPunct = append(justPunct, ")")
		case x[i] == 42:
			justPunct = append(justPunct, "*")
		case x[i] == 43:
			justPunct = append(justPunct, "+")
		case x[i] == 44:
			justPunct = append(justPunct, ",")
		case x[i] == 45:
			justPunct = append(justPunct, "-")
		case x[i] == 46:
			justPunct = append(justPunct, ".")
		case x[i] == 47:
			justPunct = append(justPunct, "/")
		case x[i] == 58:
			justPunct = append(justPunct, ":")
		case x[i] == 59:
			justPunct = append(justPunct, ";")
		case x[i] == 61:
			justPunct = append(justPunct, "=")
		case x[i] == 63:
			justPunct = append(justPunct, "?")
		case x[i] == 96:
			justPunct = append(justPunct, "`")
		case x[i] == 126:
			justPunct = append(justPunct, "~")
		}
	}
	return justPunct
}

func main() {

	data, err := ioutil.ReadFile("./The Fall of the House of Usher.txt")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	data_str := string(data[:])
	fmt.Println(punctuation(data_str))
}
