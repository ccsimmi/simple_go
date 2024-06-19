package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 5)
	expected := 7

	assertTest(t, expected, result)
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 1)
	expected := 4

	assertTest(t, expected, result)
}

func TestMultiply(t *testing.T) {
	result := Multiply(6, 3)
	expected := 18

	assertTest(t, expected, result)
}

func assertTest(t *testing.T, expected, result int) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
