package arrays

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, result []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	}

	t.Run("should sum all tails from the arrays", func(t *testing.T) {
		expected := []int{2, 9}
		result := SumAllTails([]int{0, 2}, []int{3, 9})

		assertCorrectMessage(t, expected, result)
	})

	t.Run("shouldn't make panic when an empty slice is given", func(t *testing.T) {
		expected := []int{0, 9}
		result := SumAllTails([]int{}, []int{3, 4, 5})

		assertCorrectMessage(t, expected, result)
	})
}
