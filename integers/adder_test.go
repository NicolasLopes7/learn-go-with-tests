package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, result int) {
		t.Helper()
		if result != expected {
			t.Errorf("expected '%d' but got '%d'", expected, result)
		}
	}

	t.Run("should add to integers correctly", func(t *testing.T) {
		sum := Add(2, 2)
		const expected = 4

		assertCorrectMessage(t, expected, sum)
	})
}
