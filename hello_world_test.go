package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to a person", func(t *testing.T) {
		got := Hello("Kate")
		want := "Hello Kate :)"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to the world when given an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello World :)"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
