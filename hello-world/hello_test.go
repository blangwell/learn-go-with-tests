package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q Want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Barent", "")
		want := "Hello, Barent!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Kimber", "Spanish")
		want := "Hola, Kimber!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in japanese", func(t *testing.T) {
		got := Hello("Kimber", "Japanese")
		want := "こんにちは, Kimber!"
		assertCorrectMessage(t, got, want)
	})
}
