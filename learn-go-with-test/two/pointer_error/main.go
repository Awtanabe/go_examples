package pointer_error

import "errors"


type Bitcoin int

type Wallet struct {
  balance Bitcoin
}

func (w *Wallet) Deposit(money Bitcoin) {
  w.balance = money
}

func (w Wallet) Balance() Bitcoin {
  return w.balance
}

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {
  if bitcoin > w.balance {
    return errors.New("預金より多くは引き出せないよ")
  }
  w.balance -= bitcoin
  return nil
}

