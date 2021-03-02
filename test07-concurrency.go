package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	cave := [5]string{"nug1", "nug2", "nug3", "nug4", "nug5"}
	nugChan := make(chan string)
	nugChan2 := make(chan string)

	wg.Add(3)
	go func(mine [5]string) {
		defer wg.Done()
		defer close(nugChan)
		defer close(nugChan2)
		for _, item := range mine {
			nugChan <- item
			nugChan2 <- item
		}
	}(cave)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			foundNug := <-nugChan
			fmt.Println("GOT a nug: " + foundNug + ". Oh, it's a beaut!")
		}
	}()

	go func() {
		defer wg.Done()
		for j := 0; j < 5; j++ {
			foundNug2 := <-nugChan2
			fmt.Println("FOUND a nug: " + foundNug2 + ". Yah there, bud.")
		}
	}()

	// <-time.After(time.Second / 500)
	wg.Wait()
}
