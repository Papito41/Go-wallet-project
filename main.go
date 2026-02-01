package main

import "fmt"

// 1. THE BLUEPRINT (Struct)
type Wallet struct {
	Balance float64
	History []float64 // The Slice (The Train)
}

// 2. THE DEPOSIT TOOL (Pointer Receiver)
func (w *Wallet) Deposit(amount float64) {
	w.Balance = w.Balance + amount
	w.History = append(w.History, amount) // Adding a "car" to the train
	fmt.Printf("Deposited: $%.2f\n", amount)
}

// 3. THE WITHDRAW TOOL (Logic + Pointer)
func (w *Wallet) Withdraw(amount float64) {
	if amount <= w.Balance {
		w.Balance = w.Balance - amount
		w.History = append(w.History, -amount) // Record as a negative
		fmt.Printf("Withdrew: $%.2f\n", amount)
	} else {
		fmt.Println("Error: Insufficient funds!")
	}
}

// 4. THE EXECUTION
func main() {
	// Initialize the wallet
	myWallet := Wallet{Balance: 100.00}

	// Perform actions
	myWallet.Deposit(50.00)
	myWallet.Withdraw(30.00)
	myWallet.Withdraw(200.00) // This will trigger the "if" error

	// See the results
	fmt.Println("---------------------------")
	fmt.Printf("Final Balance: $%.2f\n", myWallet.Balance)
	fmt.Println("Full Statement:", myWallet.History)
}
