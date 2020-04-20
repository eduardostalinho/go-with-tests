package wallet

import "testing"

func TestWallet(t *testing.T) {

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

		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, ErrInsufficientFunds)
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

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error but didn't got one")
	}

	if err != want {
		t.Errorf("expected error %q got %q", want, err)
	}
}
