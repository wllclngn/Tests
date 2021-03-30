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
	data, err := ioutil.ReadFile("[FILE]")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	hash := sha256.Sum256(data)
	fmt.Printf("%x", hash[:])
}
