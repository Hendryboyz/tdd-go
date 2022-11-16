package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		// arrange
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(64.2))

		// action
		actual := wallet.Balance()
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		// assert
		expected := Bitcoin(64.2)
		if actual != expected {
			t.Errorf("Expected %s, but actual %s", expected, actual)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		// Arrange
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(30.2))

		// Action
		wallet.Withdraw(Bitcoin(15.2))

		// Assert
		actual := wallet.Balance()
		expected := Bitcoin(15.0)
		if actual != expected {
			t.Errorf("Expected %s, but actual %s", expected, actual)
		}
	})

}
