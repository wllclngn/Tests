package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func convSHA256(data []byte) []byte {
	hashish := sha256.Sum256(data)
	return hashish[:]
}

func main() {
	data, err := ioutil.ReadFile("[FILE]")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	hasher := convSHA256(data)
	fmt.Printf("RESULT: %x", hasher[:])
}
