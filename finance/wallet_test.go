package finance

import "testing"

func TestWallet(t *testing.T) {

	compare := func(t testing.TB, expected Bitcoin, actual Wallet) {
		t.Helper()
		if expected != actual.balance {
			t.Errorf("got %d want %d", actual.balance, expected)
		}
	}

	t.Run("deposit bitcoin to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)
		compare(t, expected, wallet)
	})

	t.Run("withdraw bitcoin from wallet", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		expected := Bitcoin(10)
		if err == nil {
			t.Error("got an error when we shouldn't have")
		}
		compare(t, expected, wallet)

	})

	t.Run("try to withdraw too much bitcoin from wallet", func(t *testing.T) {
		start := Bitcoin(20)
		wallet := Wallet{start}

		err := wallet.Withdraw(Bitcoin(25))

		compare(t, start, wallet)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
