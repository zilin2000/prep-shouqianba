package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know words", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown words", func(t *testing.T) {
		_, got := dictionary.Search("unknow")

		assertError(t, got, ErrNotFound)
	})
}

func assertError(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("got error %s want %s", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
