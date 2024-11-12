package bank

import (
	"errors"
	"fmt"
)

type Ethereum int

func (e Ethereum) String() string {
	return fmt.Sprintf("%d ETH", e)
}

type Wallet struct {
	balance Ethereum
}

// struct pointers are automatically dereferenced.
func (w *Wallet) Deposit(amount Ethereum) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("insufficient funds to withdraw")

func (w *Wallet) Withdraw(amount Ethereum) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Ethereum {
	return w.balance
}
