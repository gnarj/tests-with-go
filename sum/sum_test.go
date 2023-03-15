package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test summing slice of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		sum := Sum(numbers)
		expected := 21

		if sum != expected {
			t.Errorf("received %d, expected %d, given %v", sum, expected, numbers)
		}
	})
}

func TestSumTails(t *testing.T) {
	checkSums := func(t testing.TB, sum, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("received %d, expected %d", sum, expected)
		}
	}
	t.Run("test summing the tails of slices", func(t *testing.T) {
		sum := SumAllTails([]int{1, 2, 3, 4, 5, 6}, []int{1, 4, 3, 6, 5, 6})
		expected := []int{20, 24}

		checkSums(t, sum, expected)

	})
	t.Run("safely sum all the tails", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{1, 4, 3, 6, 5, 6})
		expected := []int{0, 24}

		checkSums(t, sum, expected)
	})
}
