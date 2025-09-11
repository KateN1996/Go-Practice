package finance

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return fmt.Errorf(
			"tried to withdraw more than you had in the account\naccount balance: %v, amount tried to withdraw: %v",
			w.balance,
			amount,
		)
	}
	w.balance -= amount
	return nil
}
