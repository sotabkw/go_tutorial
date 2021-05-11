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

	assertBalance := func(t *testing.T, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		wallet.Withdraw(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})
}
