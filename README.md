# Go concurrency

This project demonstrates Go's concurrency primitives through a fun, food-themed example. It includes both a Go program simulating a party kitchen with multiple concurrent tasks, and an interactive HTML visualization to help understand how goroutines and channels work together.

## Project Structure

- `main.go` — The main Go program simulating a kitchen where several dishes (lemon cake, strawberry cupcakes, grilled chicken, goat stew) are prepared concurrently using goroutines, channels, WaitGroups, and semaphores.
- `concurrency_viz.html` — An interactive browser-based visualization of the Go program's concurrency, showing the flow and synchronization between goroutines.
- `go.mod` — Go module definition.

## How It Works

- **Lemon Cake**: Decides a sugar level and shares it via a channel.
- **Strawberry Cupcakes**: Waits for the sugar level from the lemon cake before starting.
- **Grilled Chicken & Goat Stew**: Start independently, limited by a semaphore (max 2 concurrent tasks).
- **Synchronization**: All tasks use a `sync.WaitGroup` to notify when they're done. The main goroutine waits for all dishes to finish.

## Running the Go Program

1. Make sure you have Go 1.23+ installed.
2. In your terminal, run:
   ```sh
   go run main.go
   ```
3. You should see output describing the concurrent preparation of each dish and the coordination between them.

## Viewing the Visualization

1. Open `concurrency_viz.html` in your web browser.
2. Click **Start Baking Party Visualization** to see each step animated, including channel communication and task completion.

## Learning Objectives

- Understand how goroutines run concurrently in Go.
- See how channels synchronize data between tasks.
- Visualize the effect of semaphores and WaitGroups in managing concurrency.

## License

This project is for educational purposes.
