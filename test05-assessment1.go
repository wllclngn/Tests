package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solution(A int, B int) int {
	x := strconv.Itoa(A)
	y := strconv.Itoa(B)
	for i := 0; i < len(y); i++ {
		if x[0] == y[i] {
			if i+1 < len(y) {
				if x[1] == y[i+1] {
					// var z [1000000]string
					z := make([]string, len(y))
					for j := range z {
						z[j] = "0"
					}
					for k := i + 2; k < len(y); k++ {
						z[k] = string(y[k])
					}
					zJoined := strings.Join(z, "")
					// C, _ := strconv.Atoi(zJoined)
					extracurricular(x, zJoined, i)
					fmt.Println(i)
					return i
				}
			}
		}
	}
	fmt.Println(-1)
	return -1
}

func extracurricular(C string, D string, index int) int {
	for l := index; l < len(D); l++ {
		if C[0] == D[l] {
			if l+1 < len(D) {
				if C[1] == D[l+1] {
					z := make([]string, len(D))
					for m := range z {
						z[m] = "0"
					}
					for n := l + 2; n < len(D); n++ {
						z[n] = string(D[n])
					}
					zJoined := strings.Join(z, "")
					extracurricular(C, zJoined, l)
					fmt.Println(l)
					return l
				}
			}
		}
	}
	return 0
}

func main() {
	start := time.Now()
	solution(53, 15353535555)
	solution(54, 15353535555)
	solution(57, 15353575557)
	duration := time.Since(start)
	fmt.Println(duration)
}
