package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate deciding the sugar level and sharing it via a channel
func bakeLemonCake(sugarLevelChan chan<- int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}        // acquire semaphore
	defer func() { <-sem }() // release semaphore
	fmt.Println("Baking lemon cake: Deciding sugar level...")
	time.Sleep(1 * time.Second) // Simulate time to decide
	sugarLevel := 5
	fmt.Printf("Baking lemon cake: Sugar level decided: %d\n", sugarLevel)
	sugarLevelChan <- sugarLevel // Send sugar level to channel
	fmt.Println("Baking lemon cake: Baking in progress...")
	time.Sleep(2 * time.Second)
	fmt.Println("Baking lemon cake: Done!")
}

func bakeStrawberryCupcakes(sugarLevelChan <-chan int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}        // acquire semaphore
	defer func() { <-sem }() // release semaphore
	fmt.Println("Baking strawberry cupcakes: Waiting for sugar level from lemon cake...")
	sugarLevel := <-sugarLevelChan // Wait for sugar level
	fmt.Printf("Baking strawberry cupcakes: Received sugar level: %d\n", sugarLevel)
	fmt.Println("Baking strawberry cupcakes: Baking in progress...")
	time.Sleep(2 * time.Second)
	fmt.Println("Baking strawberry cupcakes: Done!")
}

func grillChicken(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}        // acquire semaphore
	defer func() { <-sem }() // release semaphore
	fmt.Println("Grilling chicken: Starting...")
	time.Sleep(3 * time.Second)
	fmt.Println("Grilling chicken: Done!")
}

func cookGoatStew(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}        // acquire semaphore
	defer func() { <-sem }() // release semaphore
	fmt.Println("Cooking goat stew: Starting...")
	time.Sleep(4 * time.Second)
	fmt.Println("Cooking goat stew: Done!")
}

func main() {
	var wg sync.WaitGroup
	// Buffered channel demonstration
	bufferedChan := make(chan int, 2)
	fmt.Println("Demonstrating buffered channel behavior:")
	fmt.Println("Sending 1 to bufferedChan...")
	bufferedChan <- 1 // does not block
	fmt.Println("Sent 1 (no block)")
	fmt.Println("Sending 2 to bufferedChan...")
	bufferedChan <- 2 // does not block
	fmt.Println("Sent 2 (no block)")
	// The next send would block because buffer is full, so we'll comment it
	// fmt.Println("Sending 3 to bufferedChan (will block)...")
	// bufferedChan <- 3 // would block here
	fmt.Println("Receiving from bufferedChan:")
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println("Buffered channel demonstration complete.")
	sugarLevelChan := make(chan int, 2) // Now buffered
	sem := make(chan struct{}, 2)       // Semaphore with 2 slots

	wg.Add(4)
	go bakeLemonCake(sugarLevelChan, &wg, sem)
	go bakeStrawberryCupcakes(sugarLevelChan, &wg, sem)
	go grillChicken(&wg, sem)
	go cookGoatStew(&wg, sem)

	wg.Wait()
	fmt.Println("All dishes are ready for the party!")
}
