package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Valya", "")
		want := "Hello, Valya"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when no name is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Pablo", "ES")
		want := "Hola, Pablo"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Camile", "FR")
		want := "Bonjour, Camile"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Gustav", "DE")
		want := "Guten tag, Gustav"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
