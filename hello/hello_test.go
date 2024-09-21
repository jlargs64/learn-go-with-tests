package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to justin", func(t *testing.T) {
		got := Hello("Justin", English)
		want := "Hello, Justin"

		assetCorrectMessage(t, got, want)
	})

	t.Run("saying hello when an empty string is provided", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello, world"

		assetCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Nerds", Spanish)
		want := "Hola, Nerds"

		assetCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Nerds", French)
		want := "Bonjour, Nerds"

		assetCorrectMessage(t, got, want)
	})
}

func assetCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
