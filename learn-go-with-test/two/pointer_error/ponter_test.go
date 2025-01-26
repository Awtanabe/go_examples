package pointer_error

import (
	"testing"
)


func TestWallet(t *testing.T) {

	checkOKAssert := func(t *testing.T, got, want Bitcoin){
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("Depositi", func(t *testing.T) {
		wallet := Wallet{}
		// 預ける
		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		checkOKAssert(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		// 預ける
		wallet.Deposit(10)
		wallet.Withdraw(5)
		got := wallet.Balance()
		want := Bitcoin(5)
		checkOKAssert(t, got, want)
	})

	t.Run("Withdraw with error", func(t *testing.T) {
		wallet := Wallet{}
		// 預ける
		wallet.Deposit(10)
		err := wallet.Withdraw(15)
		want := "預金より多くは引き出せないよ"

		if err == nil {
			t.Error("want en error")
		}

		if err.Error() != want {
			t.Errorf("want error %q got %q", err, want)
		}

	})
}