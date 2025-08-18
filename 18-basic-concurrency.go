/*
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
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Generates a slice of random unique nuggets
func generateNuggets(r *rand.Rand, maxNuggets int) []int {
	usedNumbers := make(map[int]bool) // Map to track used numbers
	numNuggets := []int{}
	for len(numNuggets) < maxNuggets {
		newNug := r.Intn(100)
		if !usedNumbers[newNug] { // Only add if the number hasn't been used
			numNuggets = append(numNuggets, newNug)
			usedNumbers[newNug] = true
		}
	}
	return numNuggets
}

// Populates channels with nuggets
func populateChannels(nuggets []int, nugChan, nugChan2 chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(nugChan)
	defer close(nugChan2)

	for _, nug := range nuggets {
		nugStr := fmt.Sprintf("nug%d", nug)
		nugChan <- nugStr
		nugChan2 <- nugStr
	}
}

// Selects and formats a random message for a nugget
func formatMessage(r *rand.Rand, messages []string, nug string) string {
	return fmt.Sprintf(messages[r.Intn(len(messages))], nug)
}

// Consumes nuggets from a channel
func consumeNuggets(nugChan chan string, messages []string, r *rand.Rand, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if foundNug, ok := <-nugChan; ok {
			fmt.Println(formatMessage(r, messages, foundNug))
		} else {
			fmt.Println("Channel closed unexpectedly.")
			break
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	messages := []string{
		"GOT a nug: %s. Oh, it's a beaut!",
		"FOUND a nug: %s. Yah there, bud.",
		"Look at this nug: %s. It's shining bright!",
		"Whoa, a nug: %s. What a treasure!",
		"Check this nug out: %s. It's amazing!",
		"Wow, a nug: %s. Pure gold!",
		"Spotted a nug: %s. Lucky find!",
		"Hey, it's a nug: %s. This one's special!",
		"Seen a nug: %s. It's remarkable!",
		"Got another nug: %s. Unbelievable!",
		"Found one more nug: %s. What a day!",
		"A nug appeared: %s. Incredible!",
		"Discovered a nug: %s. Such luck!",
		"Unearthed a nug: %s. Truly valuable!",
		"Just dug up a nug: %s. Fantastic!",
		"Picked up a nug: %s. Amazing find!",
		"Unearthing nuggets like this: %s. A gem!",
		"Lucky strike with this nug: %s. Outstanding!",
		"Treasure hunting success: %s. What a haul!",
		"Another shiny nug: %s. Marvelous!",
	}

	// Generate random nuggets
	numNuggets := generateNuggets(r, 100)

	// Create channels
	nugChan := make(chan string, len(numNuggets))
	nugChan2 := make(chan string, len(numNuggets))

	wg.Add(3)

	// Populate channels
	go populateChannels(numNuggets, nugChan, nugChan2, wg)

	// Consume nuggets from channels
	go consumeNuggets(nugChan, messages, r, wg)
	go consumeNuggets(nugChan2, messages, r, wg)

	wg.Wait()
}
