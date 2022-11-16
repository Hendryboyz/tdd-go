package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	// String() function like ToString() in java
	String() string
}

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// In Go, when you call a function or a method the arguments are copied.
// use pointer herer to change the private variable
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// return (*w).balance is the same notion below
	return w.balance // automatically dereferenced
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return errors.New("Cannot withdraw, insufficient funds!")
	}
	w.balance -= amount
	return nil
}
