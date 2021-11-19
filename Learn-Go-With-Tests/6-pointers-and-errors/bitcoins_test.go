package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("show balance 10 BTC after depositing 10 BTC", func(t *testing.T) {
		wallet := Wallet{}

		t.Helper()
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertValue(t, got, want)
	})

	t.Run("show balance after withdrawing 10 from a wallet with initial balance of 30", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}

		t.Helper()
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(20)

		assertValue(t, got, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertValue(t, wallet.Balance(), startingBalance)
		assertError(t, err)
	})
}

func assertValue(t testing.TB, got, want Bitcoin) {

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
}
