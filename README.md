# limited-buffer-go
A concurrent bounded buffer system in Go using goroutines, channels, and WaitGroups. Supports multiple producers and consumers with customizable test scenarios.


# 🌀 Bounded Buffer System with Producers and Consumers in Go

This project implements a concurrent system in Go using goroutines and channels to simulate a **bounded buffer**. Multiple producers generate items and place them into the buffer, while multiple consumers remove and process them. The system ensures safe synchronization and graceful shutdown using WaitGroups and channel closing.

## 🚀 Features

- Configurable number of **producers** and **consumers**
- Adjustable **buffer size**
- Safe synchronization using `sync.WaitGroup`
- Communication via buffered `chan`
- Automatic test scenarios with various configurations

## 📦 Project Structure

├── main.go  
├── buffer/ 
 └── buffer.go 
├── dashboard/ │ 
  └── manager.go  
├── processes/ │ 
 └── producer.go  
 └── consumer.go 


 
## 🧪 Test Scenarios

The `main.go` file runs 5 different scenarios, including:

- Varying numbers of producers and consumers
- Small, medium, and large buffer sizes
- Balanced and unbalanced producer/consumer loads

Each test is run sequentially and logs are printed for better visualization.

## 🛠️ How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/limited-buffer-golang.git
   cd limited-buffer-golang
   ```
2. Check Go Modules and Dependencies
  ```bash
    go mod tidy
  ```
3. Execute
  ```bash
   go run main.go
  ```
  
⚠️ Handling Panics
The system gracefully handles concurrency shutdown by properly closing the buffer only after the last producer finishes. If the buffer is closed too early or accessed incorrectly, it may lead to a panic: send on closed channel. This has been fixed in the current version.

📚 Concepts Covered
Goroutines & Channels

Synchronization with WaitGroups

Channel closing patterns

Producer-Consumer Problem (Concurrency 101)

🪪 License
This project is licensed under the MIT License.
