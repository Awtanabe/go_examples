package main

import (
	"testing"
)


func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if wallet.balance != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	assertErr := func(t *testing.T, got error, want error) {

		if got == nil {
			t.Fatal("wanted err but didn't get one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet {}
		wallet.Deposit(Bitcoin(10))	
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{ balance: Bitcoin(20)}
		wallet.Deposit(10)
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insuffient funds", func (t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		// ダメだったらerrを返すのがgoの慣習
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertErr(t, err, ErrInsufficientFunds)
	})
}