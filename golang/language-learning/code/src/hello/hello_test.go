package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectHello := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	//3.9 add Spanish cases
	t.Run("this is testing for Spanish hello", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose"
		assertCorrectHello(t, got, want)
	})

	t.Run("this is testing for hello with name", func(t *testing.T) {
		got := Hello("Bruce", "English")
		want := "Hello, Bruce"

		assertCorrectHello(t, got, want)
	})

	t.Run("this is testing for hello with empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectHello(t, got, want)
	})

	t.Run("this is testing for French hello", func(t *testing.T) {
		got := Hello("Louis", "French")
		want := "Bonjour, Louis"

		assertCorrectHello(t, got, want)
	})

}
