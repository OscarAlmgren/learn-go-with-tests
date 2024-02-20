package pointerserrors

import "fmt"

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() (balance Bitcoin) {
	balance = w.balance
	return
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
