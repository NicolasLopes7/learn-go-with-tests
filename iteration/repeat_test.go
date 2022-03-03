package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	const char = "a"
	const repeatCount = 5

	repeatedCharString := Repeat(char, repeatCount)

	fmt.Println(repeatedCharString)
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, got string) {
		t.Helper()

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}

	}

	t.Run("should repeat the string 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
