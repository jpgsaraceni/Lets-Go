package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// empty names
	t.Run("empty name and language defaults to english", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty name in Spanish defaults to 'Mundo'", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Mundo"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty name in French defaults to 'Monde'", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, Monde"
		assertCorrectMessage(t, got, want)
	})

	// with names
	t.Run("saying hello to people with empty language defaults to english", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Usted", "Spanish")
		want := "Hola, Usted"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("ma cherie", "French")
		want := "Bonjour, ma cherie"
		assertCorrectMessage(t, got, want)
	})
}
