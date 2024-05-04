package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("returns correct defualt number of chars", func(t *testing.T) {

		output := Repeat("f", 0)
		expected := "fffff"

		assertCorrectMessage(t, output, expected)
	})

	t.Run("returns correct specified number of chars", func(t *testing.T) {
		output := Repeat("m", 3)
		expected := "mmm"

		assertCorrectMessage(t, output, expected)
	})
}

func assertCorrectMessage(t testing.TB, output, expected string) {
	t.Helper()

	if output != expected {
		t.Errorf("expected %q, but got %q", expected, output)
	}
}

func ExampleRepeat_default() {
	output := Repeat("x", 0)
	fmt.Println(output)
	// Output: xxxxx
}

func ExampleRepeat_specific() {
	output := Repeat("a", 3)
	fmt.Println(output)
	// Output: aaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
