package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats 'a' 5 times", func(t *testing.T) {
		expected := "aaaaa"
		actual := Repeat("a", 5)
		assertCorrectIteration(t, actual, expected)
	})

	t.Run("repeats 'b' 0 times", func(t *testing.T) {
		expected := ""
		actual := Repeat("b", 0)
		assertCorrectIteration(t, actual, expected)
	})

}

func assertCorrectIteration(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but received %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}
