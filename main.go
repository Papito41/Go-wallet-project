package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 1. Updated Wallet: Added gorm.Model so it can be saved to a database
type Wallet struct {
	gorm.Model
	Balance float64
	// mu sync.Mutex  // We'll bring the Mutex back in a second if needed
}

func main() {
	// 2. Initialize the Database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 3. Create the table
	db.AutoMigrate(&Wallet{})

	// 4. Create our wallet in the database (Starting at 0)
	myWallet := Wallet{Balance: 0}
	db.Create(&myWallet)

	// 5. Your Concurrency Loop from earlier
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// For now, we update the local object
			myWallet.Balance += 10.0

			// PRO TIP: Tomorrow we will learn how to save this
			// back to the DB safely inside the goroutine!
		}()
	}

	wg.Wait()

	// 6. Save the final result back to the database
	db.Save(&myWallet)

	fmt.Printf("Final Balance saved to DB: %.2f\n", myWallet.Balance)
}
