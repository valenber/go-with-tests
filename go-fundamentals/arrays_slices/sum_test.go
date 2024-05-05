package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		assertResult(t, numbers, got, want)
	})
}

func assertResult(t *testing.T, input []int, got int, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d instead of %d given %v", got, want, input)
	}
}
