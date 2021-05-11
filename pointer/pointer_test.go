package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	got := wallet.Balance()

	want := BitCoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestWalletSecond(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(BitCoin(10))

		got := wallet.Balance()

		want := BitCoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}

		wallet.Withdraw(BitCoin(10))

		got := wallet.Balance()

		want := BitCoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
