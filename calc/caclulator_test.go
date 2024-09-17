package calc

import "testing"

func TestAddition(t *testing.T) {
	assertEqual(t, 3, Addition{}.Calculate(1, 2))
	assertEqual(t, 4, Addition{}.Calculate(2, 2))
}

func TestSubtraction(t *testing.T) {
	assertEqual(t, 3, Subtraction{}.Calculate(5, 2))
	assertEqual(t, 0, Subtraction{}.Calculate(2, 2))
}
func TestMultiplication(t *testing.T) {
	assertEqual(t, 4, Multiplication{}.Calculate(2, 2))
	assertEqual(t, 3, Multiplication{}.Calculate(1, 3))
}
func TestDivision(t *testing.T) {
	assertEqual(t, 1, Division{}.Calculate(2, 2))
}

func assertEqual(t *testing.T, expected, actual any) {
	if actual != expected {
		t.Errorf("Expected: %v -> Actual: %v", expected, actual)
	}
}
