package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	actual := Sum(numbers)
	expected := 15

	if actual != expected {
		t.Errorf("got %d \nexpected %d\n given: %v", actual, expected, numbers)
	}

}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(actual, expected) {
		t.Errorf("got %d \nexpected %d\n", actual, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, actual []int) {
		t.Helper()
		if !slices.Equal(actual, expected) {
			t.Errorf("got %v \nexpected %v\n", actual, expected)
		}

	}

	t.Run("Make the sum of tails", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, expected, actual)

	})

	t.Run("Make the sum of tails with empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, expected, actual)

	})

}
