package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to a specific person", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Empty string defaults to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"

		assertCorrectMessage(t, got, want)

	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
