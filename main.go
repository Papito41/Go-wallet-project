package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Balance float64
	History []float64
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.Balance {
		return errors.New("insufficient funds withdrawal")
	}
	w.Balance -= amount
	w.History = append(w.History, -amount)
	return nil
}
func main() {
	myWallet := Wallet{Balance: 100.00}

	err := myWallet.Withdraw(500.00)
	if err != nil {
		fmt.Println("CRITICAL ERROR: ", err)
		return
	}
	fmt.Println("Success! New balance:", myWallet.Balance)

}
