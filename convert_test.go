package main

import "testing"

func TestConvertsToCorrectBase(t *testing.T) {

	t.Run("2 in base 3 is 2 in base 10", func(t *testing.T) {
		got := convertToBase10(2, 3)
		want := 2

		assertConversion(t, got, want)
	})

	t.Run("25 in base 10 should be 31 in base 8", func(t *testing.T) {
		got := convertToBase(25, 8)
		want := 31

		assertConversion(t, got, want)
	})

	t.Run("21345 in base 10 should return 116142 in base 7", func(t *testing.T) {
		got := convertFromBaseTo(21345, 10, 7)
		want := 116142

		assertConversion(t, got, want)
	})

	t.Run("5 in base 10 should return 5 in base 6", func(t *testing.T) {
		got := convertFromBaseTo(5, 10, 6)
		want := 5

		assertConversion(t, got, want)
	})
}

func TestDigits(t *testing.T) {
	t.Run("42 to base 3 digits is 1120", func(t *testing.T) {
		got := toDigits(42, 3)
		want := "1120"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("1120 in base 3 is 42", func(t *testing.T) {
		got := fromDigits("1120", 3)
		want := 42

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func BenchmarkConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertFromBaseTo(12561, 8, 9)
	}
}

func BenchmarkDigitsConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertBaseWithDigits("1120", 3, 6)
	}
}

func assertConversion(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
