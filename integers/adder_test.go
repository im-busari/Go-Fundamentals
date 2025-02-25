package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		sum := Add(3, 2);
		expected := 5;

		if sum != expected {
			// We use %d instead of %q because we want to print an integer rather than strings
			t.Errorf("Expected '%d' but got '%d'", expected, sum);
		}

	})
}