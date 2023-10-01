package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "spanish")
		want := "hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := hello("", "")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
