package bank

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Ethereum) {
		t.Helper()
		got := wallet.Balance()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Ethereum(10))

		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		assertBalance(t, wallet, Ethereum(10))
	})

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't need one")
		}
	}

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Ethereum(20)}
		err := wallet.Withdraw(Ethereum(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Ethereum(15))
	})

	assertError := func(t testing.TB, err, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't got one")
		}

		if err != want {
			t.Errorf("got %q want %q", err, want)
		}
	}

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Ethereum(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Ethereum(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
