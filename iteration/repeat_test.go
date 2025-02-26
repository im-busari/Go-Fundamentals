package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	assertMessage(t, repeated, expected)
}

// Run with go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 6)
	fmt.Println(repeated)
	// Output: bbbbbb
}

func assertMessage(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}
