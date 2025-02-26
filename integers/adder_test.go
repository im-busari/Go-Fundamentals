package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		sum := Add(3, 2)
		expected := 5

		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		// We use %d instead of %q because we want to print an integer rather than strings
		t.Errorf("Expected '%d' but got '%d'", want, got)
	}
}
