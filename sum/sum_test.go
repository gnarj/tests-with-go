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

func TestSumAll(t *testing.T) {
	t.Run("test summing multiple slices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		sum := SumAll([]int{1, 2, 3, 4, 5, 6}, []int{1, 4, 3, 6, 5, 6})
		expected := []int{21, 25}

		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("received %d, expected %d, given %v", sum, expected, numbers)
		}
	})
}
