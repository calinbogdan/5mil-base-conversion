package main

import "testing"

func TestConvertsToCorrectBase(t *testing.T) {

	t.Run("2 in base 3 is 2 in base 10", func (t *testing.T) {
		got := convertToBase10(2, 3)
		want := 2

		assertConversion(t, got, want)
	})

	t.Run("25 in base 10 should be 31 in base 8", func(t *testing.T) {
		got := convertToBase(25, 8)
		want := 31

		assertConversion(t, got, want)
	})

	t.Run("21345 in base 10 should return 116142 in base 7", func (t *testing.T) {
		got := convertFromBaseTo(21345, 10, 7)
		want := 116142

		assertConversion(t, got, want)
	})

	t.Run("5 in base 10 should return 5 in base 6", func (t *testing.T) {
		got := convertFromBaseTo(5, 10, 6)
		want := 5

		assertConversion(t, got, want)
	})
}


func assertConversion(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
