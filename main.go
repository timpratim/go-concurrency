package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate deciding the sugar level and sharing it via a channel
func bakeLemonCake(sugarLevelChan chan<- int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{} // acquire semaphore
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
	sem <- struct{}{} // acquire semaphore
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
	sem <- struct{}{} // acquire semaphore
	defer func() { <-sem }() // release semaphore
	fmt.Println("Grilling chicken: Starting...")
	time.Sleep(3 * time.Second)
	fmt.Println("Grilling chicken: Done!")
}

func cookGoatStew(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{} // acquire semaphore
	defer func() { <-sem }() // release semaphore
	fmt.Println("Cooking goat stew: Starting...")
	time.Sleep(4 * time.Second)
	fmt.Println("Cooking goat stew: Done!")
}

func main() {
	var wg sync.WaitGroup
	sugarLevelChan := make(chan int)
	sem := make(chan struct{}, 2) // Semaphore with 2 slots

	wg.Add(4)
	go bakeLemonCake(sugarLevelChan, &wg, sem)
	go bakeStrawberryCupcakes(sugarLevelChan, &wg, sem)
	go grillChicken(&wg, sem)
	go cookGoatStew(&wg, sem)

	wg.Wait()
	fmt.Println("All dishes are ready for the party!")
}
