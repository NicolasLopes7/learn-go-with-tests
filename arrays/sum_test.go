package arrays

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, result int) {
		t.Helper()

		if expected != result {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		const expected = 6
		result := Sum(numbers)

		assertCorrectMessage(t, expected, result)

	})

}
