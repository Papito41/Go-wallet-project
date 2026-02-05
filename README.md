# Go Wallet: Concurrency & Persistence Exploration

This project is a hands-on exploration of the Go programming language, specifically focusing on the **Standard Library** and **Concurrency** patterns. 

## 🚀 Project Overview
I built a digital wallet system that handles multiple deposit transactions simultaneously. This project demonstrates how Go manages memory and thread safety using built-in synchronization primitives.

## 🧠 Key Concepts Covered (Week 2 Preview)
Although I am currently mastering Week 1 fundamentals, this code explores:
* **Goroutines**: Using the `go` keyword to run functions in the background.
* **WaitGroups (`sync.WaitGroup`)**: Orchestrating multiple workers to ensure the main program waits for all tasks to complete.
* **Mutexes (`sync.Mutex`)**: Implementing a "Mutual Exclusion" lock to prevent **Race Conditions** when multiple workers access the same balance.
* **Pointers**: Using memory addresses to update state across different functions.

## 🛠️ How it Works
1. The program initializes a `Wallet` struct with a `Balance`.
2. It spawns **10 concurrent goroutines**, each attempting to deposit $10.0.
3. A **Mutex** ensures that only one worker can update the balance at a time, ensuring the final balance is exactly $100.0.



## 🏗️ Future Roadmap
- [ ] Connect to a persistent database (PostgreSQL/SQLite) using **GORM**.
- [ ] Implement **Channels** for communication between workers.
- [ ] Build a REST API using the `net/http` standard library.

---
*Learning journey guided by the February Go Bootcamp Roadmap.*