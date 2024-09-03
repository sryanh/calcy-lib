package calc

import "testing"

func TestAddition(t *testing.T) {
	assertEqual(t, 3, Addition{}.Calculate(1, 2))
	assertEqual(t, 4, Addition{}.Calculate(2, 2))
}

func assertEqual(t *testing.T, expected, actual any) {
	if actual != expected {
		t.Errorf("Expected: %v -> Actual: %v", expected, actual)
	}
}
