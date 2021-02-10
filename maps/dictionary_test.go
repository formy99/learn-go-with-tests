package maps

import "testing"

func assertMessage(t testing.TB, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q ,want %q ,given %q", got, want, given)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertMessage(t, got, want, "test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertMessage(t, err.Error(), want, "unknown")

	})

}
