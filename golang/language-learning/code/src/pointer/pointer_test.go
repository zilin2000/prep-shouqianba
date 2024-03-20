package pointer

import (
	"fmt"
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

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("want an error but didn't")
		}

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		withdrawAmount := Bitcoin(100)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(withdrawAmount)

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, InsufficientFundsError)
	})

}

func TestString(t *testing.T) {
	wallet := Wallet{balance: Bitcoin(20)}

	myBalance := wallet.balance

	fmt.Println(myBalance)
}
