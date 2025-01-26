package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}	
	}
	
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello Chris"
	
		assertMessage(t, got, want)
	})

	t.Run("saey hello with empty string", func(t *testing.T){
		got := Hello("", "English")
		want := "Hello World"
		assertMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"
		assertMessage(t, got, want)
	})

	t.Run("say hello in spanish with blank", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola Elodie"
		assertMessage(t, got, want)
	})
}