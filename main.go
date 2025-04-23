package main

import (
	"fmt"
	"sync"
	"time"
)

// useOven demonstrates the use of a semaphore (ovenSlots) to limit concurrent oven usage
func useOven(dish string, ovenSlots chan struct{}) {
	ovenSlots <- struct{}{} // acquire (semaphore pattern)
	fmt.Println("Using oven for", dish)
	time.Sleep(2 * time.Second) // simulate oven time
	<-ovenSlots // release (semaphore pattern)
	fmt.Println("Done with", dish)
}

// bakeLemonCake demonstrates sending data through a channel and using a WaitGroup
func bakeLemonCake(sugarLevelChan chan<- int, wg *sync.WaitGroup, ovenSlots chan struct{}) {
	defer wg.Done()
	fmt.Println("Baking lemon cake: Deciding sugar level...")
	time.Sleep(1 * time.Second) // Simulate time to decide
	sugarLevel := 5
	fmt.Printf("Baking lemon cake: Sugar level decided: %d\n", sugarLevel)
	sugarLevelChan <- sugarLevel // Send sugar level to channel
	useOven("lemon cake", ovenSlots)
	fmt.Println("Baking lemon cake: Done!")
}

// bakeStrawberryCupcakes demonstrates receiving from a channel and using a WaitGroup
func bakeStrawberryCupcakes(sugarLevelChan <-chan int, wg *sync.WaitGroup, ovenSlots chan struct{}) {
	defer wg.Done()
	fmt.Println("Baking strawberry cupcakes: Waiting for sugar level from lemon cake...")
	sugarLevel := <-sugarLevelChan // Wait for sugar level
	fmt.Printf("Baking strawberry cupcakes: Received sugar level: %d\n", sugarLevel)
	useOven("strawberry cupcakes", ovenSlots)
	fmt.Println("Baking strawberry cupcakes: Done!")
}

// grillChicken demonstrates using a WaitGroup, semaphore, and sending a result via a channel
func grillChicken(wg *sync.WaitGroup, ovenSlots chan struct{}, done chan<- string) {
	defer wg.Done()
	useOven("grilled chicken", ovenSlots)
	done <- "Chicken"
	fmt.Println("Grilling chicken: Done!")
}

// cookGoatStew demonstrates using a WaitGroup, semaphore, and sending a result via a channel
func cookGoatStew(wg *sync.WaitGroup, ovenSlots chan struct{}, done chan<- string) {
	defer wg.Done()
	useOven("goat stew", ovenSlots)
	done <- "Stew"
	fmt.Println("Cooking goat stew: Done!")
}

func main() {
	var wg sync.WaitGroup

	// --- Buffered Channel Example ---
	wg.Add(1)
	go func() {
		defer wg.Done()
		bufferedChan := make(chan int, 2) // buffered channel
		fmt.Println("Demonstrating buffered channel behavior:")
		fmt.Println("Sending 1 to bufferedChan...")
		bufferedChan <- 1 // does not block
		fmt.Println("Sent 1 (no block)")
		fmt.Println("Sending 2 to bufferedChan...")
		bufferedChan <- 2 // does not block
		fmt.Println("Sent 2 (no block)")
		// The next send would block because buffer is full
		// fmt.Println("Sending 3 to bufferedChan (will block)...")
		// bufferedChan <- 3 // would block here
		fmt.Println("Receiving from bufferedChan:")
		fmt.Println(<-bufferedChan)
		fmt.Println(<-bufferedChan)
		fmt.Println("Buffered channel demonstration complete.")
	}()

	// --- Channels and Semaphores Example ---
	// sugarLevelChan: buffered channel for sugar level communication
	// ovenSlots: semaphore channel for oven slot management (concurrency control)
	// chickenDone, stewDone: channels for completion signals
	
	sugarLevelChan := make(chan int, 2) // Now buffered
	ovenSlots := make(chan struct{}, 2) // Semaphore with 2 slots (ovens)
	chickenDone := make(chan string, 1)
	stewDone := make(chan string, 1)

	wg.Add(4)
	go bakeLemonCake(sugarLevelChan, &wg, ovenSlots)
	go bakeStrawberryCupcakes(sugarLevelChan, &wg, ovenSlots)
	go grillChicken(&wg, ovenSlots, chickenDone)
	go cookGoatStew(&wg, ovenSlots, stewDone)

	// --- Select Statement Example ---
	// Wait for either chicken or stew to finish first, or timeout after 5 seconds
	select {
	case dish := <-chickenDone:
		fmt.Println(dish, "team finished first and gets served!")
	case dish := <-stewDone:
		fmt.Println(dish, "team finished first and gets served!")
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout: No dish finished in time!")
	}

	wg.Wait()
	fmt.Println("All dishes are ready for the party!")
}
