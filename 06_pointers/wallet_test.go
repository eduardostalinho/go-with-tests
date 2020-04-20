package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(0)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(11))

		if err == nil {
			t.Error("expected error on withdraw")
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
