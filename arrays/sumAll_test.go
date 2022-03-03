package arrays

import (
	"reflect"
	"testing"
)

func TestSumALl(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, result []int) {
		t.Helper()

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	}

	t.Run("Should return an array with the sum of each array given", func(t *testing.T) {

		expected := []int{3, 9}
		result := SumAll([]int{1, 2}, []int{0, 9})
		assertCorrectMessage(t, expected, result)
	})
}
