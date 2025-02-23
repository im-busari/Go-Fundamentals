package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "en");
		want := "Hello, Chris!";

		assertCorrectMessage(t, got, want);
	});

	t.Run("say 'Hello, World!' when empty string is applied", func(t *testing.T) {
		got := Hello("", "en");	
		want := "Hello, World!";

		assertCorrectMessage(t, got, want);
	});

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Lisa", "es")

		want := "Hola, Lisa!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to English version if language is not available", func(t *testing.T) {
		got := Hello("Lisa", "de");
		want := "Hello, Lisa!";

		assertCorrectMessage(t, got, want);
	});

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lisa", "fr");
		want := "Bonjour, Lisa!";

		assertCorrectMessage(t, got, want);
	});
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper();

	if got != want {
		t.Errorf("got %q want %q", got, want);	
	}
}