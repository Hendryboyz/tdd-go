package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	// arrange
	wallet := Wallet{}
	wallet.Deposit(10)

	// action
	actual := wallet.Balance()
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	// assert
	expected := 10.0
	if actual != expected {
		t.Errorf("Expected %.2f, but actual %.2f", expected, actual)
	}
}
