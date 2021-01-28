package pointers_and_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s ,want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("wanted an error but didnt get one")
		}
		if got.Error() != want {
			t.Errorf("got %s ,want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Deposit(10)
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(100)
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err)
	})

