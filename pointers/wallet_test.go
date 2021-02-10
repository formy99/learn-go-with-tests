package pointers

import (
	"testing"
)

func assertMessage(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s ,want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}

}
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but did")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestWallet(t *testing.T) {
	t.Run("Despoint Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Despoint(10)
		want := Bitcoin(10)
		assertMessage(t, wallet, want)
	})

	t.Run("Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertMessage(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startigBalance := Bitcoin(20)
		wallet := Wallet{startigBalance}
		err := wallet.Withdraw(Bitcoin(10))
		assertMessage(t, wallet, Bitcoin(10))
		assertError(t, err, ErrInsufficientFunds)
	})

}
