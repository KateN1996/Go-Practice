package search

import "testing"

func TestSearch(t *testing.T) {
	compare := func(t testing.TB, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("got %q wanted %q", actual, expected)
		}
	}
	t.Run("search dictionary", func(t *testing.T) {
		dict := map[string]string{"test": "just testing lol"}
		actual := Search(dict, "test")
		expected := "just testing lol"
		compare(t, actual, expected)

	})
}
