package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("test for Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test for Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(0))
	})

}
