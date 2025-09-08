package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	actual := Sum(numbers)
	expected := 15

	if actual != expected {
		t.Errorf("got %d \nexpected %d\n given: %v", actual, expected, numbers)
	}

}
