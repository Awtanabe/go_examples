package pointer

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// このWalletが*じゃないと更新されないから上手く行かない
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	// 引き出す額が、残高より多い場合はエラーを返す
	if amount > w.balance {
		  // ここは学習項目
			return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil
}