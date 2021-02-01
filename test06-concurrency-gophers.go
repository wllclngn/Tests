package main

import (
	"fmt"
	"math"
	"sync"
)

// GOHPERS BURNING C++ MANUALS THING
// Thx to Chance Dinkins, Tim Heckman and Nathan Bass in Gophers' Slack channel.

func main() {

	wg := &sync.WaitGroup{}

	// done := make(chan bool, 1)
	//manuals := make(chan []string, 43)
	manuals := make(chan string)
	manuals2 := make(chan string)
	foundMan := make(chan string)
	foundMan2 := make(chan string)

	// GENERATE C++ MANUALS FOR INCINERATOR
	wg.Add(5)
	go func() {
		defer wg.Done()
		defer close(manuals)
		for i := 0; i < 86; i++ {
			maths := math.Mod(float64(i), 2)
			if maths == 0 {
				manuals <- "C++ Manual"
			}
		}

	}()

	// GENERATE C++ MANUALS FOR INCINERATOR
	go func() {
		defer wg.Done()
		defer close(manuals2)
		for i := 0; i < 86; i++ {
			maths := math.Mod(float64(i), 2)
			if maths == 0 {
				manuals2 <- "C++ Manual"
			}
		}

	}()

	// FIND C++ MANUALS FOR INCINERATOR
	go func(x chan string) {
		defer wg.Done()
		defer close(foundMan)
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
	}(manuals)

	// FIND C++ MANUALS FOR INCINERATOR
	go func(y chan string) {
		defer wg.Done()
		defer close(foundMan2)
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

	}(manuals2)

	go func(fire chan string, fire2 chan string) {
		defer wg.Done()
		for l := 0; l <= 200; l++ {
			select {
			case <-fire:
				fmt.Println(<-fire, fire, "manuals have been burnt. #1")
				<-fire
			case <-fire2:
				fmt.Println(<-fire2, fire2, "manuals have been burnt. #2")
				<-fire2
			}
		}
	}(foundMan, foundMan2)

	wg.Wait()
}
