package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func convSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func main() {
	data, err := ioutil.ReadFile("d://DOCUMENTS [EXTHD]/tester.txt")
	// data, err := ioutil.ReadFile("/run/media/EXTHD/DOCUMENTS [EXTHD]/tester.txt")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	hasher := convSHA256(data)
	fmt.Printf("RESULT: %x", hasher[:])
}
