package main

import (
    "fmt"
    "math"
    "sync"
)

// GOHPERS RECYCLING C++ MANUALS, ala https://talks.golang.org/2012/waza.slide#1
// Thx to Chance Dinkins, Tim Heckman and Nathan Bass in Gophers' Slack channel.

func main() {

    wg := &sync.WaitGroup{}

    manuals := make(chan string)
    manuals2 := make(chan string)
    foundMan := make(chan string)
    foundMan2 := make(chan string)

    // GENERATE C++ MANUALS FOR RECYCLING
    wg.Add(6)
    go func() {
        defer wg.Done()
        defer close(manuals)
        for i := 0; i < 100; i++ {
            maths := math.Mod(float64(i), 2)
            if maths == 0 {
                manuals <- "C++ Manual"
            }
        }
    }()

    // GENERATE C++ MANUALS FOR RECYCLING
    go func() {
        defer wg.Done()
        defer close(manuals2)
        for i := 0; i < 100; i++ {
            maths := math.Mod(float64(i), 2)
            if maths == 0 {
                manuals2 <- "C++ Manual"
            }
        }
    }()

    // FIND C++ MANUALS FOR RECYCLING
    go func(x chan string) {
        defer wg.Done()
        defer close(foundMan)
        for j := 0; j < 50; j++ {
            foundMan := <-x
            if j < 15 {
                fmt.Println("GOT a " + foundMan + ". It's outdated!")
            } else if j < 25 {
                fmt.Println("GOT a " + foundMan + ". Yellow pages, and rats have ate on it.")
            } else {
                fmt.Println("GOT a " + foundMan + ". Almost need a break.")
            }
        }
    }(manuals)

    // FIND C++ MANUALS FOR RECYCLING
    go func(y chan string) {
        defer wg.Done()
        defer close(foundMan2)
        for k := 0; k < 50; k++ {
            foundMan2 := <-y
            if k < 15 {
                fmt.Println("FOUND a " + foundMan2 + "! Yah there, bud.")
            } else if k < 25 {
                fmt.Println("FOUND a " + foundMan2 + ". You betcha.")
            } else {
                fmt.Println("FOUND a " + foundMan2 + ". Workin' up a sweat!")
            }
        }
    }(manuals2)

    go func(fire chan string) {
        defer wg.Done()
        for l := 0; l < 50; l++ {
            select {
            case <-fire:
                fmt.Println("GOPHER #1 recycled a C++ manual.")
            }
        }
    }(foundMan)

    go func(fire2 chan string) {
        defer wg.Done()
        for l := 0; l < 50; l++ {
            select {
            case <-fire2:
                fmt.Println("GOPHER #2 recycled a C++ manual.")
            }
        }
    }(foundMan2)

    wg.Wait()
}
