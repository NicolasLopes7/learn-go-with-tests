package helloWorld

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("when receive a parameter, should say hello to people", func(t *testing.T) {
		const name = "Nicolas"
		got := Hello(name, "")
		const want = "Hello " + name

		assertCorrectMessage(t, got, want)
	})

	t.Run("when non receive a parameter, should say hello to the world", func(t *testing.T) {

		got := Hello("", "")
		const want = "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("when receive a language parameter as Spanish", func(t *testing.T) {

		got := Hello("Nico", "Spanish")
		const want = "Hola Nico"

		assertCorrectMessage(t, got, want)
	})

	t.Run("when receive a language parameter as French", func(t *testing.T) {

		got := Hello("Nico", "French")
		const want = "Bonjour Nico"

		assertCorrectMessage(t, got, want)
	})
}
