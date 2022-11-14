package wallet

import "fmt"

type Wallet struct {
	balance float64
}

// In Go, when you call a function or a method the arguments are copied.
// use pointer herer to change the private variable
func (w *Wallet) Deposit(amount float64) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	// return (*w).balance is the same notion below
	return w.balance // automatically dereferenced
}
