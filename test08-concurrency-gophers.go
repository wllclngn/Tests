package main

import (
	"fmt"
	"math"
	"sync"
)

// import "time"

// GOHPERS BURNING C++ MANUALS THING

func main() {

	var wg sync.WaitGroup

	// done := make(chan bool, 1)

	// var farm []string
	//manuals := make(chan []string, 43)
	manuals := make(chan string)
	manuals2 := make(chan string)
	foundMan := make(chan string)
	foundMan2 := make(chan string)

	// GENERATE C++ MANUALS FOR INCINERATOR
	wg.Add(5)
	go func() {
		for i := 0; i < 86; i++ {
			maths := math.Mod(float64(i), 2)
			if maths == 0 {
				manuals <- "C++ Manual"
			}
		}
		// close(manuals)
		//defer wg.Done()
	}()

	// GENERATE C++ MANUALS FOR INCINERATOR
	go func() {
		for i := 0; i < 86; i++ {
			maths := math.Mod(float64(i), 2)
			if maths == 0 {
				manuals2 <- "C++ Manual"
			}
		}
		// close(manuals)
		//defer wg.Done()
	}()

	// FIND C++ MANUALS FOR INCINERATOR
	go func(x chan string) {
		for j := 0; j < 43; j++ {
			foundMan := <-x
			if j < 14 {
				fmt.Println("GOT a " + foundMan + ". It's outdated!")
			} else if j < 28 {
				fmt.Println("GOT a " + foundMan + ". Yellow pages, and rats have ate on it.")
			} else {
				fmt.Println("GOT a " + foundMan + ". Almost need a break")
			}
		}
		// close(foundMan)
		//defer wg.Done()
	}(manuals)

	// FIND C++ MANUALS FOR INCINERATOR
	go func(y chan string) {
		for k := 0; k < 43; k++ {
			foundMan2 := <-y
			if k < 14 {
				fmt.Println("FOUND a " + foundMan2 + "! Yah there, bud.")
			} else if k < 28 {
				fmt.Println("FOUND a " + foundMan2 + ". You betcha.")
			} else {
				fmt.Println("FOUND a " + foundMan2 + ". Workin' up a sweat!")
			}
		}
		// close(foundMan2)
		//defer wg.Done()
	}(manuals2)

	go func(fire chan string, fire2 chan string) {
		for l := 0; l <= 86; l++ {
			select {
			case <-fire:
				fmt.Println("s have been burnt. #1")
			case <-fire2:
				fmt.Println("s have been burnt. #2")
			default:
				fmt.Println("Wrong.")
			}
		}
		//defer wg.Done()
	}(foundMan, foundMan2)

	wg.Wait()
}
