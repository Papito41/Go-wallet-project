package main

import "fmt"

type balance struct {
	amount float64
}

// Move this OUTSIDE of main
func addMoney(b *float64, amount float64) {
	if amount > 0 {
		*b += amount
	} else {
		fmt.Println("Error: Cannot deposit zero or negative money!")
	}
}

func main() {
	myWallet := balance{amount: 100.00}

	fmt.Println("Initial Balance:", myWallet.amount)

	// Use the pointer to change the original amount
	addMoney(&myWallet.amount, 50.50)

	fmt.Println("New Balance:", myWallet.amount)
}
