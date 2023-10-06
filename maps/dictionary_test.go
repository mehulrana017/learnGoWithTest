package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "This is the test"}

	t.Run("known words", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "This is the test"

		assertString(t, got, want)
	})

	t.Run("unknown words", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)

	})

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
