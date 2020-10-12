package main

import "testing"

func TestAdder(t *testing.T) {

	assertCorrectSum := func(t *testing.T, expected, sum int) {
		t.Helper()
		if expected != sum {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectSum(t, expected, sum)
	})

	t.Run("", func(t *testing.T) {
		sum := Add(1, 2)
		expected := 3
		assertCorrectSum(t, expected, sum)
	})
}
