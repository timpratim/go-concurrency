package main

import (
	"fmt"
	"sync"
	"time"
)

func useOven(dish string, ovenSlots chan struct{}) {
	ovenSlots <- struct{}{} // acquire
	fmt.Println("Using oven for", dish)
	time.Sleep(2 * time.Second) // simulate oven time
	<-ovenSlots // release
	fmt.Println("Done with", dish)
}

// Simulate deciding the sugar level and sharing it via a channel
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

func bakeStrawberryCupcakes(sugarLevelChan <-chan int, wg *sync.WaitGroup, ovenSlots chan struct{}) {
	defer wg.Done()
	fmt.Println("Baking strawberry cupcakes: Waiting for sugar level from lemon cake...")
	sugarLevel := <-sugarLevelChan // Wait for sugar level
	fmt.Printf("Baking strawberry cupcakes: Received sugar level: %d\n", sugarLevel)
	useOven("strawberry cupcakes", ovenSlots)
	fmt.Println("Baking strawberry cupcakes: Done!")
}

func grillChicken(wg *sync.WaitGroup, ovenSlots chan struct{}) {
	defer wg.Done()
	useOven("grilled chicken", ovenSlots)
	fmt.Println("Grilling chicken: Done!")
}

func cookGoatStew(wg *sync.WaitGroup, ovenSlots chan struct{}) {
	defer wg.Done()
	useOven("goat stew", ovenSlots)
	fmt.Println("Cooking goat stew: Done!")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
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
	}()

	sugarLevelChan := make(chan int, 2) // Now buffered
	ovenSlots := make(chan struct{}, 2) // Semaphore with 2 slots (ovens)

	wg.Add(4)
	go bakeLemonCake(sugarLevelChan, &wg, ovenSlots)
	go bakeStrawberryCupcakes(sugarLevelChan, &wg, ovenSlots)
	go grillChicken(&wg, ovenSlots)
	go cookGoatStew(&wg, ovenSlots)

	wg.Wait()
	fmt.Println("All dishes are ready for the party!")
}
