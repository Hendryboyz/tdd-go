package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBitcoin := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()
		if actual != expected {
			t.Errorf("Expected %s, but actual %s", expected, actual)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		// arrange
		wallet := Wallet{}

		// action
		wallet.Deposit(Bitcoin(64.2))

		// assert
		assertBitcoin(t, wallet, Bitcoin(64.2))
	})

	t.Run("Withdraw", func(t *testing.T) {
		// Arrange
		wallet := Wallet{balance: Bitcoin(30.2)}

		// Action
		wallet.Withdraw(Bitcoin(15.2))

		// Assert
		assertBitcoin(t, wallet, Bitcoin(15.0))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		// Arrange
		startingBalance := Bitcoin(10.2)
		wallet := Wallet{balance: startingBalance}

		// Action
		err := wallet.Withdraw(Bitcoin(16.2))

		// Assert
		assertBitcoin(t, wallet, startingBalance)
		if err == nil {
			t.Error("Need to return an error while withdrawing the amount exceed the balance.")
		}
	})
}
