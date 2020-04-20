package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(0)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestBitcoinStringer(t *testing.T) {
	btc := Bitcoin(10)

	want := "10 BTC"

	got := btc.String()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
