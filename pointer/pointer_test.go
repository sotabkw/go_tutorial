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
