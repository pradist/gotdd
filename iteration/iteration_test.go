package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrect := func(t *testing.T, expected, repeated string) {

		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("a to aaaaa", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrect(t, expected, repeated)
	})

	t.Run("b to bbbbb", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"
		assertCorrect(t, expected, repeated)
	})

	t.Run("b to bbbbbb", func(t *testing.T) {
		repeated := Repeat("b", 6)
		expected := "bbbbbb"
		assertCorrect(t, expected, repeated)
	})

	t.Run("a to aaaaaaa", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"
		assertCorrect(t, expected, repeated)
	})

	t.Run("a to ", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrect(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}
