package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	Balance float64
}

func main() {
	myWallet := &Wallet{Balance: 0} // The '&' means this is the REAL wallet address
	var wg sync.WaitGroup           // The Headcounter
	var mu sync.Mutex               // The Key

	for i := 0; i < 10; i++ {
		wg.Add(1) // Tell the counter: "One more worker is starting"

		go func() {
			defer wg.Done() // When this worker finishes, tell the counter: "I'm done"

			mu.Lock()                // 1. Grab the key
			myWallet.Balance += 10.0 // 2. Change the money safely
			mu.Unlock()              // 3. Hand the key back
		}()
	}

	wg.Wait() // Wait until all 10 workers shout "Done!"
	fmt.Println("Final Balance:", myWallet.Balance)
}
