package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectHello := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("this is testing for hello with name", func(t *testing.T) {
		got := Hello("Bruce")
		want := "Hello, Bruce"

		assertCorrectHello(t, got, want)
	})

	t.Run("this is testing for hello with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectHello(t, got, want)
	})

}
