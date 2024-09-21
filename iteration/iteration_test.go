package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestGenerateCommaList(t *testing.T) {
	t.Run("successfully generate list", func(t *testing.T) {
		got := GenerateCommaList([]string{"foo", "bar", "baz"})
		expected := "foo,bar,baz"
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})

	t.Run("generates nothing when list is empty", func(t *testing.T) {
		var emptyList []string
		got := GenerateCommaList(emptyList)
		expected := ""
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("Justin", 3)
	fmt.Println(repeated)
	// Output: JustinJustinJustin
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
